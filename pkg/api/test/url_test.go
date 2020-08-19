package test

import (
	"fmt"
	"onesky-sdk-cli/pkg/api"
	"testing"
)

type TestPair struct {
	given    interface{}
	expected interface{}
	err      bool
}

func TestNewUrl(t *testing.T) {
	tests := []TestPair{
		{"some.url/with/slashes/", "some.url/with/slashes", true},
		{"http://some.url/with/slashes/", "http://some.url/with/slashes", false},
		{"some.url/without/slashes", "some.url/without/slashes", true},
		{"http://some.url/without/slashes", "http://some.url/without/slashes", false},
	}

	for i := range tests {

		url, err := api.NewUrl(fmt.Sprint(tests[i].given))

		if tests[i].err == true && err == nil {
			t.Error(
				"\nFor", tests[i],
				"\nexpected", "ERROR",
				"\ngot", url.String(),
			)

		} else if tests[i].err == false {
			if err != nil {
				t.Error(err)
			}

			if url.String() != tests[i].expected {
				t.Error(
					"\nFor", tests[i],
					"\nexpected", tests[i].expected,
					"\ngot", url.String(),
				)
			}
		}

	}
}

func TestUrl_Join(t *testing.T) {
	baseUrl := "http://some.url/with/slashes"
	tests := []TestPair{
		{"/join/with/slashes/", baseUrl + "/join/with/slashes/", false},
		{"/join/with/slashes", baseUrl + "/join/with/slashes", false},
		{"join/without/slashes/", baseUrl + "/join/without/slashes/", false},
		{"join/without/slashes", baseUrl + "/join/without/slashes", false},
	}

	for i := range tests {
		url, _ := api.NewUrl(baseUrl)
		err := url.Join(fmt.Sprint(tests[i].given))

		if !tests[i].err && err != nil {
			t.Error("\nUnexpected error: ", err)

		} else if url.String() != tests[i].expected {
			t.Error(
				"\nFor", tests[i],
				"\nexpected", tests[i].expected,
				"\ngot", url.String(),
			)
		}
	}
}
