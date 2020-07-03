// +build linux

package build

import "os"

const PS = "/"

//const DefaultConfigPath = "$HOME/.config/onesky.toml"
var CONFIG_PATH = os.Getenv("HOME") + "/.config/onesky.toml"
var PRODUCT_NAME = "onesky-cli-linux"
