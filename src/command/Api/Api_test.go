package Api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/onesky/onesky-sdk-cli/src/command/test"
	"github.com/urfave/cli/v2"
	"strconv"
	"strings"
	"testing"
)

func TestApi_List(t *testing.T) {

	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	gotBytes := test.Cap(func() {
		if err := List(cliCtx); err != nil {
			t.Error("Unexpected error:", err)
		}
	})

	got := string(gotBytes)
	var expected map[string]interface{}
	for k, v := range structs.Map(appCtx.Config().Api) {
		t1 := strings.Index(got, k)
		t2 := strings.Index(got, fmt.Sprint(v))

		if t1 < 0 || t2 < 0 {
			t.Error(
				fmt.Sprintf("\nGot: '%s'\n", got),
				fmt.Sprintf("Expected: '%s'\n", expected),
			)
		}
	}
}

func TestApi_Set(t *testing.T) {
	cliCtx := test.CreateTestCliContext()
	appCtx := test.CreateTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	cliCtx.App.Commands = []*cli.Command{
		Command,
	}

	expectedTimeout := 11
	expectedUrl := "test-url"

	err := cliCtx.App.Run([]string{"app", "api", "set", "--timeout", strconv.Itoa(expectedTimeout), "--url", expectedUrl})
	if err != nil {
		t.Fatal(err)
	}

	if expectedTimeout != appCtx.Config().Api.Timeout {
		t.Error(
			fmt.Sprintf("\nGot: '%v'\n", appCtx.Config().Api.Timeout),
			fmt.Sprintf("Expected: '%v'\n", expectedTimeout),
		)
	}

	if expectedUrl != appCtx.Config().Api.Url {
		t.Error(
			fmt.Sprintf("\nGot: '%v'\n", appCtx.Config().Api.Url),
			fmt.Sprintf("Expected: '%v'\n", expectedUrl),
		)
	}

}
