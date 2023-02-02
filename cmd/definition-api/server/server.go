package server

import (
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "serve",
		Usage: "starts a server for definition-api",
		Action: func(ctx *cli.Context) error {
			startServer()
			return nil
		},
	}
}

func startServer() {
	s := &http.Server{
		Addr:    ":80",
		Handler: getRouter(),
	}

	log.Println("starting server")
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Println("ERROR: ", err)
	}
	log.Println("exiting server")
}
