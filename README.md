# rest
--
    import "bitbucket.org/matchmove/rest"


## Usage

```go
const (
	// NewInstanceMsg sets the message to indicate the start of the log
	NewInstanceMsg = "START"
	// EndInstanceMsg sets the message to indicate the end of the log
	EndInstanceMsg = "END"
	// LogLevelDebug defines a normal debug log
	LogLevelDebug = "DEBUG"
	// LogLevelPanic defines a panic log
	LogLevelPanic = "PANIC"
	// LogLevelFatal defines a fatal log
	LogLevelFatal = "FATAL"
)
```

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
	ConfigExt = ".yml"
)
```

```go
var (
	// EmptyHandler creates an empty pass through handler
	EmptyHandler = func(r *mux.Router) http.Handler { return r }
)
```

#### func  LoadConfig

```go
func LoadConfig(path string, out interface{}) error
```
LoadConfig creates a new instance of configuration from a file

#### func  NewRouter

```go
func NewRouter(routes Routes, s *Server) *mux.Router
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
	Level   string
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

#### func (*Log) Panic

```go
func (l *Log) Panic(v ...interface{})
```
Panic then throws a panic with the same message afterwards

#### func (*Log) Print

```go
func (l *Log) Print(v ...interface{})
```
Print a regular log

#### type Resource

```go
type Resource struct {
	Vars     map[string]string
	Response http.ResponseWriter
	Request  *http.Request
	Log      *Log
	Server   *Server
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
func (c *Resource) Set(vars map[string]string, w http.ResponseWriter, r *http.Request, l *Log, s *Server)
```
Set method to set the following properties

#### func (*Resource) SetContentType

```go
func (c *Resource) SetContentType(ctype string)
```
SetContentType method to set the content type

#### type ResourceType

```go
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

#### func  NewRoute

```go
func NewRoute(n string, p string, r ResourceType) Route
```
NewRoute creates a new route

#### func (Route) GetHandler

```go
func (route Route) GetHandler(s *Server) func(http.ResponseWriter, *http.Request)
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

#### func (*Server) Listen

```go
func (server *Server) Listen(h func(*mux.Router) http.Handler)
```
Listen initiates the handlers

#### func (*Server) Routes

```go
func (server *Server) Routes(r Routes)
```
Routes sets up the configuration of the server and creates an instance
