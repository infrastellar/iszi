package space

import (
	"github.com/urfave/cli/v2"
)

var (
	flagSpace, flagRegion, flagAccountId, flagOperator string

	flagOutputJson bool
)

func BaseFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "space",
			Aliases:     []string{"s"},
			Usage:       "The space to act upon",
			Destination: &flagSpace,
		},
	}
}

func BaseOutputFlags() []cli.Flag {
	baseflags := BaseFlags()
	outflags := []cli.Flag{
		&cli.BoolFlag{
			Name:        "json",
			Aliases:     []string{"j"},
			Usage:       "Whether to provide output formatted as json",
			Destination: &flagOutputJson,
		},
	}
	baseflags = append(baseflags, outflags...)
	return baseflags
}

func BaseNewFlags() []cli.Flag {
	baseflags := BaseFlags()
	newflags := []cli.Flag{
		&cli.StringFlag{
			Name:        "region",
			Aliases:     []string{"r"},
			Usage:       "Use the provided region",
			Destination: &flagRegion,
		},
		&cli.StringFlag{
			Name:        "account-id",
			Aliases:     []string{"a"},
			Usage:       "Use the provided AWS account id",
			Destination: &flagAccountId,
		},
		&cli.StringFlag{
			Name:        "operator",
			Aliases:     []string{"o"},
			Usage:       "Use the provided operator role arn",
			Destination: &flagOperator,
		},
	}
	baseflags = append(baseflags, newflags...)
	return baseflags
}
