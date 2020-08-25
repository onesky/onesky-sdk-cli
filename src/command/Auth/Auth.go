package Auth

import (
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/urfave/cli/v2"
)

const ErrNoToken = "Token not found"
const ErrUpdate = "Unable to update config: %s"
const MsgNewToken = "New token: (%s) %s"
const MsgGetToken = "Access token: (%s) %s"
const MsgRevokeToken = "Access token was revoked"

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
		fmt.Printf(MsgNewToken, a.Config().Credentials.Type, a.Config().Credentials.Token)
		fmt.Println()
	} else {
		fmt.Printf(ErrUpdate, e.Error())
		fmt.Println()
	}
	return e
}

func List(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	if tok := a.Config().Credentials.Token; tok != "" {
		fmt.Printf(MsgGetToken, a.Config().Credentials.Type, a.Config().Credentials.Token)
		fmt.Println()
	} else {
		fmt.Println(ErrNoToken)
	}
	return e

}

func Revoke(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	a.Config().Credentials.Token = ""

	e = a.Config().Update()
	if e == nil {
		fmt.Print(MsgRevokeToken)
		fmt.Println()
	} else {
		fmt.Printf(ErrUpdate, e.Error())
		fmt.Println()
	}

	return e

}
