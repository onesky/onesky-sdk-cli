package main

import (
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/build"
	"github.com/onesky/onesky-sdk-cli/src/command/Api"
	"github.com/onesky/onesky-sdk-cli/src/command/Auth"
	"github.com/onesky/onesky-sdk-cli/src/command/File"
	"github.com/onesky/onesky-sdk-cli/src/command/Lang"
	"github.com/onesky/onesky-sdk-cli/src/help"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
	"time"
)

//var onTerminate = func(code syscall.Signal) {}

func init() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	//SIG_INT(&onTerminate)

}

func main() {

	/////////////////////////////////
	// CLI-INTERFACE
	////////////////////////////////
	cli.AppHelpTemplate = help.AppHelpTemplate
	cli.CommandHelpTemplate = help.CommandHelpTemplate
	cli.SubcommandHelpTemplate = help.SubcommandHelpTemplate
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Display information about built-in commands",
	}
	cli.VersionPrinter = func(c *cli.Context) {
		_, _ = fmt.Fprintln(c.App.Writer, c.App.Version)
	}

	Cli := &cli.App{
		Name:     "onesky",
		Version:  "0.0.2",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
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
			"it's easy to perform many common localization tasks like uploading and downloading string files," +
			"either from the command line or in scripts and other automation.",

		Commands: []*cli.Command{
			Auth.Command,
			Lang.Command,
			File.Command,
			Api.Command,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "access-token",
				Usage:    "Set global `TOKEN` (access-token from config will be ignored)",
				Required: false,
			},
			&cli.StringFlag{
				Name:     "access-type",
				Usage:    "Set global authorization `TYPE` (access-type from config will be ignored)",
				Required: false,
				Value:    "Bearer",
			},
			&cli.StringFlag{
				Name:     "config-file",
				Usage:    "Set alternative config `PATH`",
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
			_, _ = fmt.Fprintf(c.App.Writer, "Thar be no %q here.\n", command)
		},
		OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
			if isSubcommand {
				return err
			}

			_, _ = fmt.Fprintf(c.App.Writer, "WRONG: %#v\n", err)
			return nil
		},
		Before: func(c *cli.Context) error {

			c.App.Setup()
			config, err := build.CreateConfig(c)
			if err != nil {
				return err
			}
			context, err := build.CreateAppContext(c, &config)
			if err != nil {
				return err
			}

			if context.Flags().Debug {
				_, _ = fmt.Fprintf(c.App.Writer, "Build: %s / %s\n", context.Build().BuildId, context.Build().BuildInfo)
				_, _ = fmt.Fprintf(c.App.Writer, "Loaded config file: %s\n", context.Config().Source())
			}

			c.App.Metadata["context"] = context

			//*Config = *context.Config()

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
