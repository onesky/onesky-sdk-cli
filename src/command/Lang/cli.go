package Lang

import "github.com/urfave/cli/v2"

var Command = &cli.Command{
	Name:      "lang",
	Usage:     "Manage languages of the app",
	UsageText: "onesky [global options] lang <command> [options]",
	Description: "List all enabled languages of the app: \n" +
		"			onesky lang list\n",
	HideHelpCommand: true,
	Subcommands: []*cli.Command{
		SubcommandList,
	},
}

var SubcommandList = &cli.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Action:      List,
	Description: "List all enabled languages of the app",
	Usage:       "List all enabled languages of the app",
	UsageText:   "onesky lang list",
}
