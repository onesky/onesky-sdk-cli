package File

import "github.com/urfave/cli"

var FLAG_Output = &cli.StringFlag{
	Name:     "output",
	Aliases:  []string{"o"},
	Usage:    "​Save output to given `FILE_NAME`",
	Required: true,
}

var FLAG_FileId = &cli.StringFlag{
	Name:     "file-id",
	Aliases:  []string{"i"},
	Usage:    "​`FILE_ID`",
	Required: true,
}

var FLAG_PluginAgent = &cli.StringFlag{
	Name: "plugin-agent",
	//Aliases: []string{"a"},
	Usage: "​​`USER_AGENT_HEADER_COMMENT` (Ex.: --plugin-agent=intellij)",
}
