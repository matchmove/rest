package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// StatusMethodNotAllowed defines the HTTP message when 405 is encountered
	StatusMethodNotAllowed = "Method Not Allowed"
	// ContentTypeJSON defines Content-Type for application/json
	ContentTypeJSON = "application/json"
)

// ResourceType represents an interface information about a rest resource.
// need to add 3 properties:
// - Vars  map[string]string
// - Write http.ResponseWriter
// - Read  *http.Request
type ResourceType interface {
	Set(map[string]string, http.ResponseWriter, *http.Request)

	Init()

	Get()

	Put()

	Post()

	Patch()

	Delete()

	Deinit()
}

type Resource struct {
	Vars  map[string]string
	Write http.ResponseWriter
	Read  *http.Request
}

func (c *Resource) Set(vars map[string]string, w http.ResponseWriter, r *http.Request) {
	c.Vars = mux.Vars(r)
	c.Write = w
	c.Read = r
}

func (c *Resource) SetContentType(ctype string) {
	c.Write.Header().Set("Content-Type", ctype)
}

func (c *Resource) Init() {
	c.SetContentType(ContentTypeJSON)
}

func (c *Resource) Get() {

}

func (c *Resource) Put() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

func (c *Resource) Post() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

func (c *Resource) Patch() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

func (c *Resource) Delete() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

func (c *Resource) Deinit() {

}
