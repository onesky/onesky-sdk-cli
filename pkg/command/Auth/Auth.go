package Auth

import (
	"OneSky-cli/pkg/context"
	"fmt"
	"github.com/urfave/cli"
)

type Auth interface {
	Login(*cli.Context) error
	List(*cli.Context) error
	Revoke(*cli.Context) error
}

func Login(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(context.AppContext)

	a.Config().Credentials.Token = c.String("access-token")

	if authType := c.String("access-type"); authType != "" {
		a.Config().Credentials.Type = authType
	}

	e = a.Config().Update()
	if e == nil {
		fmt.Println("New token: ", a.Config().Credentials.Token)
	} else {
		fmt.Println("Unable to update config: ", e.Error())
	}
	return e
}

func List(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(context.AppContext)

	if tok := a.Config().Credentials.Token; tok != "" {
		fmt.Println("Access token:", a.Config().Credentials.Token)
	} else {
		fmt.Println("Token not found")
	}
	return e

}

func Revoke(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(context.AppContext)

	a.Config().Credentials.Token = ""

	e = a.Config().Update()
	if e == nil {
		fmt.Println("Access token was revoked")
	} else {
		fmt.Println("Unable to update config: ", e.Error())
	}

	return e

}
