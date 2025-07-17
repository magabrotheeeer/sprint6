package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Log    *log.Logger
	Server *http.Server
}

func createRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.IndexHtml)
	router.HandleFunc("/upload", handlers.Upload)
	return router
}

func NewServer(logger *log.Logger) *Server {
	router := createRouter()

	httpServer := &http.Server{
		Addr: 	     "localhost:8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}


	return &Server{
		Log:    logger,
		Server: httpServer,
	}
}
