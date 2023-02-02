package main

import (
	"log"
	"os"

	"github.com/haydenccarroll/definition-api/cmd/definition-api/server"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.App{
		Name: "definition-api",
		Commands: []*cli.Command{
			server.Command(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Println("Exited with error: ", err.Error())
		os.Exit(1)
	}
}
