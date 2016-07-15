package rest

import (
	"fmt"
	"os"
	"time"
)

const (
	// NewInstanceMsg sets the message to indicate the start of the log
	NewInstanceMsg = "NEW Log Instance"
	// EndInstanceMsg sets the message to indicate the end of the log
	EndInstanceMsg = "End of Log Entry"
)

// Log represents information about a rest server log.
type Log struct {
	Entry []Entry
}

// Entry represents information about a rest server log entry.
type Entry struct {
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

// Print a regular log
func (l *Log) Print(v ...interface{}) {
	l.Entry = append(
		l.Entry,
		Entry{
			Message: fmt.Sprint(v...),
			Time:    time.Now(),
		},
	)
}

// Fatal is equivalent to Print() and followed by a call to os.Exit(1)
func (l *Log) Fatal(v ...interface{}) {
	l.Print(v...)
	l.Dump()
	os.Exit(1)
}

// Dump will print all the messages to the io.
func (l *Log) Dump() {
	l.Print(EndInstanceMsg)

	len := len(l.Entry)
	for i := 0; i < len; i++ {
		fmt.Printf("%s - %s\n", l.getDate(l.Entry[i].Time), l.Entry[i].Message)
	}
}
