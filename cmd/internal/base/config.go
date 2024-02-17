package base

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/adrg/xdg"
)

var (
	ConfigDir      string = fmt.Sprintf("%s/%s", UserConfigHome(), CommandName)
	ConfigFilePath string = fmt.Sprintf("%s/%s", ConfigDir, CommandConfigFileName)
	ConfigLogDir   string = fmt.Sprintf("%s/logs", ConfigDir)

	ConfigProgramFilePath string = fmt.Sprintf("%s/%s", ConfigDir, CommandProgramFileName)
)

// UserConfigHome manages the user preference for where we store the
// configuration for our command. Use XDG_CONFIG_HOME if it is set, otherwise
// derive a similar path ourselves.
func UserConfigHome() string {
	// Check whether the user configured a standard configuration home
	_, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		return xdg.ConfigHome
	}
	// If not, let's put our configuration in a standard place
	return filepath.Join(os.Getenv("HOME"), ".config")
}

// CommandRepository is a helper command to point to the repository
func CommandRepository() string {
	info, _ := debug.ReadBuildInfo()
	return info.Path
}

// CommandRepositoryURL is a helper command to point to the repository remote
func CommandRepostoryURL() string {
	repo := CommandRepository()
	return fmt.Sprintf("https://%s", repo)
}

// ConfigNew creates a new command configuration setup
func ConfigNew() error {
	repo := CommandRepository()

	cfgdata := CommandConfig{
		LogsDir:    ConfigLogDir,
		Repository: repo,
	}

	for _, dir := range []string{ConfigLogDir} {
		err := os.MkdirAll(dir, 0750)
		if err != nil {
			return err
		}
	}

	json, err := json.MarshalIndent(cfgdata, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFilePath, json, 0640)
	if err != nil {
		return err
	}

	fmt.Printf("%s configuration has been initialized\n", strings.ToUpper(CommandName))

	return nil
}

func ConfigInitialized() bool {
	if _, err := os.Stat(ConfigDir); !os.IsNotExist(err) {
		return true
	}
	return false
}

// ConfigUpdatePrograms manages the file where we stash our current program path
func ConfigUpdatePrograms(program *Program) error {
	raw, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return err
	}

	var cfg CommandConfig

	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return err
	}

	cfg.Programs = append(cfg.Programs, program)

	json, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFilePath, json, 0640)
	if err != nil {
		return err
	}

	return nil
}
