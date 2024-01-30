package program

import (
	"github.com/urfave/cli/v2"
)

var SubCmdAccount = &cli.Command{
	Name:        "account",
	Usage:       "Configure the main (root) AWS account for the Infrastellar Space Program",
	Description: `Configure the main (root) AWS account for the infrastructure program.`,
	Action: func(cCtx *cli.Context) error {
		return nil
	},
}
