// +build linux,!test

package build

import "os"

const PS = "/"

var DefaultConfigPath = os.Getenv("HOME") + PS + ".config" + PS + "onesky.toml"
var ProductName = "onesky-cli-linux"
