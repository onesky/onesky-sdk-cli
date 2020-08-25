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

	//var Config = &app.Config{}
	fmt.Println(os.Args)
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
			// AUTH
			{
				Name:      "auth",
				Usage:     "Manage credentials for the OneSky CLI",
				UsageText: "onesky [global options] auth <command> [options]",
				Description: "Authorize to access the OneSky API with access token: \n" +
					"			onesky auth login --access-token=some-token-string --access-type=Bearer\n\n" +
					"   List credentialed access token: \n" +
					"			onesky auth list\n\n" +
					"	Revoke access credential: \n" +
					"			onesky auth revoke\n",
				ArgsUsage:       "{login|list|revoke}",
				HideHelpCommand: true,

				Subcommands: []*cli.Command{
					//Auth.GetLoginCmd(),
					{
						Name:        "login",
						Action:      Auth.Login,
						Description: "Authorize to access the OneSky API with access token",
						Usage:       "Authorize to access the OneSky API with access token",
						UsageText:   "onesky auth login --access-token=ACCESS_TOKEN [--access-type=TYPE]",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "access-token",
								Usage:    "Set `ACCESS_TOKEN`",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "access-type",
								Usage:    "Set authorization type `TYPE` ('Bearer', 'Basic', etc.)",
								Required: false,
								Value:    "Bearer",
							},
						},
					},
					{
						Name:        "list",
						Action:      Auth.List,
						Description: "List credentialed access token",
						Usage:       "List credentialed access token",
						UsageText:   "onesky auth list [options]",
					},
					{
						Name:        "revoke",
						Action:      Auth.Revoke,
						Description: "Revoke access credential",
						Usage:       "Revoke access credential",
						UsageText:   "onesky auth revoke [options]",
					},
				},

				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					_, _ = fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
			},
			// LANG
			{
				Name:      "lang",
				Usage:     "Manage languages of the app",
				UsageText: "onesky [global options] lang <command> [options]",
				Description: "List all enabled languages of the app: \n" +
					"			onesky lang list\n",
				HideHelpCommand: true,
				Subcommands: []*cli.Command{
					{
						Name:        "list",
						Aliases:     []string{"l"},
						Action:      Lang.List,
						Description: "List all enabled languages of the app",
						Usage:       "List all enabled languages of the app",
						UsageText:   "onesky lang list",
					},
				},
			},
			// FILE
			{
				Name:  "file",
				Usage: "Manage files of the app",
				//Usage:           "onesky file upload --platform-id=web --language-id=en_US --file-name=en_US.json --content='{"apple-key": "Apple"}'",
				UsageText: "onesky [global options] file <command> [options]",
				Description: "Upload content from command-line: \n" +
					"			onesky file upload --platform-id=web --language-id=en_US --file-name=en_US.json --content='{\"apple-key\": \"Apple\"}'\n\n" +
					"   Upload content from local file: \n" +
					"			onesky file upload --platform-id=web --language-id=en_US --file-name=en_US.json --path=path/to/file/with/valid.ext\n\n" +
					"	Download file by file id: \n" +
					"			onesky file download --file-id=\"09547d3f-3734-4efd-801a-0aea74fc301e\" --plugin-agent=intellij\n\n" +
					"	List all files of the app: \n" +
					"			onesky file list\n",
				HideHelpCommand: true,
				Subcommands: []*cli.Command{
					{
						Name:        "list",
						Action:      File.List,
						Aliases:     []string{"l"},
						Description: "List all files of the app",
						Usage:       "List all files of the app",
						UsageText:   "`onesky file list`",
					},
					{
						Name:        "upload",
						Action:      File.Upload,
						Aliases:     []string{"u"},
						Description: "Upload a new file to the app",
						Usage:       "Upload a new file to the app",
						UsageText:   "onesky file upload --platform-id=PLATFORM_ID --language-id=LANGUAGE_ID --file-name=FILE_NAME {--content=CONTENT_ENCODED_IN_UTF8|--path=PATH} [--plugin-agent=intellij]",
						Flags: []cli.Flag{
							File.FlagPlatformId,
							File.FlagLanguageId,
							File.FlagFileName,
							File.FlagFileContent,
							File.FlagFilePath,
							File.FlagPluginAgent,
						},
					},
					{
						Name:        "download",
						Aliases:     []string{"d"},
						Action:      File.Download,
						Description: "Download file by file id",
						Usage:       "Download file by file id",
						UsageText:   "onesky file download --file-id=FILE_ID [--plugin-agent=USER_AGENT_HEADER_COMMENT]",
						Flags: []cli.Flag{
							File.FlagFileId,
							File.FlagOutput,
							File.FlagPluginAgent,
						},
					},
				},

				OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
					_, _ = fmt.Fprintf(c.App.Writer, "for shame\n")
					return err
				},
			},

			// API
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
