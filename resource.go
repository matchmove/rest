package rest

// Test cases are covered in server_test.go
import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	// ContentTypeJSON defines Content-Type for application/json
	ContentTypeJSON = "application/json"
	// ContentTypeTextPlain defines Content-Type for text/plain
	ContentTypeTextPlain = "text/plain"
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

// SetStatus method to set the header status code
func (c *Resource) SetStatus(code int) {
	c.Response.WriteHeader(code)
}

// Init method that initialized the Resource.
func (c *Resource) Init() {}

// Get represents http.get
func (c *Resource) Get() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Put represents http.put
func (c *Resource) Put() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Post represents http.post
func (c *Resource) Post() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Patch represents http.patch
func (c *Resource) Patch() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Delete represents http.delete
func (c *Resource) Delete() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Deinit method that finalizes the Resource.
func (c *Resource) Deinit() {}
