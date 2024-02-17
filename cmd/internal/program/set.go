package program

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdSet = &cli.Command{
	Name:  "set",
	Usage: "Set the current working directory to the active program",
	Description: fmt.Sprintf(
		"Set the current working directory to the active program and stash it in %s.\n\nThe environment variable %s is prioritized over this command and the resulting file.",
		base.ConfigProgramFilePath,
		base.EnvVarProgram,
	),
	Action: func(cCtx *cli.Context) error {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		if initialized := base.ConfigInitialized(); !initialized {
			return fmt.Errorf("%s has not been initialized, run '%s config init'", base.CommandName, base.CommandName)
		}

		err = os.WriteFile(base.ConfigProgramFilePath, []byte(cwd), 0640)
		if err != nil {
			return err
		}
		fmt.Printf(">>> The active program has been set to: %s\n", cwd)
		return nil
	},
}
