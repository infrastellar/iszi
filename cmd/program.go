package cmd

import (
	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/program"
)

var CmdProgram = &cli.Command{
	Name:        "program",
	Usage:       "Interface with an Infrastellar Space Program",
	Description: "Interact with an Infrastellar Space Program infrastructure setup",
	Subcommands: []*cli.Command{
		program.SubCmdActive,
		program.SubCmdNew,
		program.SubCmdSet,
		program.SubCmdUnset,
	},
}
