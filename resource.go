package rest

import (
	"net/http"
	"github.com/gorilla/mux"
)


type Resource func (w http.ResponseWriter, r *http.Request, vars map[string]string)

func (resource Resource) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource(w, r, vars)
}

