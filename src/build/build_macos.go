// +build darwin,!test

package build

import (
	"os"
)

//const PS = "/"

//const DefaultConfigPath = "$HOME/.config/onesky.toml"
var DefaultConfigPath = os.Getenv("HOME") + "/.config/onesky.toml"
var ProductName = "onesky-cli-macos"
