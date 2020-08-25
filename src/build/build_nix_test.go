// +build darwin linux
// +build test

package build

import "os"

//const PS = "/"

//const DefaultConfigPath = "$HOME/.config/onesky.toml"
var DefaultConfigPath = os.Getenv("HOME") + "/.config/onesky-test.toml"
var ProductName = "onesky-cli-test"
