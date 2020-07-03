// +build windows

package build

import "os"

const PS = '\\'

//const DefaultConfigPath = "%APPDATA%\\onesky.toml"
var CONFIG_PATH = os.Getenv("APPDATA") + PS + "onesky.toml"
var PRODUCT_NAME = "onesky-cli-windows"
