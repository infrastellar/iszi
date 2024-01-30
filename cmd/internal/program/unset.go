package program

import (
	"os"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdUnset = &cli.Command{
	Name:        "unset",
	Usage:       "Unset the active program",
	Description: "Unset the active program",
	Action: func(cCtx *cli.Context) error {
		err := os.RemoveAll(base.ConfigProgramFilePath)
		if err != nil {
			return err
		}
		return nil
	},
}
