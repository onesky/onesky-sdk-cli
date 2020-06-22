package command

import (
	"OneSky-cli/pkg/config"
)

type Command interface {
	Config() *config.OneskyConfig
}

type command struct {
	config *config.OneskyConfig
}

func New(config *config.OneskyConfig) Command {
	return &command{config}
}

func (c *command) Config() *config.OneskyConfig {
	return c.config
}
