package config

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var SubCmdInit = &cli.Command{
	Name:        "init",
	Usage:       fmt.Sprintf("Initialize the configuration for %s", base.CommandName),
	Description: fmt.Sprintf("Initialize the configuration for %s", base.CommandName),
	Action: func(cCtx *cli.Context) error {
		err := base.ConfigNew()
		if err != nil {
			return err
		}
		return nil
	},
}
