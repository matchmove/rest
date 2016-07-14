# rest
--
    import "bitbucket.org/matchmove/rest"


## Usage

```go
const (
	// StatusMethodNotAllowed defines the HTTP message when 405 is encountered
	StatusMethodNotAllowed = "Method Not Allowed"
	// ContentTypeJSON defines Content-Type for application/json
	ContentTypeJSON = "application/json"
)
```

```go
const (
	// ConfigExt defines the configuration extention that can be used
	ConfigExt = ".yaml"
)
```

```go
const (
	// NewInstanceMsg sets the initial message to indicate the start of the log
	NewInstanceMsg = "NEW Log Instance"
)
```

```go
var (
	// EmptyHandler creates an empty pass through handler
	EmptyHandler = func(r *mux.Router) http.Handler { return r }
)
```

#### func  NewConfig

```go
func NewConfig(path string, out interface{}) error
```
NewConfig creates a new instance of configuration from a file

#### func  NewRouter

```go
func NewRouter(routes Routes) *mux.Router
```
NewRouter set the Routes given the array of route

#### type Config

```go
type Config struct {
}
```

Config represents information about a rest config.

#### func (Config) NewTempFile

```go
func (c Config) NewTempFile(text string) (*os.File, string)
```
NewTempFile creates a configuration file

#### type Entry

```go
type Entry struct {
	Message string
	Time    time.Time
}
```

Entry represents information about a rest server log entry.

#### type Log

```go
type Log struct {
	Entry []Entry
}
```

Log represents information about a rest server log.

#### func  NewLog

```go
func NewLog() Log
```
NewLog creates new instance of Log

#### func (*Log) Dump

```go
func (l *Log) Dump()
```
Dump will print all the messages to the io.

#### func (*Log) Fatal

```go
func (l *Log) Fatal(v ...interface{})
```
Fatal is equivalent to Print() and followed by a call to os.Exit(1)

#### func (*Log) Print

```go
func (l *Log) Print(v ...interface{})
```
Print a regular log

#### type Resource

```go
type Resource struct {
	Vars  map[string]string
	Write http.ResponseWriter
	Read  *http.Request
}
```

Resource represents the information about the Resource. need to add 3
properties: - Vars map[string]string - Write http.ResponseWriter - Read
*http.Request

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
func (c *Resource) Init()
```
Init method that initialized the Resource.

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

#### func (*Resource) Set

```go
func (c *Resource) Set(vars map[string]string, w http.ResponseWriter, r *http.Request)
```
Set method to set the following properties: - Vars map[string]string Represents
the Url Variables - Write http.ResponseWriter Represents the http Response -
Read *http.Request Represents the http Request

#### func (*Resource) SetContentType

```go
func (c *Resource) SetContentType(ctype string)
```
SetContentType method to set the content type

#### type ResourceType

```go
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
```

ResourceType represents an interface information about a rest resource.

#### type Route

```go
type Route struct {
	Name     string
	Pattern  string
	Resource ResourceType
}
```

Route represents the struct of Route properties: - Name string Route name -
Pattern string Pattern or Url Pattern - Resource ResourceType

#### func (*Route) GetHandler

```go
func (route *Route) GetHandler() func(http.ResponseWriter, *http.Request)
```
GetHandler is the method that handles the http.HandlerFunc

#### type Routes

```go
type Routes []Route
```

Routes represents a array/collection of Route

#### type Server

```go
type Server struct {
	Port        string
	Environment string
	AccessLog   string

	AccessLogFile *os.File
	Router        *mux.Router
}
```

Server represents information about a rest server.

#### func  NewServer

```go
func NewServer(path string, r Routes) Server
```
NewServer sets up the configuration of the server and creates an instance

#### func (*Server) Listen

```go
func (server *Server) Listen(h func(*mux.Router) http.Handler)
```
Listen initiates the handlers
