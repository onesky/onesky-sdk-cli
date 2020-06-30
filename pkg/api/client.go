package api

import (
	"OneSky-cli/pkg/config"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type Client interface {
	DoRequest(request *Request, debug bool) ([]byte, error)
}

type client struct {
	config *config.OneskyConfig
}

func newClient(config *config.OneskyConfig) Client {
	return &client{config}
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

	httpClient := &http.Client{}
	if c.config.Api.Timeout > 0 {
		httpClient.Timeout = time.Duration(c.config.Api.Timeout) * time.Second
	}

	response, err := httpClient.Do(request.Request)
	if err == nil {

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
		}
		data, err = ioutil.ReadAll(response.Body)

	}
	return data, err
}
