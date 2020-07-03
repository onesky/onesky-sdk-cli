package context

import "OneSky-cli/pkg/config"

type AppFlags struct {
	AuthString string // global token
	AuthType   string // global token
	Debug      bool
	ConfigPath string
}

type AppVars struct {
	//DefaultConfig		config.OneskyConfig
	//ConfigPath	string
}

type AppBuild struct {
	ConfigPath     string
	ProductName    string
	ProductVersion string
	BuildId        string
	BuildInfo      string
}

type appContext struct {
	config *config.OneskyConfig
	build  *AppBuild
	vars   *AppVars
	flags  *AppFlags
}

type AppContext interface {
	Config() *config.OneskyConfig
	Build() *AppBuild
	Vars() *AppVars
	Flags() *AppFlags
}

func New(c *config.OneskyConfig) AppContext {

	if c == nil {
		c = &config.OneskyConfig{}
	}

	return &appContext{
		c,
		&AppBuild{},
		&AppVars{},
		&AppFlags{},
	}
}

func (a *appContext) Config() *config.OneskyConfig {
	return a.config
}

func (a appContext) Build() *AppBuild {
	return a.build
}

func (a appContext) Vars() *AppVars {
	return a.vars
}

func (a appContext) Flags() *AppFlags {
	return a.flags
}
