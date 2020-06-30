package Api

import (
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"fmt"
	"github.com/fatih/structs"
	"github.com/urfave/cli"
)

type Api interface {
	List(*cli.Context) error
	Set(*cli.Context) error
}

type api struct {
	command.Command
}

func New(config *config.OneskyConfig) Api {
	return &api{
		command.New(config),
	}
}

func (a *api) List(c *cli.Context) (err error) {

	for k, v := range structs.Map(a.Config().Api) {
		fmt.Println(k+":", v)
	}
	return err
}

func (a *api) Set(c *cli.Context) (e error) {

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
