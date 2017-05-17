package logs_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"gopkg.in/matchmove/rest.v2/logs"
)

func TestNew(t *testing.T) {
	l := logs.New()
	lastEntry := l.LastEntry()

	if logs.NewInstanceMsg != lastEntry.Message {
		t.Errorf(
			"Expected first entry.Message to be `%v`, got `%v`",
			logs.NewInstanceMsg,
			lastEntry.Message,
		)
	}
}

func TestGetDateRFC3339(t *testing.T) {
	now := time.Now().Format(logs.DateFormat)

	if _, e := time.Parse(time.RFC3339, now); e != nil {
		t.Errorf("Expected to be RFC3339 format, got `%v`", now)
	}
}

func TestPrint(t *testing.T) {
	l := logs.New()
	msg := "Test log message"
	i := struct {
		i    int
		text string
	}{
		21,
		"interface",
	}

	l.Print(msg, i)
	lastEntry := l.LastEntry()

	if c := l.Count(); 2 != c {
		t.Error("Expected 2 entries, got", c)
	}

	if tEntry := fmt.Sprint(msg, i); lastEntry.Message != tEntry {
		t.Errorf("Expected entry.Message to be `%v`, got `%v`", tEntry, lastEntry.Message)
	}
}

func TestPanic(t *testing.T) {
	l := logs.New()
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
			lastEntry := l.LastEntry()

			if c := l.Count(); 2 != c {
				t.Error("Expected 2 entries, got", c)
			}

			if logs.LogLevelPanic != lastEntry.Level {
				t.Errorf("Expected entry.Status to be `%v`, got `%v`", logs.LogLevelPanic, lastEntry.Level)
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
	l := logs.New()

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	msg := "DEMO"
	l.ThrowFatalTest(msg)

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

	fatalMsg := strings.Split(strings.Split(out, "\n")[1], "\t")

	if logs.LogLevelFatal != fatalMsg[1] {
		t.Errorf("Expected entry.Status to be `%v`, got `%v`", logs.LogLevelFatal, fatalMsg[1])
	}

	if msg != fatalMsg[2] {
		t.Errorf("Expected entry.Status to be `%v`, got `%v`", msg, fatalMsg[2])
	}
}

func TestDump(t *testing.T) {
	l := logs.New()
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
