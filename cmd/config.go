package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
	"github.com/infrastellar/iszi/cmd/internal/config"
)

var CmdConfig = &cli.Command{
	Name:        "config",
	Usage:       fmt.Sprintf("Manage the configuration for %s", base.CommandName),
	Description: fmt.Sprintf("Manage the configuration for %s", base.CommandName),
	Subcommands: []*cli.Command{
		config.SubCmdInit,
		config.SubCmdDir,
		config.SubCmdFile,
	},
}
