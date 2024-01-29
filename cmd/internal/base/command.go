package base

const (
	CommandName        = "iszi"
	CommandUsage       = "ISZI: Infrastellar Systems Zeropoint Interface"
	CommandDescription = `ISZI is the primary interface to the Infrastellar ecosystem.`

	CommandConfigFileName  = "config.json"
	CommandProgramFileName = "PROGRAM"
	CommandProgramEnvVar   = "INFRASTELLAR_PROGRAM"
)

var SubCommands = []*cli.Command{
	CmdConfig,
	CmdProgram,
	// CmdEnvironment,
	// CmdMission,
	// CmdSpace,
}
