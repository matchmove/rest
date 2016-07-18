package rest

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"testing"
	"time"
)

func TestNewLog(t *testing.T) {
	l := NewLog()
	lastEntry := l.Entry[len(l.Entry)-1]

	if NewInstanceMsg != lastEntry.Message {
		t.Errorf(
			"Expected first entry.Message to be `%v`, got `%v`",
			NewInstanceMsg,
			lastEntry.Message,
		)
	}
}

func TestGetDateRFC3339(t *testing.T) {
	l := NewLog()
	now := l.getDate(time.Now())

	if _, e := time.Parse(time.RFC3339, now); e != nil {
		t.Errorf("Expected to be RFC3339 format, got `%v`", now)
	}
}

func TestPrint(t *testing.T) {
	l := NewLog()
	msg := "Test log message"
	i := struct {
		i    int
		text string
	}{
		21,
		"interface",
	}

	l.Print(msg, i)
	lastEntry := l.Entry[len(l.Entry)-1]

	if 2 != len(l.Entry) {
		t.Error("Expected 2 entries, got", len(l.Entry))
	}

	if tEntry := fmt.Sprint(msg, i); lastEntry.Message != tEntry {
		t.Errorf("Expected entry.Message to be `%v`, got `%v`", tEntry, lastEntry.Message)
	}
}

func TestPanic(t *testing.T) {
	l := NewLog()
	msg := "Test log message"
	i := struct {
		i    int
		text string
	}{
		21,
		"interface",
	}

	defer func() {
		if r := recover(); r != nil {
			lastEntry := l.Entry[len(l.Entry)-1]

			if 2 != len(l.Entry) {
				t.Error("Expected 2 entries, got", len(l.Entry))
			}

			if tEntry := fmt.Sprint(msg, i); lastEntry.Message != tEntry {
				t.Errorf("Expected entry.Message to be `%v`, got `%v`", tEntry, lastEntry.Message)
			}
			return
		}

		t.Error("Expected panic() with message but got `nil`")
	}()

	l.Panic(msg, i)
	t.Error("Expected panic() to be thrown")
}

func TestFatal(t *testing.T) {
	l := NewLog()
	if os.Getenv("AS_FATAL") == "1" {
		l.Fatal("Crash and burn!")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "AS_FATAL=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("Process ran with err %v, want exit status 1", err)
}

func TestDump(t *testing.T) {
	l := NewLog()
	l.Print("This is a test message")

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	l.Dump()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC   // reading our temp stdout

	if out == "" {
		t.Error("Standard output NOT to be \"\" on Dump")
	}

}
