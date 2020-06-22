package Auth

import (
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"fmt"
	"github.com/urfave/cli"
)

type Auth interface {
	Login(*cli.Context) error
	List(*cli.Context) error
	Revoke(*cli.Context) error
}

type auth struct {
	command.Command
}

func New(config *config.OneskyConfig) Auth {
	return &auth{
		command.New(config),
	}
}

func (a *auth) Login(c *cli.Context) (e error) {

	fmt.Println(c.String("access-token"))
	return e
}

func (a *auth) List(c *cli.Context) (e error) {
	fmt.Println(a.Config().Credentials.Token)
	return e

}

func (a *auth) Revoke(c *cli.Context) (e error) {
	return e

}
