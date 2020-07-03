package api

import "strings"

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
	parts := strings.Fields(authString)

	switch len(parts) {
	case 0:
		{ // name / version
			panic("Invalid string for RequestAuthorization: " + authString)
		}
	case 1:
		{ //name / version plugin
			return &requestAuthorization{parts[0], ""}
		}
	default:
		return &requestAuthorization{strings.Join(parts[1:], " "), parts[0]}
	}
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
