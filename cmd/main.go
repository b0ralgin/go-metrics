package main

import (
	"github.com/urfave/cli"
	"metrics/cmd/server"
	"os"
)

var (
	name    string //nolint:gochecknoglobals
	version string //nolint:gochecknoglobals
)

func main() {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Commands = []cli.Command{
		server.StartServerCommand(),
	}

	if runErr := app.Run(os.Args); runErr != nil {
		panic(runErr)
	}
}
