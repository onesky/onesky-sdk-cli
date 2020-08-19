package api

import (
	"OneSky-cli/pkg/app"
)

type Api interface {
	Client() Client
	CreateRequest(method, path string) (*Request, error)
}

type api struct {
	baseUrl *Url
	context app.Context
	client  Client
}

func New(context app.Context) (Api, error) {

	api := &api{
		context: context,
		client:  nil,
	}

	url, err := NewUrl(context.Config().Api.Url)
	if err == nil {
		api.SetBaseUrl(url)
	}

	return api, err
}

func (a *api) Client() Client {
	if a.client == nil {
		a.client = NewClient()
	}
	return a.client
}

func (a *api) CreateRequest(method, path string) (r *Request, err error) {

	url, err := NewUrl(a.BaseUrl().String())
	if err == nil {
		if err = url.Join(path); err == nil {
			r = NewRequest(nil)
			r.Method = method
			r.URL = url.URL

			a.markRequest(r)
			a.authorizeRequest(r)
		}
	}

	return r, err
}

func (a *api) BaseUrl() *Url {
	return a.baseUrl
}

func (a *api) SetBaseUrl(url *Url) {
	a.baseUrl = url
}

func (a *api) markRequest(r *Request) {
	if a.context.Build().ProductName+a.context.Build().ProductVersion != "" {
		r.SetAgent(NewRequestAgent(a.context.Build().ProductName, a.context.Build().ProductVersion, ""))
	}
}

func (a *api) authorizeRequest(r *Request) {
	if token := a.context.Auth().Token; token != "" {
		r.SetAuth(NewRequestAuthorization(token, a.context.Auth().Type))
	}
}
