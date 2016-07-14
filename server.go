package rest

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Server represents information about a rest server.
type Server struct {
	Port        string
	Environment string
	AccessLog   string

	AccessLogFile *os.File
	Router        *mux.Router
}

var (
	// EmptyHandler creates an empty pass through handler
	EmptyHandler = func(r *mux.Router) http.Handler { return r }
)

// NewServer sets up the configuration of the server and creates an instance
func NewServer(path string, r Routes) Server {
	var server Server

	if err := NewConfig(path, &server); err != nil {
		log.Fatalf("Server configuration cannot be loaded with error `%v`", err)
	}

	server.Router = NewRouter(r)

	accesLog, err := os.Create(server.AccessLog)

	if err != nil {
		log.Fatalf("Failed to create accesslog file with error `%v`", err)
	}

	server.AccessLogFile = accesLog

	return server
}

// Listen initiates the handlers
func (server *Server) Listen(h func(*mux.Router) http.Handler) {
	handler := handlers.LoggingHandler(server.AccessLogFile, h(server.Router))

	if err := http.ListenAndServe(":"+server.Port, handler); err != nil {
		log.Fatalf("Failed to start server with error `%v`", err)
	}
}
