package Api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/urfave/cli/v2"
)

type Api interface {
	List(*cli.Context) error
	Set(*cli.Context) error
}

func List(c *cli.Context) (err error) {
	a := c.App.Metadata["context"].(app.Context)

	for k, v := range structs.Map(a.Config().Api) {
		fmt.Println(k+":", v)
	}
	return err
}

func Set(c *cli.Context) (e error) {
	a := c.App.Metadata["context"].(app.Context)

	if baseUrl := c.String("url"); baseUrl != "" {
		a.Config().Api.Url = baseUrl
	}

	if timeout := c.Int("timeout"); timeout > -1 {
		a.Config().Api.Timeout = timeout
	}

	e = a.Config().Update()
	if e == nil {
		fmt.Println("Successful!")
	} else {
		fmt.Println("Unable to update config: ", e.Error())
	}
	return e
}
