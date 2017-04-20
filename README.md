# ReST Server

## Example:
<code>
func m1ain () {
    var (
        aLog        = createTempFile()
        h           *rest.Host
        s           *rest.Server
        err         error
    )

    if aLog, err = ioutil.TempFile("", ""); err != nil {
        log.Fatal(err)
    }

    handler = handlers.LoggingHandler(s.AccessLog, h(s.Router))
    defer s.AccessLog.Close()


    if err = aLog.Close(); err != nil {
        log.Fatal(err)
    }

    if s, err = rest.NewServer("http://0.0.0.0:1"); err != nil {
        panic(err)
    }
    
    s.SetRoutes(
        mux.NewRouter().StrictSlash(true),
        rest.NewRoutes().
            Add("Test", "/test2", new(Mock2Resource)).
            Add("TestId", "/test/{id}", new(MockResource)).
            Root(func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, ResponseRoot)
            }).
            NotFound(rest.DefaultNotFoundRouteHandler))

        s.Listen(func(m *mux.Router) http.Handler {
            channelResp = channelOK
            return m
        })

    s.Listen()
}


const (
    channelOK = "ServerOK!"
)

if h, err = rest.NewHost("http://0.0.0.0:8999"); err != nil {
    panic(err)
}

defer os.Remove(aLog.Name())

if s, err = h.NewServer(aLog.Name()); err != nil {
    panic(err)
}

s.SetRoutes(
    mux.NewRouter().StrictSlash(true),
    rest.NewRoutes().
        Add("Test", "/test2", new(Mock2Resource)).
        Add("TestId", "/test/{id}", new(MockResource)).
        Root(func(w http.ResponseWriter, r *http.Request) {
            fmt.Fprint(w, ResponseRoot)
        }).
        NotFound(rest.DefaultNotFoundRouteHandler))

    s.Listen()

// Output:
//
</code>

# rest
--
    import "bitbucket.org/matchmove/rest"


## Usage

```go
const (
	// ContentTypeJSON defines Content-Type for application/json
	ContentTypeJSON = "application/json"
	// ContentTypeTextPlain defines Content-Type for text/plain
	ContentTypeTextPlain = "text/plain"
)
```

```go
const (
	// ServerEnvDev defines the DEVELOPMENT environment
	ServerEnvDev = "DEVELOPMENT"
	// ServerEnvTesting defines the TESTING environment
	ServerEnvTesting = "TESTING"
)
```

#### func  DefaultNotFoundRouteHandler

```go
func DefaultNotFoundRouteHandler(w http.ResponseWriter, r *http.Request)
```
DefaultNotFoundRouteHandler is the initial 404 route handler

#### func  EmptyHandler

```go
func EmptyHandler(r *mux.Router) http.Handler
```
EmptyHandler creates an empty pass through handler

#### type Resource

```go
type Resource struct {
	Vars     map[string]string
	Response http.ResponseWriter
	Request  *http.Request
	Log      *logs.Log
	Route    Route
}
```

Resource represents the information about the Resource.

#### func (*Resource) Deinit

```go
func (c *Resource) Deinit()
```
Deinit method that finalizes the Resource.

#### func (*Resource) Delete

```go
func (c *Resource) Delete()
```
Delete represents http.delete

#### func (*Resource) Get

```go
func (c *Resource) Get()
```
Get represents http.get

#### func (*Resource) Init

```go
func (c *Resource) Init() bool
```
Init method that initialized the Resource. Returning false will skip executing
the method and proceed to deinit()

#### func (*Resource) Options

```go
func (c *Resource) Options()
```
Options represents http.options

#### func (*Resource) Patch

```go
func (c *Resource) Patch()
```
Patch represents http.patch

#### func (*Resource) Post

```go
func (c *Resource) Post()
```
Post represents http.post

#### func (*Resource) Put

```go
func (c *Resource) Put()
```
Put represents http.put

#### func (*Resource) SetContentType

```go
func (c *Resource) SetContentType(ctype string)
```
SetContentType method to set the content type

#### func (*Resource) SetStatus

```go
func (c *Resource) SetStatus(code int)
```
SetStatus method to set the header status code

#### type ResourceType

```go
type ResourceType interface {
	SetContentType(string)

	Init() bool

	Get()

	Put()

	Post()

	Patch()

	Options()

	Delete()

	Deinit()
	// contains filtered or unexported methods
}
```

ResourceType represents an interface information about a rest resource.

#### type Route

```go
type Route struct {
	Name     string
	Pattern  string
	Resource ResourceType
	Server   *Server
}
```

Route represents the struct of Route

#### func (Route) GetHandler

```go
func (r Route) GetHandler(s *Server) func(http.ResponseWriter, *http.Request)
```
GetHandler is the method that handles the http.HandlerFunc

#### func (Route) GetSimplePattern

```go
func (r Route) GetSimplePattern() string
```
GetSimplePattern returns the pattern without the regex rules

#### type Routes

```go
type Routes struct {
}
```

Routes represents a array/collection of Route

#### func  NewRoutes

```go
func NewRoutes() Routes
```
NewRoutes simplifies the initialization of the routes.

#### func (Routes) Add

```go
func (rs Routes) Add(name string, pattern string, c ResourceType) Routes
```
Add a new Route to the stack

#### func (Routes) NotFound

```go
func (rs Routes) NotFound(custom func(http.ResponseWriter, *http.Request)) Routes
```
NotFound overwrites the current 404 handler

#### func (Routes) Root

```go
func (rs Routes) Root(root func(http.ResponseWriter, *http.Request)) Routes
```
Root assigns the "/" handler

#### type Server

```go
type Server struct {
	URL *url.URL
	Env string

	// Handler are used in setting http activities like activity logs.
	// see: github.com/gorilla/handlers
	Handler http.Handler
	Router  *mux.Router
}
```

Server represents information about a rest server.

#### func  NewServer

```go
func NewServer(urlRaw string) (*Server, error)
```
NewServer initializes an new ReST server

#### func (*Server) Listen

```go
func (s *Server) Listen() error
```
Listen initiates the handlers

#### func (*Server) Listener

```go
func (s *Server) Listener() io.Closer
```
Listener returns the value of the server.Listen when invoked.

#### func (*Server) SetRoutes

```go
func (s *Server) SetRoutes(router *mux.Router, routes Routes) *mux.Router
```
SetRoutes set the Routes given the array of route
