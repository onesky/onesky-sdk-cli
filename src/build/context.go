package build

import (
	"OneSky-cli/pkg/config"
	. "OneSky-cli/pkg/context"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"time"
)

func CreateAppContext(c *cli.Context) (context AppContext, err error) {

	appContext := New(nil)
	*appContext.Build() = AppBuild{
		BuildId:        time.Now().Format("20060102-1504"),
		BuildInfo:      runtime.GOOS + "(" + runtime.GOARCH + ")",
		ConfigPath:     DefaultConfigPath,
		ProductName:    ProductName,
		ProductVersion: c.App.Version,
	}
	*(appContext.Flags()) = AppFlags{
		ConfigPath: c.String("config-file"),
		AuthString: c.String("access-token"),
		AuthType:   c.String("access-type"),
		Debug:      c.Bool("debug"),
	}

	/////////////////////////////////////// CONFIG ///////////////////////////////////
	//var Config *config.OneskyConfig

	// TRY TO LOAD ALTERNATIVE CONFIG
	currentConfigPath := appContext.Flags().ConfigPath
	if currentConfigPath != "" {
		if _, err = os.Stat(currentConfigPath); os.IsNotExist(err) {
			return context, errors.New("Config not found in " + currentConfigPath)
		} else {
			*appContext.Config(), err = config.NewConfigFromFile(currentConfigPath)
		}

		// LOAD DEFAULT CONFIG
	} else {
		currentConfigPath = DefaultConfigPath

		// Create new default config file
		if _, confErr := os.Stat(currentConfigPath); os.IsNotExist(confErr) {
			fmt.Print("Initializing to: ", currentConfigPath)

			// Even if file can't be saved, we can go continue
			*appContext.Config() = DefaultConfig

			confErr = config.SaveConfig(currentConfigPath, appContext.Config())
			if confErr != nil {
				fmt.Println("\nWARNING:", confErr)
			} else {
				fmt.Println(".......... OK")
			}

			// Load default config
		} else {
			*appContext.Config(), err = config.NewConfigFromFile(currentConfigPath)
		}
	}
	/////////////////////////////////////// end of CONFIG ///////////////////////////////////

	return appContext, err
}
