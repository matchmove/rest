package rest

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

// GetHandler is the method that handles the http.HandlerFunc
func (route *Route) GetHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		route.Resource.Set(mux.Vars(r), w, r)

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

// NewRouter set the Routes given the array of route
func NewRouter(routes Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Path(route.Pattern).
			Name(route.Name).
			Handler(http.HandlerFunc(route.GetHandler()))
	}

	return router
}
