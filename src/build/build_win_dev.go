// +build windows,test

package build

import "os"

const PS = "\\"

var DefaultConfigPath = os.Getenv("APPDATA") + PS + "onesky-test.toml"
var ProductName = "onesky-cli-test"
