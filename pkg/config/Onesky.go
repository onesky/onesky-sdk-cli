package config

import (
	"errors"
)

type OneskyConfig struct {
	source      string
	Title       string
	Credentials Credentials
	Api         Api
}

func (o *OneskyConfig) Update() error {
	if o.source != "" {
		return SaveConfig(o.source, o)
	}
	return errors.New("unknown source")
}

func (o *OneskyConfig) Source() string {
	return o.source
}

type Credentials struct {
	Token string
	Type  string // Bearer, Basic and etc
}

type Api struct {
	Url     string
	Timeout int
}
