package rest

import (
	"log"
	"net/http"
	"os"

	"net"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	// ServerEnvDev defines the DEVELOPMENT environment
	ServerEnvDev = "DEVELOPMENT"
	// ServerEnvTesting defines the TESTING environment
	ServerEnvTesting = "TESTING"
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

// Routes sets up the configuration of the server and creates an instance
func (server *Server) Routes(r Routes, def func(http.ResponseWriter, *http.Request), router *mux.Router) {

	if router == nil {
		router = mux.NewRouter().StrictSlash(true)
	}

	if def != nil {
		router.HandleFunc("/", def)
	}

	server.Router = ApplyRoutes(router, r, server)

	accessLog, err := os.Create(server.AccessLog)

	if err != nil {
		log.Fatalf("Failed to create accesslog file with error `%v`", err)
	}

	server.AccessLogFile = accessLog
}

// Listen initiates the handlers
func (server *Server) Listen(h func(*mux.Router) http.Handler) {
	handler := handlers.LoggingHandler(server.AccessLogFile, h(server.Router))
	defer server.AccessLogFile.Close()

	ln, err := net.Listen("tcp4", ":"+server.Port)
	if err != nil {
		log.Fatalf("Failed to start server with error `%v`", err)
	}
	log.Fatal(http.Serve(ln, handler))
}
