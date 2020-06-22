package config

type OneskyConfig struct {
	source      string
	Title       string
	Credentials Credentials
}

type Credentials struct {
	Token string
}

func (o *OneskyConfig) Update() {
	if o.source != "" {
		SaveConfig(o.source, o)
	}
}
