package build

import (
	"errors"
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
	"time"
)

func CreateAppContext(ctx *cli.Context, conf *app.Config) (context app.Context, err error) {

	appContext := app.NewContext(conf)
	*appContext.Build() = app.Build{
		BuildId:        time.Now().Format("20060102-1504"),
		BuildInfo:      runtime.GOOS + "(" + runtime.GOARCH + ")",
		ConfigPath:     DefaultConfigPath,
		ProductName:    ProductName,
		ProductVersion: ctx.App.Version,
	}
	*(appContext.Flags()) = app.Flags{
		ConfigPath: ctx.String("config-file"),
		AuthString: ctx.String("access-token"),
		AuthType:   ctx.String("access-type"),
		Debug:      ctx.Bool("debug"),
	}

	*(appContext.Auth()) = buildAuth(ctx, appContext.Config())

	return appContext, err
}

func CreateConfig(c *cli.Context) (conf app.Config, err error) {
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
			fmt.Print("Initializing the configuration........")

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
