package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
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

func NewConfig(config string) *OneskyConfig {

	var Config *OneskyConfig
	if _, err := toml.Decode(config, &Config); err != nil {
		log.Fatalln(err)
	}

	return Config
}

func NewConfigFromFile(path string) *OneskyConfig {
	var Config *OneskyConfig
	if _, err := toml.DecodeFile(path, &Config); err != nil {
		log.Println(err)
	}
	Config.source = path

	return Config
}
