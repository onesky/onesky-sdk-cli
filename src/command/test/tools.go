package test

import (
	"flag"
	"fmt"
	"github.com/onesky/onesky-sdk-cli/src/app"
	"github.com/onesky/onesky-sdk-cli/src/build"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
)

func CleanupTestConfig() {
	fmt.Print("Cleanup...")

	_, err := os.Stat(build.DefaultConfigPath)
	if os.IsNotExist(err) {
		log.Println(err)
	} else {
		_ = os.Remove(build.DefaultConfigPath)
		fmt.Println("done")
	}
}

func CreateTestConfig(ctx *cli.Context) (app.Config, error) {
	CleanupTestConfig()
	return build.CreateConfig(ctx)
}

func CreateTestAppContext(ctx *cli.Context) app.Context {

	config, err := CreateTestConfig(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	appCtx, err := build.CreateAppContext(ctx, &config)
	if err != nil {
		log.Fatalln(err)
	}

	return appCtx
}

func CreateTestCliContext() *cli.Context {
	ctx := cli.NewContext(
		cli.NewApp(),
		flag.NewFlagSet("testSet", flag.ContinueOnError),
		nil,
	)
	ctx.App.Setup()
	return ctx
}

func Cap(f func()) []byte {
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
