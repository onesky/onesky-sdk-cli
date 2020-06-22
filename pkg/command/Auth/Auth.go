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

	a.Config().Credentials.Token = c.String("access-token")
	a.Config().Update()
	fmt.Println("New token: ", a.Config().Credentials.Token)
	return e
}

func (a *auth) List(c *cli.Context) (e error) {
	if tok := a.Config().Credentials.Token; tok != "" {
		fmt.Println("Access token:", a.Config().Credentials.Token)
	} else {
		fmt.Println("Token not found")
	}
	return e

}

func (a *auth) Revoke(c *cli.Context) (e error) {
	a.Config().Credentials.Token = ""
	a.Config().Update()
	fmt.Println("Access token was revoked")
	return e

}
