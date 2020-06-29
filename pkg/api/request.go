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
}

func newApiRequest(r *http.Request) *Request {
	request := &Request{
		Request: r,
		params:  &RequestParams{},
	}

	request.Header.Add("Accept", TypeJson)
	request.Header.Add("Content-Type", TypeJson)

	return request
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
			fmt.Println(err)
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
