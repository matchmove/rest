package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Server represents information about a rest server.
type Server struct {
	Port        string
	Environment string
}

func NewServer(path string, r Routes) {
	var server Server

	if err := NewConfig(path, &server); err != nil {
		panic(err)
	}

	server.listen(NewRouter(r))
}

func (server *Server) listen(router *mux.Router) {
	http.ListenAndServe(":"+server.Port, router)
}
