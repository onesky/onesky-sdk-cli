package api

import (
	"strings"
)

type RequestAuthorizationType int

type RequestAuthorization interface {
	Type() string
	Value() string
	SetType(string)
	String() string
}

type requestAuthorization struct {
	authValue string
	authType  string
}

func NewRequestAuthorization(value, authType string) RequestAuthorization {

	if value == "" {
		panic("required RequestAuthorization.Value")
	}

	return &requestAuthorization{value, authType}
}

func NewRequestAuthorizationFromString(authString string) RequestAuthorization {
	parts := strings.SplitN(authString, " ", 2)

	pLen := len(parts)
	if authString != "" && pLen > 0 {
		authValue := authString
		authType := ""

		if pLen > 1 {
			if parts[0] != "" {
				authType = parts[0]
				authValue = authString[len(parts[0])+1:]
			}
		}
		return &requestAuthorization{authValue, authType}
	}

	panic("Invalid string for RequestAuthorization: " + authString)
}

func (r *requestAuthorization) Value() string {
	return r.authValue
}

func (r *requestAuthorization) Type() string {
	return r.authType
}

func (r *requestAuthorization) SetType(s string) {
	r.authType = s
}

func (r *requestAuthorization) String() string {
	authString := r.authValue
	if r.Type() != "" {
		authString = r.Type() + " " + r.Value()
	}

	return authString
}
