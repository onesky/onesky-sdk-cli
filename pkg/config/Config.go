package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
)

func SaveConfig(path string, config *OneskyConfig) {
	f, err := os.Create(path)

	if err != nil {
		fmt.Println(err)
	}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		fmt.Println(err)
	}
	if err := f.Close(); err != nil {
		fmt.Println(err)
	}
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

	return Config
}
