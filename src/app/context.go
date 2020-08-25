package app

type Flags struct {
	AuthString string // global token
	AuthType   string // global token
	Debug      bool
	ConfigPath string
}

type Vars struct {
	//DefaultConfig		config.Config
	//ConfigPath	string
}
type Auth Credentials

type Build struct {
	ConfigPath     string
	ProductName    string //required
	ProductVersion string //required
	BuildId        string //required
	BuildInfo      string
}

type appContext struct {
	config *Config
	build  *Build
	vars   *Vars
	flags  *Flags
	auth   *Auth
}

type Context interface {
	Config() *Config
	Build() *Build
	Vars() *Vars
	Flags() *Flags
	Auth() *Auth
}

func NewContext(c *Config) Context {

	if c == nil {
		c = &Config{}
	}

	ctx := &appContext{
		c,
		&Build{},
		&Vars{},
		&Flags{},
		&Auth{},
	}

	if c.Credentials.Token != "" {
		*ctx.auth = Auth(c.Credentials)
	}

	return ctx
}

func (a *appContext) Config() *Config {
	return a.config
}

func (a *appContext) Build() *Build {
	return a.build
}

func (a *appContext) Vars() *Vars {
	return a.vars
}

func (a *appContext) Flags() *Flags {
	return a.flags
}

func (a *appContext) Auth() *Auth {
	return a.auth
}
