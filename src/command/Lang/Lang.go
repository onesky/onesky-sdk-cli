package Lang

import (
	"fmt"
	. "github.com/onesky/onesky-sdk-cli/src/api"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/urfave/cli/v2"
)

type Lang interface {
	List(*cli.Context) error
}

func List(c *cli.Context) (err error) {
	ctx := c.App.Metadata["context"].(app.Context)

	api, err := New(ctx)
	if err == nil {
		request, err := api.CreateRequest("GET", "/languages")
		if err == nil {

			responseString, e := api.Client().DoRequest(request, ctx.Flags().Debug)
			if e == nil && !ctx.Flags().Debug {
				fmt.Println(string(responseString))
			}
		}
	}

	return err
}
