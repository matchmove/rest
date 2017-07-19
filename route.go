package rest

// Test cases are covered in server_test.go
import (
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
	"gopkg.in/matchmove/rest.v2/logs"
)

// Route represents the struct of Route
type Route struct {
	Name                 string
	Pattern              string
	ResourceInstantiator func() ResourceType
	Server               *Server
}

// Routes represents a array/collection of Route
type Routes struct {
	stack    []Route
	root     func(http.ResponseWriter, *http.Request)
	notFound func(http.ResponseWriter, *http.Request)
}

// DefaultNotFoundRouteHandler is the initial 404 route handler
func DefaultNotFoundRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentTypeTextPlain)
	w.WriteHeader(http.StatusNotFound)
}

// NewRoutes simplifies the initialization of the routes.
func NewRoutes() Routes {
	return Routes{
		notFound: DefaultNotFoundRouteHandler,
	}
}

// Add a new Route to the stack
func (rs Routes) Add(name string, pattern string, resourceInstantiator func() ResourceType) Routes {
	rs.stack = append(rs.stack, Route{
		Name:                 name,
		Pattern:              pattern,
		ResourceInstantiator: resourceInstantiator,
	})

	return rs
}

// Root assigns the "/" handler
func (rs Routes) Root(root func(http.ResponseWriter, *http.Request)) Routes {
	rs.root = root
	return rs
}

// NotFound overwrites the current 404 handler
func (rs Routes) NotFound(custom func(http.ResponseWriter, *http.Request)) Routes {
	rs.notFound = custom
	return rs
}

// GetSimplePattern returns the pattern without the regex rules
func (r Route) GetSimplePattern() string {
	reg, _ := regexp.Compile(`:[:()?a-zA-Z0-9\[\]\-\|\{\}\\\.]+`)

	return reg.ReplaceAllString(r.Pattern, "}")
}

// GetHandler is the method that handles the http.HandlerFunc
func (r Route) GetHandler(s *Server) func(http.ResponseWriter, *http.Request) {
	if s == nil {
		panic("(s *Server) cannot be `nil`")
	}

	return func(w http.ResponseWriter, rq *http.Request) {
		l := logs.New()

		defer func() {
			if ServerEnvTesting != s.Env {
				l.Dump()
			}
		}()

		resource := r.ResourceInstantiator()

		resource.set(resource, mux.Vars(rq), w, rq, l, r)
	}
}
