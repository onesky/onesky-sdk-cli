package api

import (
	"fmt"
	"strings"
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

	return &requestAgent{name, version, plugin}
}

func NewRequestAgentFromString(agentString string) RequestAgent {
	parts := strings.Fields(agentString)

	switch len(parts) {
	case 3:
		{ // name / version
			return &requestAgent{parts[0], parts[2], ""}
		}
	case 4:
		{ //name / version plugin
			return &requestAgent{parts[0], parts[2], parts[3]}
		}
	default:
		panic("Invalid string for RequestAgent: \"" + agentString + "\"")
	}
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
		return fmt.Sprintf("%s / %s %s​", r.Name(), r.Version(), r.Plugin())
	}
	return ""
}
