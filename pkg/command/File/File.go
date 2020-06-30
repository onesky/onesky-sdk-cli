package File

import (
	"OneSky-cli/pkg/api"
	"OneSky-cli/pkg/command"
	"OneSky-cli/pkg/config"
	"fmt"
	"github.com/urfave/cli"
)

type File interface {
	Upload(*cli.Context) error
	List(*cli.Context) error
	Download(*cli.Context) error
}

type file struct {
	command.Command
	api api.Api
}

func New(config *config.OneskyConfig) File {
	return &file{
		Command: command.New(config),
		api:     api.New(config),
	}
}

func (f *file) Upload(c *cli.Context) (err error) {

	request, err := f.api.NewApiRequest("POST", "/files")
	if err == nil {
		request.SetParam("platformId", c.String("platform-id"))
		request.SetParam("languageId", c.String("language-id"))
		request.SetParam("fileName", c.String("file-name"))
		request.SetParam("content", c.String("content"))

		isDebug := c.Bool("debug")
		responseString, e := f.api.Client().DoRequest(request, isDebug)
		if e == nil && !isDebug {
			fmt.Println(string(responseString))
		}
	}

	return err
}

func (f *file) List(c *cli.Context) (err error) {

	request, err := f.api.NewApiRequest("GET", "/files")
	if err == nil {
		isDebug := c.Bool("debug")
		responseString, e := f.api.Client().DoRequest(request, isDebug)
		if e == nil && !isDebug {
			fmt.Println(string(responseString))
		}
	}

	return err
}

func (f *file) Download(c *cli.Context) (e error) {

	path := fmt.Sprintf("/files/%s/contents", c.String("file-id"))
	request, err := f.api.NewApiRequest("GET", path)
	if err == nil {
		isDebug := c.Bool("debug")
		responseString, e := f.api.Client().DoRequest(request, isDebug)
		if e == nil && !isDebug {
			fmt.Println(string(responseString))
		}
	}

	return err
}
