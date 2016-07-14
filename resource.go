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
	Set(map[string]string, http.ResponseWriter, *http.Request)

	Init()

	Get()

	Put()

	Post()

	Patch()

	Delete()

	Deinit()
}

// Resource represents the information about the Resource.
// need to add 3 properties:
// - Vars  map[string]string
// - Write http.ResponseWriter
// - Read  *http.Request
type Resource struct {
	Vars  map[string]string
	Write http.ResponseWriter
	Read  *http.Request
}

// Set method to set the following properties:
// - Vars  map[string]string  	Represents the Url Variables
// - Write http.ResponseWriter	Represents the http Response
// - Read  *http.Request				Represents the http Request
func (c *Resource) Set(vars map[string]string, w http.ResponseWriter, r *http.Request) {
	c.Vars = mux.Vars(r)
	c.Write = w
	c.Read = r
}

// SetContentType method to set the content type
func (c *Resource) SetContentType(ctype string) {
	c.Write.Header().Set("Content-Type", ctype)
}

// Init method that initialized the Resource.
func (c *Resource) Init() {
	c.SetContentType(ContentTypeJSON)
}

// Get represents http.get
func (c *Resource) Get() {

}

// Put represents http.put
func (c *Resource) Put() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

// Post represents http.post
func (c *Resource) Post() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

// Patch represents http.patch
func (c *Resource) Patch() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

// Delete represents http.delete
func (c *Resource) Delete() {
	c.Write.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(c.Write, StatusMethodNotAllowed)
}

// Deinit method that finalizes the Resource.
func (c *Resource) Deinit() {

}
