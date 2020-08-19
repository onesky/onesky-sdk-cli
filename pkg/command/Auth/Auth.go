package Auth

import (
	"fmt"
	"github.com/urfave/cli"
	"onesky-sdk-cli/pkg/app"
)

type Auth interface {
	Login(*cli.Context) error
	List(*cli.Context) error
	Revoke(*cli.Context) error
}

func Login(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	a.Config().Credentials.Token = c.String("access-token")
	a.Config().Credentials.Type = c.String("access-type")

	e = a.Config().Update()
	if e == nil {
		fmt.Printf("NewContext token: (%s) %s\n", a.Config().Credentials.Type, a.Config().Credentials.Token)
	} else {
		fmt.Println("Unable to update config: ", e.Error())
	}
	return e
}

func List(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	if tok := a.Config().Credentials.Token; tok != "" {
		fmt.Printf("Access token: (%s) %s\n", a.Config().Credentials.Type, a.Config().Credentials.Token)
	} else {
		fmt.Println("Token not found")
	}
	return e

}

func Revoke(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	a.Config().Credentials.Token = ""

	e = a.Config().Update()
	if e == nil {
		fmt.Println("Access token was revoked")
	} else {
		fmt.Println("Unable to update config: ", e.Error())
	}

	return e

}
