package test

import (
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/api"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func handlers() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/infinity", func(w http.ResponseWriter, r *http.Request) {
		for {
			time.Sleep(1 * time.Second)
		}
	})

	r.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("{\"ok\":1}"))
		if err != nil {
			log.Println(err)
		}

	})

	r.HandleFunc("/404", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	r.HandleFunc("/403", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
	})

	r.HandleFunc("/500", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	})

	return r
}

func TestClient(t *testing.T) {

	client := api.NewClient()

	if client.Timeout() != api.ClientDefaultTimeout {
		t.Error(
			"Expected", api.ClientDefaultTimeout,
			"\ngot", client.Timeout(),
		)
	}
}

func TestClient_Timeout(t *testing.T) {

	log.SetFlags(log.Lshortfile)

	chSrv := make(chan string)
	go func(ch chan string) {
		srv := httptest.NewServer(handlers())
		//srv := httptest.NewServer(handlers())
		defer srv.Close()
		ch <- srv.URL
		<-ch
	}(chSrv)
	defer func() { chSrv <- "" }()

	var conf = &app.Config{
		Credentials: app.Credentials{
			Token: "token string",
			Type:  "Bearer",
		},
		Api: app.Api{
			Url: <-chSrv,
		},
	}

	client := api.NewClient()

	apiContext := app.NewContext(conf)
	apiInstance, err := api.New(apiContext)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	request, err := apiInstance.CreateRequest("GET", "/infinity")
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	requestCh := make(chan bool)
	requestTm := func(done chan bool) {
		_, _ = client.DoRequest(request, true)
		done <- true
	}

	for _, timeout := range []uint{1, 2, 10} {
		t.Logf("Start timeout %v sec\n", timeout)

		client.SetTimeout(timeout)
		timer := time.NewTimer(time.Duration(client.Timeout()+1) * time.Second)

		go requestTm(requestCh)
		//timeCh := make(chan bool)

		select {
		case <-timer.C:
			t.Error("Unexpected error: no timeout")
			break
		case <-requestCh:
			t.Logf("PASS timeout %v sec\n\n", timeout)
			break //OK
		}
	}
}

func TestClient_TimeoutPanic(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic")
		}
	}()

	client := api.NewClient()
	client.SetTimeout(0)
}

func TestClient_Response(t *testing.T) {

	log.SetFlags(log.Lshortfile)

	chSrv := make(chan string)
	go func(ch chan string) {
		srv := httptest.NewServer(handlers())
		defer srv.Close()
		ch <- srv.URL
		<-ch
	}(chSrv)
	defer func() { chSrv <- "" }()

	var conf = &app.Config{
		Credentials: app.Credentials{
			Token: "token string",
			Type:  "Bearer",
		},
		Api: app.Api{
			Url: <-chSrv,
		},
	}

	apiContext := app.NewContext(conf)
	apiInstance, err := api.New(apiContext)
	if err != nil {
		t.Error("Unexpected error:", err)
	}
	client := apiInstance.Client()

	var localTests = []struct {
		name         string
		path         string
		responseText string
	}{
		{"ok", "/ok", "{\"ok\":1}"},
		{"404", "/404", "404 " + http.StatusText(http.StatusNotFound)},
		{"403", "/403", "403 " + http.StatusText(http.StatusForbidden)},
		{"500", "/500", "500 " + http.StatusText(http.StatusInternalServerError)},
	}

	for _, test := range localTests {
		t.Run(test.name, func(t *testing.T) {
			request, _ := apiInstance.CreateRequest("GET", test.path)
			response, err := client.DoRequest(request, true)

			fmt.Println(string(response), err)

			if err == nil {
				if test.responseText != string(response) {
					t.Error(
						fmt.Sprintln(test.name),
						fmt.Sprintf("\nExpected: '%v'\n", test.responseText),
						fmt.Sprintf("    Got: '%v'\n", string(response)),
					)
				}
			} else {
				if test.responseText != err.Error() {
					t.Error(
						fmt.Sprintln(test.name),
						fmt.Sprintf("\nExpected: '%v'\n", test.responseText),
						fmt.Sprintf("    Got: '%v'\n", err.Error()),
					)
				}
			}

		})
	}
}
