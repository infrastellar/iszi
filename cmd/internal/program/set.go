package program

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdSet = &cli.Command{
	Name:        "set",
	Usage:       "Set the current working directory to the active program",
	Description: "Set the current working directory to the active program",
	Action: func(cCtx *cli.Context) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		err := os.WriteFile(base.CommandProgramFilePath, []byte(cwd), 0640)
		if err != nil {
			return err
		}
		fmt.Printf(">>> The active program has been set to: %s\n", cwd)
		return nil
	},
}
