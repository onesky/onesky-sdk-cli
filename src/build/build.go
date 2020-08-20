package build

import (
	"github.com/onesky/onesky-sdk-cli/src/app"
)

var DefaultConfig = app.Config{
	Title: "OneSky config",
	Api: app.Api{
		Url:     "https://management-api.onesky.app/v1",
		Timeout: 30,
	},
}
