package main

import (
	"os"
	"time"

	"github.com/urfave/cli/v2"

	"github.com/infrastellar/iszi/cmd"
)

func main() {
	// We pass this as metadata to our app command for use later
	cfg := map[string]interface{}{
		"start": time.Now().UTC(),
	}

	cli.OsExiter = func(code int) {
		os.Exit(code)
	}

	app := cmd.NewApp(cfg)

	_ = cmd.RunApp(app, os.Args...)
}
