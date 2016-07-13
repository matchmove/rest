package rest

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Route struct {
	Name     string
	Pattern  string
	Controller Resource
}

type Routes []Route

func NewRouter(routes Routes) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
		Path(route.Pattern).
		Name(route.Name).
		Handler(http.HandlerFunc(route.Controller.MainFunc()))
	}

	return router
}
