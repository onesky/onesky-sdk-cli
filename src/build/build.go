package build

import "OneSky-cli/pkg/config"

var DefaultConfig = config.OneskyConfig{
	Title: "OneSky config",
	Api: config.Api{
		Url:     "https://management-api.onesky.app/v1",
		Timeout: 30,
	},
}
