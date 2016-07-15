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
	TestResource200Message = "FooBar"
	TestServerPort         = "8999"
	TestServerDomain       = "http://0.0.0.0:"
)

type TestResource struct {
	Resource
}

func (c *TestResource) Get() {
	c.Write.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Write, TestResource200Message)
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

	server := NewServer(fileName, Routes{
		NewRoute("Root", "/", new(TestResource)),
		NewRoute("Test", "/test", new(TestResource)),
		NewRoute("TestId", "/test/{client}", new(TestResource)),
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

func TestSampleRouteWithUrlParamater(t *testing.T) {
	//Create request with JSON body
	request, err := http.NewRequest("GET", TestServerDomain+TestServerPort+"/test", strings.NewReader(""))

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

	if got, want := string(body), TestResource200Message; got != want {
		t.Errorf("Encountering response body `%s`, should be `%s`", got, want)
	}
}

func TestSampleRouteWithoutUrlParamater(t *testing.T) {
	//Create request with JSON body
	request, err := http.NewRequest("GET", TestServerDomain+TestServerPort+"/test", strings.NewReader(""))

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

	if got, want := string(body), TestResource200Message; got != want {
		t.Errorf("Encountering response body `%s`, should be `%s`", got, want)
	}
}
