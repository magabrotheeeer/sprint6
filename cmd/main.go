package main

import (
	"log"
	"os"

	"github.com/magabrotheeeer/sprint6/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	_, err := server.NewServer(logger)

	if err != nil {
		logger.Fatalf("Error when starting server")
	}
	logger.Println("Starting server")
}

