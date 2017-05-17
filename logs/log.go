package logs

import (
	"fmt"
	"os"
	"time"
)

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

var (
	osExit = os.Exit
)

// Log represents information about a rest server log.
type Log struct {
	entry []Entry
}

// Entry represents information about a rest server log entry.
type Entry struct {
	Level   string
	Message string
	Time    time.Time
}

func (l Log) getDate(t time.Time) string {
	return t.Format(DateFormat)
}

// New creates new instance of Log
func New() *Log {
	var log Log
	log.entry = make([]Entry, 1)
	log.entry[0] = Entry{
		Message: NewInstanceMsg,
		Time:    time.Now(),
	}
	return &log
}

func (l *Log) addEntry(level string, v ...interface{}) {
	l.entry = append(
		l.entry,
		Entry{
			Level:   level,
			Message: fmt.Sprint(v...),
			Time:    time.Now(),
		},
	)
}

// Print a regular log
func (l *Log) Print(v ...interface{}) {
	l.addEntry(LogLevelDebug, v...)
}

// Panic then throws a panic with the same message afterwards
func (l *Log) Panic(v ...interface{}) {
	l.addEntry(LogLevelPanic, v...)
	panic(fmt.Sprint(v...))
}

// ThrowFatalTest allows Fatal to be testable
func (l *Log) ThrowFatalTest(msg string) {
	defer func() { osExit = os.Exit }()
	osExit = func(int) {}
	l.Fatal(msg)
}

// Fatal is equivalent to Print() and followed by a call to os.Exit(1)
func (l *Log) Fatal(v ...interface{}) {
	l.addEntry(LogLevelFatal, v...)
	l.Dump()
	osExit(1)
}

// LastEntry returns the last inserted log
func (l *Log) LastEntry() Entry {
	return l.entry[len(l.entry)-1]
}

// Count returns number of inserted logs
func (l *Log) Count() int {
	return len(l.entry)
}

// Dump will print all the messages to the io.
func (l *Log) Dump() {
	l.addEntry("", EndInstanceMsg)

	len := len(l.entry)
	for i := 0; i < len; i++ {
		fmt.Printf("%s\t%s\t%s\n", l.getDate(l.entry[i].Time), l.entry[i].Level, l.entry[i].Message)
	}
}
