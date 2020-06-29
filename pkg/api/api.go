package api

import (
	"OneSky-cli/pkg/config"
	"net/http"
)

type Api interface {
	Client() Client
	NewApiRequest(method, path string) (*Request, error)
}

type api struct {
	config *config.OneskyConfig
	client Client
}

func New(config *config.OneskyConfig) Api {
	return &api{
		config: config,
		client: nil,
	}
}

func (a *api) Client() Client {
	if a.client == nil {
		a.client = newClient(a.config)
	}
	return a.client
}

func (a *api) NewApiRequest(method, path string) (r *Request, err error) {

	endpointUrl, err := NewUrl(a.config.Api.Url)
	if err != nil {
		return r, err
	}

	if err = endpointUrl.Join(path); err == nil {
		httpRequest := &http.Request{}
		a.authorizeHttpRequest(httpRequest)
		r = newApiRequest(httpRequest)
		r.Method = method
		r.URL = endpointUrl.URL
	}

	return r, err
}

func (a *api) authorizeHttpRequest(request *http.Request) {
	request.Header = http.Header{}
	if a.config.Credentials.Type == "" {
		request.Header.Add("authorization", a.config.Credentials.Token)
	} else {
		request.Header.Add("authorization", a.config.Credentials.Type+" "+a.config.Credentials.Token)
	}
}