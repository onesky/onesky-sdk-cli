package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type RequestParams map[string]interface{}

const TypeJson = "application/json"

type Request struct {
	*http.Request
	params *RequestParams
	agent  RequestAgent
	auth   RequestAuthorization
}

func NewRequest(p RequestParams) *Request {
	httpRequest, _ := http.NewRequest("", "", nil)
	return WrapHttpRequest(httpRequest, p)
}

func WrapHttpRequest(r *http.Request, p RequestParams) *Request {
	request := &Request{
		Request: r,
		agent:   &requestAgent{},
		auth:    &requestAuthorization{},
		params:  &RequestParams{},
	}

	if request.Request.Header == nil {
		request.Request.Header = http.Header{}
	}

	if p != nil {
		request.SetParams(p)
	} else {
		request.SetParams(RequestParams{})
	}

	request.Header.Add("Accept", TypeJson)
	request.Header.Add("Content-Type", TypeJson)

	request.upSyncHeaders()

	return request
}

func (r *Request) upSyncHeaders() *Request {
	if ua := r.Header.Get("User-Agent"); ua != r.agent.String() {
		r.agent = NewRequestAgentFromString(ua)
	}

	if au := r.Header.Get("Authorization"); au != r.auth.String() {
		r.auth = NewRequestAuthorizationFromString(au)
	}
	return r
}

func (r *Request) Agent() RequestAgent {
	return r.upSyncHeaders().agent
}

func (r *Request) SetAgent(agent RequestAgent) {
	r.Header.Add("User-Agent", agent.String())
	r.upSyncHeaders()
}

func (r *Request) Auth() RequestAuthorization {
	return r.auth
}

func (r *Request) SetAuth(auth RequestAuthorization) {
	r.Header.Add("Authorization", auth.String())
	r.upSyncHeaders()
}

func (r *Request) Params() RequestParams {
	return *r.params
}
func (r *Request) Param(key string) interface{} {
	if val, ok := (*r.params)[key]; ok {
		return val
	}
	return nil
}

func (r *Request) SetParams(params RequestParams) {
	*r.params = params
	r.saveParams()
}

func (r *Request) SetParam(key string, value interface{}) {
	(*r.params)[key] = value
	r.saveParams()
}

func (r *Request) saveParams() {

	var params string
	if r.Header.Get("Content-Type") == TypeJson {
		bParams, err := json.Marshal(*r.params)
		if err != nil {
			fmt.Println("Fatal:", err)
			os.Exit(1)
		}
		params = string(bParams)

	} else {
		uParams := url.Values{}
		for k, v := range *r.params {
			uParams.Set(k, fmt.Sprint(v))
		}
		params = uParams.Encode()
	}
	r.Body = ioutil.NopCloser(strings.NewReader(params))
}
