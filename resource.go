package rest

// Test cases are covered in server_test.go
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
type ResourceType interface {
	Set(map[string]string, http.ResponseWriter, *http.Request, *Log, *Server)

	Init()

	Get()

	Put()

	Post()

	Patch()

	Delete()

	Deinit()
}

// Resource represents the information about the Resource.
type Resource struct {
	Vars     map[string]string
	Response http.ResponseWriter
	Request  *http.Request
	Log      *Log
	Server   *Server
}

// Set method to set the following properties
func (c *Resource) Set(vars map[string]string, w http.ResponseWriter, r *http.Request, l *Log, s *Server) {
	c.Vars = mux.Vars(r)
	c.Response = w
	c.Request = r
	c.Log = l
	c.Server = s
}

// SetContentType method to set the content type
func (c *Resource) SetContentType(ctype string) {
	c.Response.Header().Set("Content-Type", ctype)
}

// Init method that initialized the Resource.
func (c *Resource) Init() {}

// Get represents http.get
func (c *Resource) Get() {
	c.Response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Response, StatusMethodNotAllowed)
}

// Put represents http.put
func (c *Resource) Put() {
	c.Response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Response, StatusMethodNotAllowed)
}

// Post represents http.post
func (c *Resource) Post() {
	c.Response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Response, StatusMethodNotAllowed)
}

// Patch represents http.patch
func (c *Resource) Patch() {
	c.Response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Response, StatusMethodNotAllowed)
}

// Delete represents http.delete
func (c *Resource) Delete() {
	c.Response.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Response, StatusMethodNotAllowed)
}

// Deinit method that finalizes the Resource.
func (c *Resource) Deinit() {}
