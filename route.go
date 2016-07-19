package rest

// Test cases are covered in server_test.go
import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents the struct of Route
// properties:
// - Name  		string  			Route name
// - Pattern  string 				Pattern or Url Pattern
// - Resource ResourceType
type Route struct {
	Name     string
	Pattern  string
	Resource ResourceType
}

// Routes represents a array/collection of Route
type Routes []Route

// NewRoute creates a new route
func NewRoute(n string, p string, r ResourceType) Route {
	return Route{Name: n, Pattern: p, Resource: r}
}

// GetHandler is the method that handles the http.HandlerFunc
func (route Route) GetHandler(s *Server) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		l := NewLog()
		defer l.Dump()

		route.Resource.Set(mux.Vars(r), w, r, &l, s)

		route.Resource.Init()

		switch r.Method {
		case http.MethodGet:
			route.Resource.Get()
			break
		case http.MethodPost:
			route.Resource.Post()
			break
		case http.MethodPut:
			route.Resource.Put()
			break
		case http.MethodPatch:
			route.Resource.Patch()
			break
		case http.MethodDelete:
			route.Resource.Delete()
			break
		}

		route.Resource.Deinit()
	}
}

// ApplyRoutes set the Routes given the array of route
func ApplyRoutes(router *mux.Router, routes Routes, s *Server) *mux.Router {
	for _, route := range routes {
		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(http.HandlerFunc(route.GetHandler(s)))
	}

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentTypeTextPlain)
		w.WriteHeader(http.StatusNotFound)
	})

	return router
}
