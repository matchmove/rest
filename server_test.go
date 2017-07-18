package rest_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/matchmove/rest"
)

func ExampleServer() {
	var (
		aLog, _  = createTempFile()
		s        *rest.Server
		chanBody = make(chan string)
		err      error
	)

	const (
		waitForResponse = 100 // # of tries before considered timeout
	)

	defer func() {
		aLog.Close()
		os.Remove(aLog.Name())
	}()

	if s, err = rest.NewServer("http://0.0.0.0:8999"); err != nil {
		panic(err)
	}

	s.Env = rest.ServerEnvTesting
	s.SetRoutes(
		mux.NewRouter().StrictSlash(true),
		rest.NewRoutes().
			Add("Test", "/test2", new(Mock2Resource)).
			Add("TestId", "/test/{id}", new(MockResource)).
			Root(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, ResponseRoot)
			}).
			NotFound(rest.DefaultNotFoundRouteHandler))

	// when .Handler is used, extend it from the .Router property
	s.Handler = handlers.LoggingHandler(
		aLog,
		func(m *mux.Router) http.Handler {
			return m
		}(s.Router),
	)

	go func() {
		s.Listen()
	}()

	var fn = func(q string, c chan string) {
		var (
			request  *http.Request
			response *http.Response
		)
		//Create request with JSON body
		if request, err = http.NewRequest("GET", s.URL.String()+"/test/1"+q, strings.NewReader("")); err != nil {
			panic(err)
		}

		for i := 0; i < waitForResponse; i++ {
			time.Sleep(10 * time.Millisecond)
			response, err = http.DefaultClient.Do(request)
			if err == nil {
				break
			}
		}

		if err != nil {
			panic(fmt.Sprintf("Retry limit exceeded: %v", err))
		}
		defer response.Body.Close()
		bytesBody, _ := ioutil.ReadAll(response.Body)
		c <- string(bytesBody)
	}

	go fn("?out=a", chanBody)
	for i := 0; i < waitForResponse; i++ {
		time.Sleep(10 * time.Millisecond)
		if err == nil {
			break
		}
	}
	go fn("", chanBody)
	fmt.Println(<-chanBody)
	fmt.Println(<-chanBody)
	// Output:
	// FooBarTest
}

func TestEmptyHandler(t *testing.T) {
	if rest.EmptyHandler(nil) == nil {
		t.Error("EmptyHandler must return the same value that it accepted")
	}
}

func TestInvalidPortSetToCharacter(t *testing.T) {
	var (
		aLog, _ = createTempFile()
		s       *rest.Server
		err     error
	)

	const (
		failedListen = "Failed to start listener with error `listen tcp4: address 9876543210: invalid port`"
	)

	defer func() {
		aLog.Close()
		os.Remove(aLog.Name())
	}()

	if s, err = rest.NewServer("http://0.0.0.0:9876543210"); err != nil {
		panic(err)
	}

	s.SetRoutes(
		mux.NewRouter().StrictSlash(true),
		rest.NewRoutes().
			Root(func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprint(w, ResponseRoot)
			}))

	sChan := make(chan *rest.Server)
	go func() {
		sChan <- s

		defer func() {
			if r := recover(); fmt.Sprintf("%v", r) != failedListen {
				t.Errorf("Should encounter: |%s|, instead got, |%v|", failedListen, r)
			}
		}()

		if err = s.Listen(); err != nil {
			panic(err)
		}
	}()

	sResult := <-sChan
	time.Sleep(10 * time.Millisecond)
	if sResult.Listener() != nil {
		t.Error("Listener should not be `open`")
		sResult.Listener().Close()
	}
}

func TestFailedServeWithInvalidHandler(t *testing.T) {
	var (
		aLog, _ = createTempFile()
		s       *rest.Server
		err     error
	)

	const (
		failedListen = "accept tcp4 0.0.0.0:8899: use of closed network connection"
	)

	defer func() {
		aLog.Close()
		os.Remove(aLog.Name())
	}()

	if s, err = rest.NewServer("http://0.0.0.0:8899"); err != nil {
		panic(err)
	}

	s.SetRoutes(mux.NewRouter(), rest.Routes{})

	sChan := make(chan *rest.Server)
	go func() {
		sChan <- s
		err = s.Listen()
		if nil == err || failedListen != err.Error() {
			t.Errorf("Should encounter: %s, instead got, %v", failedListen, err)
		}
	}()

	sResult := <-sChan
	time.Sleep(10 * time.Millisecond)
	if sResult.Listener() != nil {
		sResult.Listener().Close()
	}
}

func TestFailedNewServer(t *testing.T) {
	if _, err := rest.NewServer(""); err.Error() != "parse : empty url" {
		t.Error("Expecting `parse : empty url` error, got, ", err)
	}

}

func createTempFile() (*os.File, error) {
	return ioutil.TempFile("", "")
}
