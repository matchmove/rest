# logs
--
    import "github.com/matchmove/rest/logs"


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
	// DateFormat defines the log date format
	DateFormat = time.RFC3339
)
```

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
}
```

Log represents information about a rest server log.

#### func  New

```go
func New() *Log
```
New creates new instance of Log

#### func (*Log) Count

```go
func (l *Log) Count() int
```
Count returns number of inserted logs

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

#### func (*Log) LastEntry

```go
func (l *Log) LastEntry() Entry
```
LastEntry returns the last inserted log

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

#### func (*Log) ThrowFatalTest

```go
func (l *Log) ThrowFatalTest(msg string)
```
ThrowFatalTest allows Fatal to be testable
