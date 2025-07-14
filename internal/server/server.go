package server

import (
	"log"
	"net/http"
	"time"

	"github.com/magabrotheeeer/sprint6/internal/handlers"
)

type Server struct {
	log    *log.Logger
	server *http.Server
}

func createRouter(log *log.Logger) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.IndexHtml)
	router.HandleFunc("upload", handlers.Upload)
	return router
}

func NewServer(logger *log.Logger) (*Server, error) {
	router := createRouter(logger)

	httpServer := &http.Server{
		Addr: 	     "localhost:8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
		IdleTimeout:  time.Second * 15,
	}

	err := http.ListenAndServe(httpServer.Addr, router)

	return &Server{
		log:    logger,
		server: httpServer,
	}, err
}
