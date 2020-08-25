package File

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
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
		SubcommandList,
		SubcommandUpload,
		SubcomandDownload,
	},

	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		_, _ = fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
}

var SubcommandList = &cli.Command{
	Name:        "list",
	Action:      List,
	Aliases:     []string{"l"},
	Description: "List all files of the app",
	Usage:       "List all files of the app",
	UsageText:   "`onesky file list`",
}

var SubcommandUpload = &cli.Command{
	Name:        "upload",
	Action:      Upload,
	Aliases:     []string{"u"},
	Description: "Upload a new file to the app",
	Usage:       "Upload a new file to the app",
	UsageText:   "onesky file upload --platform-id=PLATFORM_ID --language-id=LANGUAGE_ID --file-name=FILE_NAME {--content=CONTENT_ENCODED_IN_UTF8|--path=PATH} [--plugin-agent=intellij]",
	Flags: []cli.Flag{
		FlagPlatformId,
		FlagLanguageId,
		FlagFileName,
		FlagFileContent,
		FlagFilePath,
		FlagPluginAgent,
	},
}

var SubcomandDownload = &cli.Command{
	Name:        "download",
	Aliases:     []string{"d"},
	Action:      Download,
	Description: "Download file by file id",
	Usage:       "Download file by file id",
	UsageText:   "onesky file download --file-id=FILE_ID [--plugin-agent=USER_AGENT_HEADER_COMMENT]",
	Flags: []cli.Flag{
		FlagFileId,
		FlagOutput,
		FlagPluginAgent,
	},
}

var FlagOutput = &cli.StringFlag{
	Name:     "output",
	Aliases:  []string{"o"},
	Usage:    "​Save output to given `FILE_NAME`",
	Required: false,
}

var FlagFileId = &cli.StringFlag{
	Name:     "file-id",
	Aliases:  []string{"i"},
	Usage:    "​`FILE_ID`",
	Required: true,
}

var FlagPluginAgent = &cli.StringFlag{
	Name: "plugin-agent",
	//Aliases: []string{"a"},
	Usage: "​​`USER_AGENT_HEADER_COMMENT` (Ex.: --plugin-agent=intellij)",
}

var FlagPlatformId = &cli.StringFlag{
	Name:    "platform-id",
	Aliases: []string{"p"},
	Usage:   "`PLATFORM_ID` - one of 'web', 'ios' or 'android' (Ex.: --platform-id=web)",
}

var FlagLanguageId = &cli.StringFlag{
	Name:    "language-id",
	Aliases: []string{"l"},
	Usage:   "`LANGUAGE_ID` (Ex.: --language-id=en_US)",
}

var FlagFileName = &cli.StringFlag{
	Name:     "file-name",
	Aliases:  []string{"f"},
	Usage:    "`FILE_NAME` when file was uploaded",
	Required: true,
}
var FlagFileContent = &cli.StringFlag{
	Name:  "content",
	Usage: "​`CONTENT_ENCODED_IN_UTF8` (can't be used with --path option)",
}

var FlagFilePath = &cli.StringFlag{
	Name:  "path",
	Usage: "​`PATH` to text file encoded in UTF8 (can't be used with --content option)",
}
