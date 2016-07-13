package rest

import (
	"net/http"
	"github.com/gorilla/mux"

)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	MethodFunc  func(http.ResponseWriter, *http.Request, map[string]string)
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(http.HandlerFunc(route.ResourceMiddleware()))
	}

	return router
}


func (route Route) ResourceMiddleware() func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mux.Vars(r)
		route.MethodFunc(w, r, data)
	}
}