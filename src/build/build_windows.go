// +build windows

package build

import "os"

const PS = '\\'

//const CONFIG_PATH = "%APPDATA%\\onesky.toml"
var CONFIG_PATH = os.Getenv("APPDATA") + "/onesky.toml"
