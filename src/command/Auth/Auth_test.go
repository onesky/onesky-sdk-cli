package Auth

import (
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/onesky/onesky-sdk-cli/src/command/test"
	"github.com/urfave/cli/v2"
	"strings"
	"testing"
)

func TestAuth_List(t *testing.T) {

	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	tests := []struct {
		name     string
		Type     string
		Token    string
		Expected string
	}{
		{"OK", "test-type", "test-token", fmt.Sprintf(MsgGetToken, "test-type", "test-token")},
		{"Empty", "", "", ErrNoToken},
		{"NoToken", "test-type", "", ErrNoToken},
		{"NoType", "", "test-token", fmt.Sprintf(MsgGetToken, "", "test-token")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			appCtx.Config().Credentials.Token = tt.Token
			appCtx.Config().Credentials.Type = tt.Type

			gotBytes := test.Cap(func() {
				if err := List(cliCtx); err != nil {
					t.Error("Unexpected error:", err)
				}
			})

			got := strings.Trim(string(gotBytes), "\n")

			if tt.Expected != got {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", got),
					fmt.Sprintf("Expected: '%s'\n", tt.Expected),
				)
			}
		})
	}
}

func TestAuth_Login(t *testing.T) {
	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	cliCtx.App.Commands = []*cli.Command{
		Command,
	}

	tests := []struct {
		name          string
		expectedType  string
		expectedToken string
	}{
		{"OK", "test-type", "test-token"},
		{"Empty", "", ""},
		{"NoToken", "test-type", ""},
		{"NoType", "", "test-token"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := cliCtx.App.Run([]string{
				cliCtx.App.Name,
				Command.Name,
				SubcommandLogin.Name,
				"--" + SubcommandLoginFlagAccessToken.Name, tt.expectedToken,
				"--" + SubcommandLoginFlagAccessType.Name, tt.expectedType,
			})
			if err != nil {
				t.Fatal(err)
			}

			if tt.expectedType != appCtx.Config().Credentials.Type {
				t.Error(
					fmt.Sprintf("\nGot: '%v'\n", appCtx.Config().Credentials.Type),
					fmt.Sprintf("Expected: '%v'\n", tt.expectedType),
				)
			}

			if tt.expectedToken != appCtx.Config().Credentials.Token {
				t.Error(
					fmt.Sprintf("\nGot: '%v'\n", appCtx.Config().Credentials.Token),
					fmt.Sprintf("Expected: '%v'\n", tt.expectedToken),
				)
			}
		})
	}

}

func TestAuth_LoginNegative(t *testing.T) {
	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	cliCtx.App.Commands = []*cli.Command{
		Command,
	}

	tests := []struct {
		name          string
		expectedType  string
		expectedToken string
	}{
		{"OK", "test-type", "test-token"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*appCtx.Config() = app.Config{}

			err := cliCtx.App.Run([]string{
				cliCtx.App.Name,
				Command.Name,
				SubcommandLogin.Name,
				"--" + SubcommandLoginFlagAccessToken.Name, tt.expectedToken,
				"--" + SubcommandLoginFlagAccessType.Name, tt.expectedType,
			})
			if err == nil {
				t.Fatal("Expected error: ", app.ErrConfigNoSource)
			}
			if err.Error() != app.ErrConfigNoSource {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", err.Error()),
					fmt.Sprintf("Expected: '%s'\n", app.ErrConfigNoSource),
				)
			}
		})
	}

}

func TestAuth_Revoke(t *testing.T) {

	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	tests := []struct {
		name     string
		Type     string
		Token    string
		Expected string
	}{
		{"OK", "test-type", "test-token", MsgRevokeToken},
		{"Empty", "", "", MsgRevokeToken},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotBytes := test.Cap(func() {
				if err := Revoke(cliCtx); err != nil {
					t.Error("Unexpected error:", err)
				}
			})

			got := strings.Trim(string(gotBytes), "\n")

			if tt.Expected != got {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", got),
					fmt.Sprintf("Expected: '%s'\n", tt.Expected),
				)
			}
		})
	}
}

func TestAuth_RevokeNegative(t *testing.T) {

	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	tests := []struct {
		name     string
		Type     string
		Token    string
		Expected string
	}{
		{"OK", "test-type", "test-token", app.ErrConfigNoSource},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			*appCtx.Config() = app.Config{}

			err := Revoke(cliCtx)
			if err == nil {
				t.Fatal("Expected error: ", app.ErrConfigNoSource)
			}
			if err.Error() != app.ErrConfigNoSource {
				t.Error(
					fmt.Sprintf("\nGot: '%s'\n", err.Error()),
					fmt.Sprintf("Expected: '%s'\n", app.ErrConfigNoSource),
				)
			}
		})
	}
}
