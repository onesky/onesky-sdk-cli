package app

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type Config struct {
	source      string
	Title       string
	Credentials Credentials
	Api         Api
}

type Credentials struct {
	Token string
	Type  string // Bearer, Basic and etc
}

type Api struct {
	Url     string
	Timeout int
}

func (o *Config) Update() error {
	if o.source != "" {
		return SaveConfig(o.source, o)
	}
	return errors.New("unknown source")
}

func (o *Config) Source() string {
	return o.source
}

func SaveConfig(path string, config *Config) (err error) {

	baseDir := filepath.Dir(path)
	if _, err = os.Stat(baseDir); os.IsNotExist(err) {
		err = os.MkdirAll(baseDir, 0660)
	}

	if err == nil {
		f, err := os.Create(path)

		if err != nil {
			fmt.Println(err)
		}

		config.source = path
		if err := toml.NewEncoder(f).Encode(config); err != nil {
			fmt.Println(err)
		}
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}

	return err
}

//func NewConfig(config string) (Config, error) {
//
//	var Config Config
//	_, err := toml.Decode(config, &Config)
//	if err == nil {
//	Config.source = path
//	}
//
//	return Config, err
//}

func NewConfigFromFile(path string) (Config, error) {
	var Config Config
	_, err := toml.DecodeFile(path, &Config)
	if err == nil {
		Config.source = path
	}

	return Config, err
}
