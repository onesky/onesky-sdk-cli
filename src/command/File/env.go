package File

import "github.com/urfave/cli/v2"

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
