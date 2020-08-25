package api

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

const ClientDefaultTimeout = 30

type Client interface {
	DoRequest(request *Request, debug bool) ([]byte, error)
	SetTimeout(timeout uint)
	Timeout() uint
}

type client struct {
	defaultTimeout uint
}

func NewClient() Client {
	return &client{ClientDefaultTimeout}
}

func (c *client) SetTimeout(timeout uint) {
	if timeout == 0 {
		log.Panicln("Expected timeout >= 0")
	}
	c.defaultTimeout = timeout
}

func (c *client) Timeout() uint {
	return c.defaultTimeout
}

func (c *client) DoRequest(request *Request, debug bool) (data []byte, err error) {

	if debug {
		fmt.Println("************* REQUEST ***************")
		if requestDump, err := httputil.DumpRequest(request.Request, true); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(requestDump))
		}
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	defer httpClient.CloseIdleConnections()

	if c.Timeout() > 0 {
		httpClient.Timeout = time.Duration(c.Timeout()) * time.Second
	}

	request.Close = true
	var response *http.Response
	response, err = httpClient.Do(request.Request)

	if err == nil {
		response.Close = true

		if debug {
			fmt.Println("************* RESPONSE ***************")
			if responseDump, err := httputil.DumpResponse(response, true); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(responseDump))
			}
		}

		if response.StatusCode != 200 {
			err = errors.New(response.Status)
		} else {
			data, err = ioutil.ReadAll(response.Body)
		}
	}
	return data, err
}
