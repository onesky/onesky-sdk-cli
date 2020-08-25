package api

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestApi_NewRequest(t *testing.T) {

	params := RequestParams{"param1": "value1", "param2": "value2"}
	request := NewRequest(params)

	if request.Params() == nil {
		t.Error("Uninitialized parameters")
	}

	if !reflect.DeepEqual(request.Params(), params) {
		t.Error(
			"\nExpected:", params,
			"\ngot:", request.Params(),
		)
	}

	// check body
	buf := make([]byte, 1024)
	var bodyStrGot string
	for {
		n, err := request.Body.Read(buf)
		bodyStrGot += string(buf[:n])
		if err == io.EOF {
			break
		}
	}

	u := url.Values{}
	for i, k := range params {
		u.Add(i, fmt.Sprint(k))
	}
	bodyStrExp := u.Encode()

	if bodyStrGot != bodyStrExp {
		t.Error(
			"\nExpected:", bodyStrExp,
			"\ngot:", bodyStrGot,
		)
	}
}

func TestApi_WrapHttpRequest(t *testing.T) {
	httpReq := &http.Request{}
	params := RequestParams{"param1": "value1", "param2": "value2"}

	request := WrapHttpRequest(httpReq, params)

	if request.Params() == nil {
		t.Error("Uninitialized parameters")
	}

	if !reflect.DeepEqual(request.Params(), params) {
		t.Error(
			"\nExpected:", params,
			"\ngot:", request.Params(),
		)
	}

	if request.Header == nil {
		t.Error("Uninitialized header")
	}

	accExpected := TypeJson
	if got := request.Header.Get("Accept"); got != accExpected {
		t.Error(
			"\nExpected (accept):", accExpected,
			"\ngot:", got,
		)
	}

	conExpected := TypeJson
	if got := request.Header.Get("Content-Type"); got != conExpected {
		t.Error(
			"\nExpected (Content-Type):", conExpected,
			"\ngot:", got,
		)
	}
}

func TestRequest_Param(t *testing.T) {
	type fields struct {
		key   string
		value interface{}
	}
	tests := []struct {
		name   string
		fields fields
		exp    fields
		err    bool
	}{
		{"OK", fields{"key", "value"}, fields{"key", "value"}, true},
		{"OK", fields{}, fields{"anyKey", nil}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := NewRequest(nil)
			request.SetParam(tt.fields.key, tt.fields.value)

			want := tt.exp.value
			got := request.Param(tt.exp.key)
			if want != got && tt.err {
				t.Error(
					"\nExpected", want,
					"\ngot", got,
				)
			}
		})
	}
}

func TestRequest_Params(t *testing.T) {
	type fields struct {
		p RequestParams
	}
	tests := []struct {
		name   string
		fields fields
		err    bool
	}{
		{"Empty", fields{RequestParams{}}, false},
		{"OK", fields{RequestParams{"some": "param"}}, false},
		{"ERR", fields{nil}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := NewRequest(nil)
			request.SetParams(tt.fields.p)

			wantType := reflect.ValueOf(RequestParams{}).Type()
			gotType := reflect.ValueOf(request.Params()).Type()
			if wantType != gotType {
				t.Error(
					"\nExpected", wantType,
					"\ngot", gotType,
				)
			}

			wantVal := tt.fields.p
			gotVal := request.Params()
			if reflect.DeepEqual(wantVal, gotVal) && tt.err {
				t.Error(
					"\nExpected", wantVal,
					"\ngot", gotVal,
				)
			}
		})
	}
}
