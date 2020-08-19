package File

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io/ioutil"
	. "onesky-sdk-cli/pkg/api"
	"onesky-sdk-cli/pkg/app"
)

type File interface {
	Upload(*cli.Context) error
	List(*cli.Context) error
	Download(*cli.Context) error
}

func Upload(c *cli.Context) (err error) {
	appContext := c.App.Metadata["context"].(app.Context)
	api, err := New(appContext)
	if err == nil {

		request, err := api.CreateRequest("POST", "/files")
		if err == nil {
			request.SetParam("platformId", c.String(FlagPlatformId.Name))
			request.SetParam("languageId", c.String(FlagLanguageId.Name))
			request.SetParam("fileName", c.String(FlagFileName.Name))

			if pa := c.String(FlagPluginAgent.Name); pa != "" {
				request.Agent().SetPlugin(pa)
			}

			content := c.String(FlagFileContent.Name)
			path := c.String(FlagFilePath.Name)
			if content != "" && path != "" {
				return errors.New("incongruous options --content and --path")
			} else if path != "" {
				if byteContent, err := ioutil.ReadFile(path); err == nil {
					content = string(byteContent)

					// Use it if something go wrong
					//import "golang.org/x/exp/utf8string"
					//content = utf8string.NewString( string(byteContent) ).String()
				}
			}

			request.SetParam("content", content)

			responseString, err := api.Client().DoRequest(request, appContext.Flags().Debug)

			if err == nil && !appContext.Flags().Debug {
				fmt.Println(string(responseString))
			}
		}
	}

	return err
}

func List(c *cli.Context) (err error) {
	appContext := c.App.Metadata["context"].(app.Context)

	var api Api
	api, err = New(appContext)

	if err == nil {

		var request *Request
		request, err = api.CreateRequest("GET", "/files")

		if err == nil {

			var responseBytes []byte
			responseBytes, err = api.Client().DoRequest(request, appContext.Flags().Debug)

			if err == nil && !appContext.Flags().Debug {
				fmt.Println(string(responseBytes))
			}
		}
	}

	return err
}

func Download(c *cli.Context) (e error) {
	appContext := c.App.Metadata["context"].(app.Context)
	api, err := New(appContext)
	if err == nil {

		path := fmt.Sprintf("/files/%s/contents", c.String(FlagFileId.Name))
		request, err := api.CreateRequest("GET", path)
		if err == nil {

			if pa := c.String(FlagPluginAgent.Name); pa != "" {
				request.Agent().SetPlugin(pa)
			}

			responseString, e := api.Client().DoRequest(request, appContext.Flags().Debug)
			if e == nil {

				if savePath := c.String(FlagOutput.Name); savePath != "" {
					return ioutil.WriteFile(savePath, responseString, 0660)
				}

				if !appContext.Flags().Debug {
					fmt.Println(string(responseString))
				}
			}
		}
	}

	return err
}
