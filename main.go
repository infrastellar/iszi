package main

import (
	"fmt"
	"os"

	"github.com/infrastellar/is-sdk-go-v1/cmd"

	"github.com/urfave/cli/v2"
)

func main() {
	cli.OsExiter = func(code int) {
		os.Exit(code)
	}

	app := &cli.App{
		Name:        "iszi",
		Usage:       "XXXXX",
		Description: "XXXXX",
	}
	app.EnableBashCompletion = true
	app.UseShortOptionHandling = true
	app.Suggest = true
	app.Commands = []*cli.Command{
		iscmd.Config,
		iscmd.Program,
		iscmd.Space,
	}

	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(app.ErrWriter, "ISZI Command error: %v\n", err)
		cli.OsExiter(1)
	}
}
