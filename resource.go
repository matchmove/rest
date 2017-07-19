package rest

// Test cases are covered in server_test.go
import (
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
	"gopkg.in/matchmove/rest.v2/logs"
)

// ResourceType represents an interface information about a rest resource.
type ResourceType interface {
	set(ResourceType, map[string]string, http.ResponseWriter, *http.Request, *logs.Log, Route)

	SetContentType(string)

	Init() bool

	Get()

	Put()

	Post()

	Patch()

	Options()

	Delete()

	Done()

	Defer()
}

const (
	// ContentTypeJSON defines Content-Type for application/json
	ContentTypeJSON = "application/json"
	// ContentTypeTextPlain defines Content-Type for text/plain
	ContentTypeTextPlain = "text/plain"
)

// Resource represents the information about the Resource.
type Resource struct {
	Vars     map[string]string
	Response http.ResponseWriter
	Request  *http.Request
	Log      *logs.Log
	Route    Route
}

// Set the Resource properties
func (c *Resource) set(self ResourceType, vars map[string]string, w http.ResponseWriter, r *http.Request, l *logs.Log, rt Route) {
	c.Vars = mux.Vars(r)
	c.Response = w
	c.Request = r
	c.Log = l
	c.Route = rt

	rc := reflect.ValueOf(self)

	var metCall = func(name string, v ...reflect.Value) []reflect.Value {
		fn := rc.MethodByName(name)
		return fn.Call(v)
	}

	defer metCall("Defer")

	if false != metCall("Init")[0].Bool() {
		// Call HTTP Method using Camelcase
		metCall(r.Method[0:1] + strings.ToLower(r.Method)[1:])
	}

	metCall("Done")
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
// Returning false will skip executing the method and proceed to deinit()
func (c *Resource) Init() bool {
	return true
}

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

// Options represents http.options
func (c *Resource) Options() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Delete represents http.delete
func (c *Resource) Delete() {
	c.SetStatus(http.StatusMethodNotAllowed)
}

// Done method that finalizes the Resource
func (c *Resource) Done() {}

// Defer is triggered after all execution (including Deinit() and faulty executions)
func (c *Resource) Defer() {}
