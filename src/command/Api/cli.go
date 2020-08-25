package Api

import (
	"github.com/onesky/onesky-sdk-cli/src/build"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:      "api",
	Usage:     "Manage api configuration",
	UsageText: "onesky [global options] api <command> [options]",
	Description: "Show information about api configuration: \n" +
		"			onesky api info\n\n" +
		"   Set options of api configuration: \n" +
		"			onesky set --url=\"https://management-api.onesky.app/v1\" --timeout=10\n",
	HideHelp:        true,
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		SubcommandInfo,
		SubcommandSet,
	},
}

var SubcommandInfo = &cli.Command{
	Name:        "info",
	Aliases:     []string{"i"},
	Action:      List,
	Description: "Show information about api configuration",
	Usage:       "Show information about api configuration",
	UsageText:   "onesky api info",
}

var SubcommandSet = &cli.Command{
	Name:        "set",
	Aliases:     []string{"s"},
	Action:      Set,
	Description: "Set options of api configuration",
	Usage:       "Set options of api configuration",
	UsageText:   "onesky api set",
	Flags: []cli.Flag{
		SubcommandSetFlagUrl,
		SubcommandSetFlagTimeout,
	},
}

var SubcommandSetFlagUrl = &cli.StringFlag{
	Name:  "url",
	Usage: "`URL` - Base url",
	Value: build.DefaultConfig.Api.Url,
}

var SubcommandSetFlagTimeout = &cli.IntFlag{
	Name:  "timeout",
	Usage: "`TIMEOUT` - Request timeout in seconds",
	Value: build.DefaultConfig.Api.Timeout,
}
