package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/adrg/xdg"
)

var (
	ConfigDir      string = fmt.Sprintf("%s/%s", UserConfigHome(), CommandName)
	ConfigFilePath string = fmt.Sprintf("%s/%s", ConfigDir, CommandConfigFileName)
	ConfigLogDir   string = fmt.Sprintf("%s/logs", ConfigDir)

	ConfigProgramFilePath string = fmt.Sprintf("%s/%s", ConfigDir, CommandProgramFileName)
)

func UserConfigHome() string {
	// Check whether the user configured a standard configuration home
	_, ok := os.LookupEnv("XDG_CONFIG_HOME")
	if ok {
		return xdg.ConfigHome
	}
	// If not, let's put our configuration in a standard place
	return filepath.Join(os.Getenv("HOME"), ".config")
}

func CommandRepostory() string {
	info, _ := debug.ReadBuildInfo()
	return info.Path
}

func CommandRepostoryURL() string {
	repo := CommandRepository()
	return fmt.Sprintf("https://%s", repo)
}

func ConfigNew() error {
	repo := CommandRepository()

	cfgdata := CommandConfig{
		LogsDir:     ConfigLogDir,
		Repository:  repo,
		RootAccount: root,
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

	fmt.Printf("%s configuration has been initialized\n", string.ToUpper(CommandName))

	return nil
}

func ConfigUpdatePrograms(program Program) error {
	raw, err := os.ReadFile(ConfigFilePath)
	if err != nil {
		return err
	}

	var cfg CommandConfig

	cfgdata, err := json.Unmarshal(raw, &cfg)

	cfgdata.Programs = append(cfgdata.Programs, program)

	json, err := json.MarshalIndent(cfgdata, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(ConfigFilePath, json, 0640)
	if err != nil {
		return err
	}

	return nil
}
