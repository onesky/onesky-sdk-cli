package Auth

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
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
		SubcommandLogin,
		SubcommandList,
		SubcommandRevoke,
	},

	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		_, _ = fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
}

var SubcommandList = &cli.Command{
	Name:        "list",
	Action:      List,
	Description: "List credentialed access token",
	Usage:       "List credentialed access token",
	UsageText:   "onesky auth list [options]",
}

var SubcommandRevoke = &cli.Command{
	Name:        "revoke",
	Action:      Revoke,
	Description: "Revoke access credential",
	Usage:       "Revoke access credential",
	UsageText:   "onesky auth revoke [options]",
}

var SubcommandLogin = &cli.Command{
	Name:        "login",
	Action:      Login,
	Description: "Authorize to access the OneSky API with access token",
	Usage:       "Authorize to access the OneSky API with access token",
	UsageText:   "onesky auth login --access-token=ACCESS_TOKEN [--access-type=TYPE]",
	Flags: []cli.Flag{
		SubcommandLoginFlagAccessToken,
		SubcommandLoginFlagAccessType,
	},
}

var SubcommandLoginFlagAccessToken = &cli.StringFlag{
	Name:     "access-token",
	Usage:    "Set `ACCESS_TOKEN`",
	Required: true,
}

var SubcommandLoginFlagAccessType = &cli.StringFlag{
	Name:     "access-type",
	Usage:    "Set authorization type `TYPE` ('Bearer', 'Basic', etc.)",
	Required: false,
	Value:    "Bearer",
}
