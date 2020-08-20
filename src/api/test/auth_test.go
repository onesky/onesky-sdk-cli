package test

import (
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/api"
	"testing"
)

type argsAuth struct {
	authValue string
	authType  string
}

type testCaseAuth struct {
	name    string
	args    argsAuth
	isPanic bool
}

var testsAuth = []testCaseAuth{
	/* #0 */ {"OK", argsAuth{"token", "Bearer"}, false},
	/* #1 */ {"OK_NoType", argsAuth{"token", ""}, false},
	/* #1 */ {"OK_TokenSpace", argsAuth{" long token-with/spaces&dig1ts ", ""}, false},
	/* #1 */ {"OK_TokenSpaceWithType", argsAuth{" long token-with/spaces&dig1ts ", "Basic"}, false},
	/* #2 */ {"Fail_Value", argsAuth{"", "test"}, true},
}

func TestRequestAuthorization_NewRequestAuthorization(t *testing.T) {

	for _, test := range testsAuth {
		t.Run(test.name, func(t *testing.T) {

			defer func(tt testCaseAuth) {
				r := recover()
				if test.isPanic && r == nil {
					t.Error("Expected panic!")
				} else if !test.isPanic && r != nil {
					t.Error("Unexpected panic!")
				}

			}(test)

			auth := api.NewRequestAuthorization(test.args.authValue, test.args.authType)

			if auth.Value() != test.args.authValue {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", test.args.authValue),
					fmt.Sprintf("    Got: '%v'\n", auth.Value()),
				)
			}

			if auth.Type() != test.args.authType {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", test.args.authType),
					fmt.Sprintf("    Got: '%v'\n", auth.Type()),
				)
			}
		})
	}
}

func TestRequestAuthorization_NewRequestAuthorizationFromString1(t *testing.T) {

	type localTests struct {
		name string
		testCaseAuth
		authString string
		isPanic    bool
	}
	var testsAuthLocal = []localTests{
		{"OK", testsAuth[0], "Bearer token", false},
		{"OK_NoType", testsAuth[1], "token", false},
		{"OK_StrangeToken", testsAuth[2], " long token-with/spaces&dig1ts ", false},
		{"OK_StrangeToken2", testsAuth[3], "Basic  long token-with/spaces&dig1ts ", false},
		{"Fail", testsAuth[4], "", true},
	}

	for _, test := range testsAuthLocal {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				r := recover()
				if test.isPanic && r == nil {
					t.Error("Expected panic!")
				} else if !test.isPanic && r != nil {
					t.Error("Unexpected panic!")
				}

			}()

			authGot := api.NewRequestAuthorizationFromString(test.authString)
			authExp := api.NewRequestAuthorization(test.args.authValue, test.args.authType)

			if authExp.Value() != authGot.Value() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.Value()),
					fmt.Sprintf("    Got: '%v'\n", authGot.Value()),
				)
			}

			if authExp.Type() != authGot.Type() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.Type()),
					fmt.Sprintf("    Got: '%v'\n", authGot.Type()),
				)
			}

			if authExp.String() != authGot.String() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.String()),
					fmt.Sprintf("    Got: '%v'\n", authGot.String()),
				)
			}
		})
	}
}

func TestRequestAuthorization_NewRequestAuthorizationFromString2(t *testing.T) {

	for _, test := range testsAuth {
		t.Run(test.name, func(t *testing.T) {

			if test.isPanic {
				return
			}

			authExp := api.NewRequestAuthorization(test.args.authValue, test.args.authType)
			authGot := api.NewRequestAuthorizationFromString(authExp.String())

			if authExp.Value() != authGot.Value() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.Value()),
					fmt.Sprintf("    Got: '%v'\n", authGot.Value()),
				)
			}

			if authExp.Type() != authGot.Type() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.Type()),
					fmt.Sprintf("    Got: '%v'\n", authGot.Type()),
				)
			}

			if authExp.String() != authGot.String() {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", authExp.String()),
					fmt.Sprintf("    Got: '%v'\n", authGot.String()),
				)
			}
		})
	}
}

func TestRequestAuthorization_SetType(t *testing.T) {

	type localTests struct {
		name string
		testCaseAuth
		authType string
	}
	var testsAuthLocal = []localTests{
		{"OK", testsAuth[0], "Bearer"},
		{"OK_Custom", testsAuth[1], "None"},
		{"OK_TypeSpace", testsAuth[2], " "},
		{"OK_TypeReset", testsAuth[3], ""},
	}

	for _, test := range testsAuthLocal {
		t.Run(test.name, func(t *testing.T) {

			auth := api.NewRequestAuthorization(test.args.authValue, test.args.authType)
			auth.SetType(test.authType)

			if auth.Type() != test.authType {
				t.Error(
					fmt.Sprintln(test.args),
					fmt.Sprintf("\nExpected: '%v'\n", test.authType),
					fmt.Sprintf("    Got: '%v'\n", auth.Type()),
				)
			}
		})
	}
}
