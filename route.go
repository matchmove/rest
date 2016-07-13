package rest

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name     string
	Pattern  string
	Resource ResourceType
}

type Routes []Route

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
