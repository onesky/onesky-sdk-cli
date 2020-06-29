package Lang

import (
	"OneSky-cli/pkg/api"
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"fmt"
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

func (l *lang) List(c *cli.Context) (err error) {

	apiClient := api.New(l.Config())
	request, err := apiClient.NewApiRequest("GET", "/languages")
	if err == nil {
		isDebug := c.Bool("debug")
		responseString, e := apiClient.Client().DoRequest(request, isDebug)
		if e == nil && !isDebug {
			fmt.Println(string(responseString))
		}
	}

	return err
}
