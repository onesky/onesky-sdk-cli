package main

import (
	"OneSky-cli/pkg/command/Auth"
	"OneSky-cli/pkg/command/File"
	"OneSky-cli/pkg/command/Lang"
	"OneSky-cli/pkg/config"
	"OneSky-cli/pkg/help"
	"OneSky-cli/src/build"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"runtime"
	"syscall"
	"time"
)

var onTerminate = func(code syscall.Signal) {}

const API_URL = "https://management-api.onesky.app/v1"

func init() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//SIG_INT(&onTerminate)

}

func main() {

	///////////////////////////////////
	// LOAD CONFIG
	///////////////////////////////////
	var Config *config.OneskyConfig
	if _, confErr := os.Stat(build.CONFIG_PATH); os.IsNotExist(confErr) {
		fmt.Println("Trying to save new config file to:", build.CONFIG_PATH)

		Config = &config.OneskyConfig{
			Title: "Onesky config",
			Api:   config.Api{Url: API_URL},
		}

		confErr = config.SaveConfig(build.CONFIG_PATH, Config)
		if confErr != nil {
			fmt.Println("WARNING:", confErr)
		}

	} else {
		Config = config.NewConfigFromFile(build.CONFIG_PATH)
	}

	/////////////////////////////////
	// CLI-INTERFACE
	////////////////////////////////
	fmt.Println("Build: ", time.Now().Format("20060102-1504"), runtime.GOOS, runtime.GOARCH)
	fmt.Println("Read config:", build.CONFIG_PATH)

	cli.AppHelpTemplate = help.AppHelpTemplate
	cli.CommandHelpTemplate = help.CommandHelpTemplate
	cli.SubcommandHelpTemplate = help.SubcommandHelpTemplate
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Display information about built-in commands",
	}

	Cli := &cli.App{
		Name:     "onesky",
		Version:  "0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "OneSky Inc.",
				Email: "https://www.onesky.app/",
			},
		},
		Copyright: "(c) 2019 OneSky Inc. All rights reserved.",
		//HelpName: "onesky",
		HideHelpCommand: true,
		Usage:           "OneSky SDK: Command-line Tool",
		ArgsUsage:       "subcommand [subcommand options and args]",
		Description: "OneSky CLI manages authentication, localization end-to-end workflow," +
			"and interactions with OneSky APIs. With the OneSky command-line tool, " +
			"it’s easy to perform many common localization tasks like uploading and downloading string files," +
			"either from the command line or in scripts and other automation.",
		Commands: []*cli.Command{
			// AUTH
			&cli.Command{
				Name:            "auth",
				Usage:           "Manage credentials for the OneSky CLI",
				UsageText:       "onesky [global options] auth <command> [options]",
				Description:     "Manage credentials for the OneSky CLI",
				ArgsUsage:       "{login|list|revoke}",
				HideHelpCommand: true,

				Subcommands: []*cli.Command{
					//Auth.GetLoginCmd(),
					&cli.Command{
						Name:        "login",
						Action:      Auth.New(Config).Login,
						Description: "Authorize to access the OneSky API with access token",
						Usage:       "Authorize to access the OneSky API with access token",
						UsageText:   "onesky auth login --access-token=ACCESS_TOKEN",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "access-token",
								Usage:    "Set `ACCESS_TOKEN`",
								Required: true,
							},
						},
					},
					&cli.Command{
						Name:        "list",
						Action:      Auth.New(Config).List,
						Description: "List credentialed access token",
						Usage:       "List credentialed access token",
						UsageText:   "onesky auth list [options]",
					},
					&cli.Command{
						Name:        "revoke",
						Action:      Auth.New(Config).Revoke,
						Description: "Revoke access credential",
						Usage:       "Revoke access credential",
						UsageText:   "onesky auth revoke [options]",
					},
				},

				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
			},
			// LANG
			&cli.Command{
				Name:            "lang",
				Usage:           "Manage languages of the app",
				UsageText:       "onesky [global options] lang <command> [options]",
				Description:     "Manage languages of the app",
				HideHelpCommand: true,
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:        "list",
						Aliases:     []string{"l"},
						Action:      Lang.New(Config).List,
						Description: "List all enabled languages of the app",
						Usage:       "List all enabled languages of the app",
						UsageText:   "onesky lang list",
					},
				},
			},
			// FILE
			&cli.Command{
				Name:            "file",
				Usage:           "onesky file upload --platform-id=web --language-id=en_US --file-name=en_US.json --content=’{“apple-key”: “Apple”}’",
				UsageText:       "onesky [global options] file <command> [options]",
				Description:     "Manage files of the app",
				HideHelpCommand: true,
				Subcommands: []*cli.Command{
					&cli.Command{
						Name:        "list",
						Action:      File.New(Config).List,
						Aliases:     []string{"l"},
						Description: "List all files of the app",
						Usage:       "List all files of the app",
						UsageText:   "`onesky file list`",
					},
					&cli.Command{
						Name:        "upload",
						Action:      File.New(Config).Upload,
						Aliases:     []string{"u"},
						Description: "Upload a new file to the app",
						Usage:       "Upload a new file to the app",
						UsageText:   "onesky file upload --platform-id=PLATFORM_ID --language-id=LANGUAGE_ID --file-name=FILE_NAME --content=CONTENT_ENCODED_IN_UTF8 [--plugin-agent=intellij]",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "platform-id",
								Aliases: []string{"p"},
								Usage:   "`PLATFORM_ID` - one of 'web', 'ios' or 'android' (Ex.: --platform-id=web)",
							},
							&cli.StringFlag{
								Name:    "language-id",
								Aliases: []string{"l"},
								Usage:   "`LANGUAGE_ID` (Ex.: --language-id=en_US)",
							},
							&cli.StringFlag{
								Name:    "file-name",
								Aliases: []string{"f"},
								Usage:   "`FILE_NAME`",
							},
							&cli.StringFlag{
								Name:  "content",
								Usage: "​`CONTENT_ENCODED_IN_UTF8`",
							},
							&cli.StringFlag{
								Name:  "plugin-agent",
								Usage: "​​`USER_AGENT_HEADER_COMMENT` (Ex.: --plugin-agent=intellij)",
							},
						},
					},
					&cli.Command{
						Name:        "download",
						Aliases:     []string{"d"},
						Action:      File.New(Config).Download,
						Description: "Get file content by id",
						Usage:       "Get file content by id",
						UsageText:   "onesky file download --file-id=FILE_ID [--plugin-agent=USER_AGENT_HEADER_COMMENT]",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "file-id",
								Aliases: []string{"f"},
								Usage:   "​`FILE_ID`",
							},
							&cli.StringFlag{
								Name: "plugin-agent",
								//Aliases: []string{"a"},
								Usage: "​​`USER_AGENT_HEADER_COMMENT` (Ex.: --plugin-agent=intellij)",
							},
						},
					},
				},

				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "access-token",
				Usage:    "Set `ACCESS_TOKEN`",
				Required: false,
			},
			&cli.BoolFlag{
				Name:     "debug",
				Usage:    "Show debug information",
				Required: false,
			},
		},
		EnableBashCompletion: true,
		CommandNotFound: func(c *cli.Context, command string) {
			fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
			return nil
		},
		Before: func(c *cli.Context) error {

			if tokenString := c.String("access-token"); tokenString != "" {
				Config.Credentials.Token = tokenString
			}
			//fmt.Println("Access-token:", Config.Credentials.Token)

			return nil
		},
		After: func(c *cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			return cli.ShowAppHelp(c)
		},
	}

	if err := Cli.Run(os.Args); err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}
