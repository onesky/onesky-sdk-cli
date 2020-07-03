package api

import (
	"OneSky-cli/pkg/context"
)

type Api interface {
	Client() Client
	CreateRequest(method, path string) (*Request, error)
}

type api struct {
	baseUrl *Url
	context context.AppContext
	client  Client
}

func New(context context.AppContext) (Api, error) {

	api := &api{
		context: context,
		client:  nil,
	}

	url, err := NewUrl(context.Config().Api.Url)
	if err == nil {
		api.baseUrl = url
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

	if err = a.baseUrl.Join(path); err == nil {
		r = NewRequest(nil)
		r.SetAgent(NewRequestAgent(a.context.Build().ProductName, a.context.Build().ProductVersion, ""))

		if token := a.context.Flags().AuthString; token != "" {
			r.SetAuth(NewRequestAuthorization(token, a.context.Flags().AuthType))

		} else if token = a.context.Config().Credentials.Token; token != "" {
			r.SetAuth(NewRequestAuthorization(token, a.context.Config().Credentials.Type))
		}

		r.Method = method
		r.URL = a.baseUrl.URL
	}

	return r, err
}
