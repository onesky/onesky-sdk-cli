package build

import (
	"flag"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"testing"
)

func createTestContext() *cli.Context {
	return cli.NewContext(
		cli.NewApp(),
		flag.NewFlagSet("testSet", flag.ContinueOnError),
		nil,
	)
}

func TestContext_BuildConfigDefault(t *testing.T) {

	ctx := createTestContext()
	if config, err := CreateConfig(ctx); err != nil {
		t.Error("Unexpected error:", err)
	} else {
		if err := os.Remove(config.Source()); err != nil {
			t.Error("Unexpected error:", err)
		}
	}
}

func TestContext_BuildConfigDefaultExist(t *testing.T) {

	ctx := createTestContext()
	_, _ = CreateConfig(ctx)
	if config, err := CreateConfig(ctx); err != nil {
		t.Error("Unexpected error:", err)
	} else {
		if err := os.Remove(config.Source()); err != nil {
			t.Error("Unexpected error:", err)
		}
	}
}

func TestContext_BuildConfigGiven(t *testing.T) {

	ctx := createTestContext()
	ctx.App.Flags = []cli.Flag{&cli.StringFlag{
		Name:     "config-file",
		Usage:    "Set alternative config `PATH`",
		Required: true,
	}}

	ctx.App.Action = func(context *cli.Context) error {
		config, err := CreateConfig(context)
		if err == nil {
			_ = os.Remove(config.Source())
		}
		return err
	}

	existedConfig, err := CreateConfig(ctx)
	if err != nil {
		t.Fatal("Unexpected error:", err)
	}

	err = ctx.App.RunContext(ctx.Context, []string{"test", "--config-file", existedConfig.Source()})
	if err != nil {
		t.Error("Unexpected error:", err)
	}
}

func TestContext_BuildConfigErrorPath(t *testing.T) {

	ctx := createTestContext()
	ctx.App.Flags = []cli.Flag{&cli.StringFlag{
		Name:     "config-file",
		Usage:    "Set alternative config `PATH`",
		Required: true,
	}}
	ctx.App.Action = func(context *cli.Context) error {
		_, err := CreateConfig(context)
		if err == nil {
			t.Error(
				fmt.Sprintf("\nGot: '%s'\n", err),
				fmt.Sprintf("Expected: '%s'\n", "Config not found"),
			)
		}
		return err
	}

	_ = ctx.App.RunContext(ctx.Context, []string{"test", "--config-file", "invalidPath"})
}

func TestContext_CreateAppContext(t *testing.T) {

	cCtx := cli.NewContext(
		cli.NewApp(),
		flag.NewFlagSet("testSet", flag.ContinueOnError),
		nil,
	)

	cCtx.App.Setup()

	config, err := CreateConfig(cCtx)
	if err != nil {
		t.Error("Unexpected error:", err)
	}

	if appContext, err := CreateAppContext(cCtx, &config); err != nil {
		t.Error("Unexpected error:", err)
	} else {
		if err := os.Remove(appContext.Config().Source()); err != nil {
			t.Error("Unexpected error:", err)
		}
	}
}
