package episode

type Producer interface {
	EpisodeSetup() error
	EpisodeRun(proceed bool) error
	EpisodeCleanup() error
}

func EpisodeDetect(procedure, path string) (*Producer, error) {
	return &Producer{}, nil
}

// Procfile is something we look for, we run a series of commands per Procfile
//
// init: <command>
// validate: <command>
// plan: <command>
// apply: <command>
// output: <command>
