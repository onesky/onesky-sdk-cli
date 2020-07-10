package test

import (
	"OneSky-cli/pkg/api"
	"OneSky-cli/pkg/app"
	"reflect"
	"testing"
)

var confDefault = &app.Config{
	Credentials: app.Credentials{
		Token: "token string",
		Type:  "Bearer",
	},
	Api: app.Api{
		Url: "https://management-api.onesky.app/v1",
	},
}

var confNoAuthType = &app.Config{
	Credentials: app.Credentials{
		Token: "token string",
		Type:  "",
	},
	Api: app.Api{
		Url: "https://management-api.onesky.app/v1",
	},
}

var confNoAuth = &app.Config{
	Credentials: app.Credentials{
		Token: "",
		Type:  "",
	},
	Api: app.Api{
		Url: "https://management-api.onesky.app/v1",
	},
}

var confInvalidUrl = &app.Config{
	Credentials: app.Credentials{
		Token: "token string",
		Type:  "Bearer",
	},
	Api: app.Api{
		Url: "management-api.onesky.app/v1",
	},
}

var defaultCtx = app.NewContext(confDefault)
var defaultApi, _ = api.New(defaultCtx)

func Test_api_New(t *testing.T) {
	type args struct {
		context app.Context
	}
	tests := []struct {
		name string
		args args
		want api.Api
	}{
		{name: "OK", args: args{defaultCtx}, want: defaultApi},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := api.New(defaultCtx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_api_Client(t *testing.T) {

	client1 := defaultApi.Client()
	client2 := defaultApi.Client()

	if !reflect.DeepEqual(client1, client2) {
		t.Errorf("Expected 'singletone', but got new instance")
	}

	if _, ok := interface{}(client1).(api.Client); !ok {
		t.Errorf("Expected 'Client' interface")
	}
}

func Test_api_CreateRequest(t *testing.T) {

	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name string
		args args
		pass bool
		conf app.Config
	}{
		{name: "OK", args: args{"GET", "/test1"}, pass: true, conf: *confDefault},
		{name: "F_URL", args: args{"GET", "/test3"}, pass: false, conf: *confInvalidUrl},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			newApi, e := api.New(app.NewContext(&tt.conf))
			if e != nil {
				if tt.pass {
					t.Error(e.Error())
					return
				} else {
					return // OK
				}
			}

			path := tt.args.path
			method := tt.args.method

			r, e := newApi.CreateRequest(method, path)
			if e != nil {
				t.Error(e.Error())
				return
			}

			got := reflect.ValueOf(r).Type()
			want := reflect.ValueOf(&api.Request{}).Type()
			if got != want {
				t.Error(
					"\nExpected", want,
					"\ngot", got,
				)
			}

			urlWant, err := api.NewUrl(confDefault.Api.Url)
			if err != nil {
				t.Error("Unexpected error:", err.Error())

			} else if err := urlWant.Join(path); err != nil {
				t.Error("Unexpected error:", err.Error())

			} else if urlWant.String() != r.URL.String() {
				t.Error(
					"\nExpected (url):", urlWant.String(),
					"\ngot:", r.URL.String(),
				)
			}

			if r.Method != method {
				t.Error(
					"\nExpected (method):", method,
					"\ngot:", r.Method,
				)
			}
		})
	}
}

func Test_api_markRequest(t *testing.T) {
	ctx := app.NewContext(confDefault)

	version, name := "0", "onesky-test"

	ctx.Build().ProductVersion = version
	ctx.Build().ProductName = name

	newApi, err := api.New(ctx)
	if err == nil {
		r, err := newApi.CreateRequest("GET", "/test")
		if err == nil {
			ua := api.NewRequestAgent(name, version, "")

			if r.Agent().String() != ua.String() {
				t.Error(
					"\nExpected:", ua.String(),
					"\ngot:     ", r.Agent().String(),
				)
			}

			if got := r.Header.Get("User-Agent"); got != ua.String() {
				t.Error(
					"\nExpected:", ua.String(),
					"\ngot:", got,
				)
			}
		}
	}

	if err != nil {
		t.Error(err.Error())
	}
}

func Test_api_authorizeRequest(t *testing.T) {

	setAuth := func(ctx app.Context, a app.Auth) app.Context {
		*(ctx.Auth()) = a
		return ctx
	}

	type args struct {
		ctx app.Context
	}

	tests := []struct {
		name string
		args args
		exp  api.RequestAuthorization
		pass bool
	}{
		{name: "OK", args: args{app.NewContext(confDefault)}, exp: api.NewRequestAuthorization(confDefault.Credentials.Token, confDefault.Credentials.Type), pass: true},
		{name: "No_auth", args: args{app.NewContext(confNoAuth)}, exp: nil, pass: false},
		{name: "Global", args: args{
			setAuth(app.NewContext(confDefault), app.Auth{Token: "someToken", Type: "Basic"}),
		}, exp: api.NewRequestAuthorization("someToken", "Basic"), pass: true},
		{name: "Override", args: args{
			setAuth(app.NewContext(confDefault), app.Auth{Token: "overrideToken", Type: "Basic"}),
		}, exp: api.NewRequestAuthorization("overrideToken", "Basic"), pass: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			newApi, err := api.New(tt.args.ctx)

			if err == nil {
				r, _ := newApi.CreateRequest("GET", "/test")

				if r.Auth() != nil && tt.pass {
					if r.Auth().String() != tt.exp.String() {
						t.Error(
							"\nExpected:", tt.exp.String(),
							"\ngot:", r.Auth().String(),
						)
					}
				}

			} else {
				t.Error(err)
			}
		})
	}
}
