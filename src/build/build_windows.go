// +build windows,!test

package build

import "os"

const PS = "\\"

var DefaultConfigPath = os.Getenv("APPDATA") + PS + "onesky.toml"
var ProductName = "onesky-cli-windows"
