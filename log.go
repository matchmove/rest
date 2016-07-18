package rest

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
)

// Log represents information about a rest server log.
type Log struct {
	Entry []Entry
}

// Entry represents information about a rest server log entry.
type Entry struct {
	Level   string
	Message string
	Time    time.Time
}

func (l Log) getDate(t time.Time) string {
	return t.Format(time.RFC3339)
}

// NewLog creates new instance of Log
func NewLog() Log {
	var log Log
	log.Entry = make([]Entry, 1)
	log.Entry[0] = Entry{
		Message: NewInstanceMsg,
		Time:    time.Now(),
	}
	return log
}

func (l *Log) addEntry(level string, v ...interface{}) {
	l.Entry = append(
		l.Entry,
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

// Fatal is equivalent to Print() and followed by a call to os.Exit(1)
func (l *Log) Fatal(v ...interface{}) {
	l.Print(v...)
	l.Dump()
	os.Exit(1)
}

// Dump will print all the messages to the io.
func (l *Log) Dump() {
	l.addEntry("", EndInstanceMsg)

	len := len(l.Entry)
	for i := 0; i < len; i++ {
		fmt.Printf("%s\t%s\t%s\n", l.getDate(l.Entry[i].Time), l.Entry[i].Level, l.Entry[i].Message)
	}
}
