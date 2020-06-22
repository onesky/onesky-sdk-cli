package Lang

import (
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"github.com/urfave/cli"
)

type Lang interface {
	List(*cli.Context) error
}

type lang struct {
	command.Command
}

func New(config *config.OneskyConfig) Lang {
	return &lang{
		command.New(config),
	}
}

func (l *lang) List(c *cli.Context) (e error) {
	return e

}
