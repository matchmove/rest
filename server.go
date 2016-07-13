package rest

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

var invokeCount int

// Server represents information about a rest server.
type Server struct {
	port        string
	environment string

}

func (server *Server) Make(config Config, r Routes) {
	server.port = config.port
	server.environment = config.environment
	server.listen(NewRouter(r))
}

func (server *Server) listen(router *mux.Router) {
	http.ListenAndServe(":" + server.port, router)
}

func InvokeHandler(handler http.Handler, routePath string, w http.ResponseWriter, r *http.Request) {

	// Add a new sub-path for each invocation since
	// we cannot (easily) remove old handler
	invokeCount++
	router := mux.NewRouter()
	http.Handle(fmt.Sprintf("/%d", invokeCount), router)

	router.Path(routePath).Handler(handler)

	// Modify the request to add "/%d" to the request-URL
	r.URL.RawPath = fmt.Sprintf("/%d%s", invokeCount, r.URL.RawPath)
	router.ServeHTTP(w, r)
}