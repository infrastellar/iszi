package program

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdActive = &cli.Command{
	Name:        "active",
	Usage:       "Retrieve the path to the active program",
	Description: "Retrieve the path to the active program",
	Action: func(cCtx *cli.Context) error {
		file, err := os.ReadFile(base.CommandProgramFilePath)
		if err != nil {
			return err
		}
		fmt.Println(string(file[:]))
		return nil
	},
}
