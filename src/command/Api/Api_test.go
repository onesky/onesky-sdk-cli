package Api

import (
	"flag"
	"fmt"
	"github.com/fatih/structs"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/onesky/onesky-sdk-cli/src/build"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"
)

func createTestAppContext(ctx *cli.Context) app.Context {

	config, err := build.CreateConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	appCtx, err := build.CreateAppContext(ctx, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return appCtx
}

func createTestCliContext() *cli.Context {
	ctx := cli.NewContext(
		cli.NewApp(),
		flag.NewFlagSet("testSet", flag.ContinueOnError),
		nil,
	)
	ctx.App.Setup()
	return ctx
}

func cap(f func()) []byte {
	backup := os.Stdout
	r, w, _ := os.Pipe()
	defer func() {
		_ = r.Close()
		_ = w.Close()
	}()
	os.Stdout = w

	f()

	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = backup

	return out
}

func TestApi_List(t *testing.T) {

	cliCtx := createTestCliContext()
	appCtx := createTestAppContext(cliCtx)

	cliCtx.App.Metadata["context"] = appCtx

	gotBytes := cap(func() {
		if err := List(cliCtx); err != nil {
			t.Error("Unexpected error:", err)
		}
	})

	got := string(gotBytes)
	var expected = ""
	for k, v := range structs.Map(appCtx.Config().Api) {
		expected += fmt.Sprintf("%s: %v\n", k, v)
	}

	if expected != got {
		t.Error(
			fmt.Sprintf("\nGot: '%s'\n", got),
			fmt.Sprintf("Expected: '%s'\n", expected),
		)
	}
}

func TestApi_Set(t *testing.T) {
	cliCtx := createTestCliContext()
	appCtx := createTestAppContext(cliCtx)

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
