package api

import (
	"fmt"
	"regexp"
)

type RequestAgent interface {
	Name() string
	Version() string
	Plugin() string
	SetPlugin(string)
	String() string
}

type requestAgent struct {
	name    string
	version string
	plugin  string
}

func NewRequestAgent(name, version, plugin string) RequestAgent {
	if name == "" {
		panic("required RequestAgent.Name")
	}

	if version == "" {
		panic("required RequestAgent.Version")
	}

	a := &requestAgent{name, version, plugin}

	return a
}

func NewRequestAgentFromString(agentString string) RequestAgent {
	re, _ := regexp.Compile(`^([\s\w-]+)/(\s*[\w-]+)(\s)?(.+)?$(?i)`)

	found := re.FindAllStringSubmatch(agentString, -1)
	if len(found) > 0 && len(found[0]) > 1 {
		parts := found[0]
		name := parts[1]
		version := parts[2]
		plugin := ""

		if parts[4] == "" {
			version += parts[3]
		} else {
			plugin = parts[4]
		}

		return NewRequestAgent(name, version, plugin)
	}

	panic("Invalid string for RequestAgent: \"" + agentString + "\"")
}

func (r *requestAgent) Name() string {
	return r.name
}

func (r *requestAgent) Version() string {
	return r.version
}

func (r *requestAgent) Plugin() string {
	return r.plugin
}

func (r *requestAgent) SetPlugin(s string) {
	r.plugin = s
}

func (r *requestAgent) String() string {
	if r.Name() != "" && r.Version() != "" {
		agentString := fmt.Sprintf("%s/%s", r.Name(), r.Version())

		if r.Plugin() != "" {
			agentString += " " + r.Plugin()
		}
		return agentString
	}
	return ""
}
