package base

import (
	"github.com/urfave/cli/v2"
)

const (
	CommandName        = "iszi"
	CommandUsage       = "ISZI: Infrastellar Systems Zeropoint Interface"
	CommandDescription = `ISZI is the primary interface to the Infrastellar ecosystem.`

	CommandConfigFileName  = "config.json"
	CommandProgramFileName = "PROGRAM"
	CommandProgramEnvVar   = "INFRASTELLAR_PROGRAM"
)

var GlobalFlags = []cli.Flag{}
