package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
)

const (
	TestResource200Root              = "FooBarTest"
	TestResource200Message           = "FooBar"
	TestResource200MessageSub        = "FooBarSub"
	TestResource200MessageWithParam1 = "FooBar1"
	TestServerPort                   = "8999"
	TestServerDomain                 = "http://0.0.0.0:"
)

type TestResource struct {
	Resource
}

func (c *TestResource) Get() {
	c.Response.WriteHeader(http.StatusOK)
	if c.Vars["id"] != "" {
		fmt.Fprintf(c.Response, TestResource200MessageWithParam1)
		return
	}

	fmt.Fprintf(c.Response, TestResource200Message)
}

type TestSubResource struct {
	Resource
}

func (c *TestSubResource) Get() {
	c.Response.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Response, TestResource200MessageSub)
}

func TestNewServer(t *testing.T) {
	lfile, err := ioutil.TempFile("", "")

	if err != nil {
		log.Fatal(err)
	}

	if err := lfile.Close(); err != nil {
		log.Fatal(err)
	}
	defer os.Remove(lfile.Name())

	cfile, fileName := new(Config).NewTempFile(
		"port: " + TestServerPort +
			"\nenvironment: TESTING" +
			"\naccesslog: " + lfile.Name())
	defer os.Remove(cfile.Name())

	var server Server

	if err := LoadConfig(fileName, &server); err != nil {
		panic(err)
	}

	server.Routes(Routes{
		NewRoute("Test", "/test", new(TestResource)),
		NewRoute("Test", "/test2", new(TestSubResource)),
		NewRoute("TestId", "/test/{id}", new(TestResource)),
	}, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, TestResource200Root)
	})

	var channelResponse string
	channelOk := "CHANNEL OK"

	go func() {
		server.Listen(func(m *mux.Router) http.Handler {
			channelResponse = channelOk
			return m
		})
	}()

	//Create request with JSON body
	request, err := http.NewRequest("GET", TestServerDomain+TestServerPort, strings.NewReader(""))

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)
	retryTilGiveUp := 100

	for i := 0; i < retryTilGiveUp; i++ {
		if err == nil {
			break
		}
		time.Sleep(10 * time.Millisecond)
		response, err = http.DefaultClient.Do(request)
	}

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("Connection response status `%d`, expecting `%d`", got, want)
	}

	if channelOk != channelResponse {
		t.Errorf("Server handler response `%v`, expecting `%v`", channelOk, channelResponse)
	}

	bContent, err := ioutil.ReadFile(lfile.Name())

	if err != nil {
		t.Errorf("Server access log must not have ERR; got `%v`", err)
	}

	if 0 == len(bContent) {
		t.Errorf("Server access log must not be empty; got `%v`", len(bContent))
	}
}

func runTestCall(t *testing.T, u string, expect string) {
	//Create request with JSON body
	request, err := http.NewRequest("GET", TestServerDomain+TestServerPort+u, strings.NewReader(""))

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Fatalf("Encountering error in request `%v`", err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		t.Fatalf("Encountering error in response `%v`", err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("Encountering HTTP status code `%d`, should be `%d`", got, want)
	}

	if got, want := string(body), expect; got != want {
		t.Errorf("Encountering response body `%s`, should be `%s`", got, want)
	}
}

func TestRootRoute(t *testing.T) {
	runTestCall(t, "/", TestResource200Root)
}

func TestSampleResource(t *testing.T) {
	runTestCall(t, "/test", TestResource200Message)
}

func TestSampleSubResource(t *testing.T) {
	runTestCall(t, "/test2", TestResource200MessageSub)
}

func TestSampleRouteWithoutUrlParamater(t *testing.T) {
	runTestCall(t, "/test/1", TestResource200MessageWithParam1)
}
