package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"testing"
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

func createConfigServerFile() (*os.File, string) {
	content := []byte("port: " + TestServerPort + "\nenvironment: TESTING")

	tmp, err := ioutil.TempFile("", "server.yaml")

	if err != nil {
		log.Fatal(err)
	}

	if _, err = tmp.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmp.Close(); err != nil {
		log.Fatal(err)
	}

	oldPath := tmp.Name()

	if err := os.Rename(oldPath, oldPath+ConfigExt); err != nil {
		log.Fatal(err)
	}

	tmp, err = os.Open(oldPath + ConfigExt) // open the new file with ext

	if err != nil {
		log.Fatal(err)
	}

	return tmp, oldPath
}

func TestNewServer(t *testing.T) {
	go func() {
		file, fileName := createConfigServerFile()
		defer os.Remove(file.Name())

		NewServer(fileName, Routes{
			Route{"Root", "/", new(TestResource)},
			Route{"Test", "/test", new(TestResource)},
			Route{"TestId", "/test/{client}", new(TestResource)},
		})
	}()

	//Create request with JSON body
	request, err := http.NewRequest("GET", TestServerDomain+TestServerPort, strings.NewReader(""))

	if err != nil {
		panic(err)
	}

	response, err := http.DefaultClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("Connection response status `%d`, expecting `%d`", got, want)
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
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("Connection response status `%d`, want `%d`", got, want)
	}

	if got, want := string(body), "welcome"; got != want {
		t.Errorf("Response body `%s`, want `%s`", got, want)
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
		panic(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("Connection response status `%d`, want `%d`", got, want)
	}

	if got, want := string(body), TestResource200Message; got != want {
		t.Errorf("Response body `%s`, want `%s`", got, want)
	}
}
