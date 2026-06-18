package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func NewServer(log *log.Logger) *Server {
	h := &handlers.Handler{
		Logger: log,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", h.GetHome)
	mux.HandleFunc("/upload", h.HandleUpload)

	var server = &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     log,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{Logger: log, Server: server}
}
