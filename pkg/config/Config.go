package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
)

func SaveConfig(path string, config *OneskyConfig) (err error) {
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

	return err
}

//func NewConfig(config string) (OneskyConfig, error) {
//
//	var Config OneskyConfig
//	_, err := toml.Decode(config, &Config)
//	if err == nil {
//	Config.source = path
//	}
//
//	return Config, err
//}

func NewConfigFromFile(path string) (OneskyConfig, error) {
	var Config OneskyConfig
	_, err := toml.DecodeFile(path, &Config)
	if err == nil {
		Config.source = path
	}

	return Config, err
}
