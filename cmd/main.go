package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

func NewApp(cfg map[string]any) *cli.App {
	app := &cli.App{
		Name:        base.CommandName,
		Usage:       base.CommandUsage,
		Description: base.CommandDescription,
	}
	// Enable bash completion
	app.EnableBashCompletion = true
	// Use sparingly...
	app.Metadata = cfg
	// Support handling short flags properly
	app.UseShortOptionHandling = true

	// Global flags
	app.Flags = append(app.Flags, base.GlobalFlags...)
	app.Commands = base.SubCommands

	return app
}

func RunApp(app *cli.App, args ...string) error {
	err := app.Run(args)
	if err == nil {
		return nil
	}

	_, _ = fmt.Fprintf(app.ErrWriter, "!!! ERROR: %v\n", err)
	cli.OsExiter(1)
	return err
}
