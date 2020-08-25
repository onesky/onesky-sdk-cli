// +build darwin linux
// +build test

package build

import (
	"os"
)

const PS = "/"

var DefaultConfigPath = os.Getenv("HOME") + PS + ".config" + PS + "onesky-test.toml"
var ProductName = "onesky-cli-test"
