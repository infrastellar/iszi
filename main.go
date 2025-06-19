package main

import (
	"context"
	"fmt"
	"os"

	infrastellar "github.com/infrastellar/is-sdk-go-v1/cmd"
	"github.com/urfave/cli/v3"
)

func main() {
	cli.OsExiter = func(code int) {
		os.Exit(code)
	}

	app := &cli.Command{
		Name:        "iszi",
		Usage:       "XXXXX",
		Description: "XXXXX",
	}

	app.EnableShellCompletion = true
	app.UseShortOptionHandling = true
	app.Suggest = true

	app.Commands = append([]*cli.Command{}, infrastellar.Commands...)

	err := app.Run(context.TODO(), os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(app.ErrWriter, "ISZI Command error: %v\n", err)
		cli.OsExiter(1)
	}
}
