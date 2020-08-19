package build

import (
	"onesky-sdk-cli/pkg/app"
)

var DefaultConfig = app.Config{
	Title: "OneSky config",
	Api: app.Api{
		Url:     "https://management-api.onesky.app/v1",
		Timeout: 30,
	},
}
