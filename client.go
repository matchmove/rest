package rest

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

const (
	StatusMethodNotAllowed = "Method Not Allowed"
)

type ClientResource struct {

}

func (c *ClientResource) Init(){

}

func (c *ClientResource) MainFunc() func(http.ResponseWriter, *http.Request) {
	return func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := mux.Vars(r)
		switch r.Method {
			case http.MethodGet:
				c.Get(w, r, data)
				break
			case http.MethodPost:
				c.Post(w, r, data)
				break
			case http.MethodPut:
				c.Put(w, r, data)
				break
			case http.MethodPatch:
				c.Patch(w, r, data)
				break
			case http.MethodDelete:
				c.Delete(w, r, data)
				break
		}

	}
}


func (c *ClientResource) Get(w http.ResponseWriter, r *http.Request, data map[string]string) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "welcome")
}

func (c *ClientResource) Put(w http.ResponseWriter,r *http.Request, data map[string]string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, StatusMethodNotAllowed)
}

func (c *ClientResource) Post(w http.ResponseWriter,r *http.Request, data map[string]string) {
	fmt.Fprint(w, "Post Test Server!\n")
}

func (c *ClientResource) Patch(w http.ResponseWriter,r *http.Request, data map[string]string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, StatusMethodNotAllowed)
}

func (c *ClientResource) Delete(w http.ResponseWriter,r *http.Request, data map[string]string) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, StatusMethodNotAllowed)
}

func (c *ClientResource) Deinit(){

}
