package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)
	srv := server.NewServer(logger)

	if err := srv.Server.ListenAndServe(); err != nil {
		logger.Println("error when starting the service")
	}
	logger.Println("Starting the server")
}

