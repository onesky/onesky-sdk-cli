// +build !windows

package build

import "os"

const PS = "/"

//const CONFIG_PATH = "$HOME/.config/onesky.toml"
var CONFIG_PATH = os.Getenv("HOME") + "/.config/onesky.toml"
