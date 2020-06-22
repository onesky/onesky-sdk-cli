package config

type OneskyConfig struct {
	Title       string
	Credentials Credentials
}

type Credentials struct {
	Token string
}
