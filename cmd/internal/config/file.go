package config

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdFile = &cli.Command{
	Name:        "file",
	Usage:       fmt.Sprintf("Return the configuration file for %s", base.CommandName),
	Description: fmt.Sprintf("Return the configuration file for %s", base.CommandName),
	Action: func(cCtx *cli.Context) error {
		fmt.Println(base.ConfigFilePath)
		return nil
	},
}
