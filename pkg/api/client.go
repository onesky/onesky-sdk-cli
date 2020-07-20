package api

import (
	"crypto/tls"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"time"
)

type Client interface {
	DoRequest(request *Request, debug bool) ([]byte, error)
	SetTimeout(timeout int)
	Timeout() int
}

type client struct {
	defaultTimeout int
}

func NewClient() Client {
	return &client{30}
}

func (c *client) SetTimeout(timeout int) {
	c.defaultTimeout = timeout
}

func (c *client) Timeout() int {
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
	//httpClient := &http.Client{}
	if c.Timeout() > 0 {
		httpClient.Timeout = time.Duration(c.Timeout()) * time.Second
	}

	var response *http.Response
	response, err = httpClient.Do(request.Request)

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
