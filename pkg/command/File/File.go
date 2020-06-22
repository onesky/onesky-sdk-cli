package File

import (
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"github.com/urfave/cli"
)

type File interface {
	Upload(*cli.Context) error
	List(*cli.Context) error
	Download(*cli.Context) error
}

type file struct {
	command.Command
}

func New(config *config.OneskyConfig) File {
	return &file{
		command.New(config),
	}
}

func (f *file) Upload(c *cli.Context) (e error) {

	return e
}

func (f *file) List(c *cli.Context) (e error) {
	return e

}

func (f *file) Download(c *cli.Context) (e error) {
	return e

}
