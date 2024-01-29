package config

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd/internal/base"
)

var (
	flagLogs bool

	SubCmdDir = &cli.Command{
		Name:        "dir",
		Usage:       fmt.Sprintf("Return the configuration path for %s", base.CommandName),
		Description: fmt.Sprintf("Return the configuration path for %s", base.CommandName),
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "logs",
				Usage:       fmt.Sprintf("Return the log path for %s", base.CommandName),
				Destination: &flagLogs,
			},
		},
		Action: func(cCtx *cli.Context) error {
			if flagLogs {
				fmt.Println(base.ConfigLogsDir)
				return nil
			}
			fmt.Println(base.ConfigDir)
			return nil
		},
	}
)
