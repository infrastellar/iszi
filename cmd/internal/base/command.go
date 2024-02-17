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

	// Environment Variables
	EnvVarProgram = "INFRASTELLAR_PROGRAM"
	// Environment Variables we can use instead of command line arguments
	// The following envvar values are all relative to the INFRASTELLAR_PROGRAM
	EnvVarEnvironment = "INFRASTELLAR_ENVIRONMENT"
	EnvVarEnvRegion   = "INFRASTELLAR_ENVIRONMENT_REGION"
	EnvVarMission     = "INFRASTELLAR_MISSION"
	EnvVarStage       = "INFRASTELLAR_STAGE"
	EnvVarProcedure   = "INFRASTELLAR_PROCEDURE"
)

// GlobalFlags represents global flags for the command
var GlobalFlags = []cli.Flag{}
