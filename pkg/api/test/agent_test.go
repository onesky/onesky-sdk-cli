package test

import (
	"fmt"
	"onesky-sdk-cli/pkg/api"
	"testing"
)

var assert = func(t *testing.T, expected, got interface{}) {
	t.Error(
		fmt.Sprintf("\nExpected: '%v'\n", expected),
		fmt.Sprintf("    Got: '%v'\n", got),
	)
}

type argsAgent struct {
	name    string
	version string
	plugin  string
}
type testCaseAgent struct {
	name    string
	args    argsAgent
	isPanic bool
}

var testsAgent = []testCaseAgent{
	/* #0 */ {"OK_NoPlugin", argsAgent{"onesky", "test", ""}, false},
	/* #1 */ {"OK_Plugin", argsAgent{"onesky", "test", "test-plugin"}, false},
	/* #2 */ {"OK_PluginSpace", argsAgent{"onesky", "test", " test space "}, false},
	/* #3 */ {"OK_VersionSpace1", argsAgent{"onesky", "test ", ""}, false},
	/* #4 */ {"OK_VersionSpace2", argsAgent{"onesky", " test ", ""}, false},
	/* #5 */ {"OK_NameSpace", argsAgent{" onesky ", "test", ""}, false},
	/* #6 */ {"OK_NoName", argsAgent{"", "test", ""}, true},
	/* #7 */ {"OK_NoVersion", argsAgent{"onesky", "", ""}, true},
	/* #8 */ {"OK_NoRequirements", argsAgent{"", "", ""}, true},
	/* #9 */ {"OK_PluginSlash", argsAgent{"onesky", "test", "test/plugin"}, false},
	/* #10 */ {"OK_AllSpaces", argsAgent{" onesky ", " test", " plugin "}, false},
}

func TestRequestAgent_New(t *testing.T) {

	for _, test := range testsAgent {
		t.Run(test.name, func(t *testing.T) {

			defer func(tt testCaseAgent) {
				r := recover()
				if test.isPanic && r == nil {
					t.Error("Expected panic!")
				} else if !test.isPanic && r != nil {
					t.Error("Unexpected panic!")
				}

			}(test)

			agent := api.NewRequestAgent(test.args.name, test.args.version, test.args.plugin)
			if agent.Name() != test.args.name {
				t.Error(
					"Expected", test.args.name,
					"\ngot", agent.Name(),
				)
			}

			if agent.Version() != test.args.version {
				assert(t, test.args.version, agent.Version())
			}

			if agent.Plugin() != test.args.plugin {
				assert(t, test.args.plugin, agent.Plugin())
			}
		})
	}
}

func TestRequestAgent_SetPlugin(t *testing.T) {

	for _, test := range testsAgent {
		t.Run(test.name, func(t *testing.T) {
			if test.isPanic {
				return
			}

			agent := api.NewRequestAgent(test.args.name, test.args.version, "")
			agent.SetPlugin(test.args.plugin)

			if agent.Plugin() != test.args.plugin {
				assert(t, test.args.plugin, agent.Plugin())
			}
		})
	}
}

func TestRequestAgent_String(t *testing.T) {

	var stringTestCases = []struct {
		testCaseAgent
		name   string
		cmp    string
		result bool
	}{
		{testsAgent[0], "OK", "onesky/test", true},
		{testsAgent[1], "OK_Plugin", "onesky/test test-plugin", true},
		{testsAgent[2], "OK_Space1", "onesky/test  test space ", true},
		{testsAgent[2], "OK_Space2", "onesky/test  test space ", true},
		{testsAgent[2], "Fail_SpacePlugin1", "onesky/test  test space", false},
		{testsAgent[2], "Fail_SpacePlugin2", "onesky/test test space ", false},
		{testsAgent[3], "OK_VersionSpace1", "onesky/test ", true},
		{testsAgent[4], "OK_VersionSpace2", "onesky/ test ", true},
	}

	for _, test := range stringTestCases {
		t.Run(test.name, func(t *testing.T) {
			if test.isPanic {
				return
			}

			agent := api.NewRequestAgent(test.args.name, test.args.version, test.args.plugin)

			if result := agent.String() == test.cmp; result != test.result {
				t.Error(
					"\nExpected", !test.result,
					fmt.Sprintf("\ncompared strings: '%s' and '%s'", agent.String(), test.cmp),
				)
			}
		})
	}
}

func TestRequestAgent_NewRequestAgentFromString1(t *testing.T) {

	var localTests = []struct {
		exp     testCaseAgent
		name    string
		arg     string
		isPanic bool
	}{
		{testsAgent[0], "OK", "onesky/test", false},
		{testsAgent[9], "OK_Spaces", " onesky / test   plugin ", false},
		{testsAgent[1], "OK_Plugin1", "onesky/test plugin", false},
		{testsAgent[1], "OK_Plugin2", "onesky/test test/plugin", false},
		{testsAgent[1], "OK_Plugin3", "onesky/test test-plugin", false},
		{testsAgent[0], "Fail_Delimiter", "onesky test plugin", true},
		{testsAgent[3], "OK_SpaceVersion", "onesky/test ", false},
		{testsAgent[2], "OK_SpacePlugin", "onesky/test  test space ", false},
		{testsAgent[2], "Fail_1", "onesky /", true},
		{testsAgent[2], "Fail_2", "onesky test", true},
		{testsAgent[2], "Fail_3", "onesky ", true},
	}

	for _, test := range localTests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if test.isPanic && r == nil {
					t.Error("Expected panic!")
				} else if !test.isPanic && r != nil {
					t.Error("Unexpected panic!:", r)
				}

			}()

			agent := api.NewRequestAgentFromString(test.arg)

			if agent.String() != test.arg {
				//assert(t, test.arg, agent.String())
				t.Error(
					fmt.Sprintln(test.exp),
					fmt.Sprintf("\nExpected: '%v'\n", test.arg),
					fmt.Sprintf("    Got: '%v'\n", agent.String()),
				)
			}
		})
	}
}

func TestRequestAgent_NewRequestAgentFromString2(t *testing.T) {

	for _, test := range testsAgent {
		t.Run(test.name, func(t *testing.T) {

			if test.isPanic {
				return
			}

			agent := api.NewRequestAgent(test.args.name, test.args.version, test.args.plugin)
			agentFromString := api.NewRequestAgentFromString(agent.String())

			if agent.Name() != agentFromString.Name() {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", agentFromString.Name()),
					fmt.Sprintf("Expected: '%s'\n", agent.Name()),
				)
			}

			if agent.Version() != agentFromString.Version() {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", agentFromString.Version()),
					fmt.Sprintf("Expected: '%s'\n", agent.Version()),
				)
			}

			if agent.Plugin() != agentFromString.Plugin() {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", agentFromString.Plugin()),
					fmt.Sprintf("Expected: '%s'\n", agent.Plugin()),
				)
			}

			if agent.String() != agentFromString.String() {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", agentFromString.String()),
					fmt.Sprintf("Expected: '%s'\n", agent.String()),
				)
			}
		})
	}
}
