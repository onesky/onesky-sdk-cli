package build

import (
	"OneSky-cli/pkg/app"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"time"
)

func CreateAppContext(c *cli.Context) (context app.Context, err error) {

	appContext := app.NewContext(nil)
	*appContext.Build() = app.Build{
		BuildId:        time.Now().Format("20060102-1504"),
		BuildInfo:      runtime.GOOS + "(" + runtime.GOARCH + ")",
		ConfigPath:     DefaultConfigPath,
		ProductName:    ProductName,
		ProductVersion: c.App.Version,
	}
	*(appContext.Flags()) = app.Flags{
		ConfigPath: c.String("config-file"),
		AuthString: c.String("access-token"),
		AuthType:   c.String("access-type"),
		Debug:      c.Bool("debug"),
	}

	*(appContext.Config()), err = buildConfig(c)

	*(appContext.Auth()) = buildAuth(c, appContext.Config())

	return appContext, err
}

func buildConfig(c *cli.Context) (conf app.Config, err error) {
	//var Config *config.Config

	// TRY TO LOAD ALTERNATIVE CONFIG
	currentConfigPath := c.String("config-file")
	if currentConfigPath != "" {
		if _, err = os.Stat(currentConfigPath); os.IsNotExist(err) {
			return conf, errors.New("Config not found in " + currentConfigPath)
		} else {
			conf, err = app.NewConfigFromFile(currentConfigPath)
		}

		// LOAD DEFAULT CONFIG
	} else {
		currentConfigPath = DefaultConfigPath

		// Create new default config file
		if _, confErr := os.Stat(currentConfigPath); os.IsNotExist(confErr) {
			fmt.Print("Initializing the configuration........", currentConfigPath)

			// Even if file can't be saved, we can go continue
			conf = DefaultConfig

			confErr = app.SaveConfig(currentConfigPath, &conf)
			if confErr != nil {
				fmt.Println("Fail\nWARNING:", confErr)
			} else {
				fmt.Println("OK")
			}

			// Load default config
		} else {
			conf, err = app.NewConfigFromFile(currentConfigPath)
		}
	}

	return conf, err
}

func buildAuth(c *cli.Context, conf *app.Config) app.Auth {
	// logic for auth of app-context
	token := c.String("config-file")
	tokenType := c.String("config-type")
	if token == "" {
		token = conf.Credentials.Token
		tokenType = conf.Credentials.Type
	}

	return app.Auth{
		Token: token,
		Type:  tokenType,
	}
}
