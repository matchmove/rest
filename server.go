package rest

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"net"

	"github.com/gorilla/mux"
)

const (
	// ServerEnvDev defines the DEVELOPMENT environment
	ServerEnvDev = "DEVELOPMENT"
	// ServerEnvTesting defines the TESTING environment
	ServerEnvTesting = "TESTING"
	// serveConnectionType defines the connection type for the server to listen
	serveConnectionType = "tcp4"
)

// Server represents information about a rest server.
type Server struct {
	URL *url.URL
	Env string

	// Handler are used in setting http activities like activity logs.
	// see: github.com/gorilla/handlers
	Handler http.Handler
	Router  *mux.Router

	// listener sets the value of the server.Listen when invoked.
	// Can be used to closed like, net.Listener
	listener io.Closer
}

// EmptyHandler creates an empty pass through handler
func EmptyHandler(r *mux.Router) http.Handler {
	return r
}

// NewServer initializes an new ReST server
func NewServer(urlRaw string) (*Server, error) {
	u, err := url.ParseRequestURI(urlRaw)

	if err != nil {
		return nil, err
	}

	return &Server{URL: u}, nil
}

// SetRoutes set the Routes given the array of route
func (s *Server) SetRoutes(router *mux.Router, routes Routes) *mux.Router {
	for _, route := range routes.stack {
		route.Server = s

		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(http.HandlerFunc(route.GetHandler(s)))
	}

	if routes.root != nil {
		router.HandleFunc("/", routes.root)
	}

	router.NotFoundHandler = http.HandlerFunc(routes.notFound)

	s.Router = router
	return router
}

// Listener returns the value of the server.Listen when invoked.
func (s *Server) Listener() io.Closer {
	return s.listener
}

// Listen initiates the handlers
func (s *Server) Listen() error {
	var (
		ln  net.Listener
		err error
	)

	if ln, err = net.Listen(serveConnectionType, ":"+s.URL.Port()); err != nil {
		return fmt.Errorf("Failed to start listener with error `%v`", err)
	}

	s.listener = ln
	return http.Serve(ln, s.Handler)
}
