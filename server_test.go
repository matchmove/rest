package rest

import (
	"testing"
	"net/http"
	"strings"
	"io/ioutil"
)

func TestNewServer(t *testing.T){
	go func () {
		routes := Routes{
			Route{ "ClientResource", "/clients", new(ClientResource) },
			Route{ "ClientResource", "/clients/{client}", new(ClientResource) },
		}
		NewServer("/Users/home/Code/Go/src/bitbucket.org/matchmove/rest/server",routes)
	}()
}

func TestSampleRouteWithUrlParamater(t *testing.T) {

	request, _ := http.NewRequest("GET", "http://localhost:9999/clients/test", strings.NewReader("")) //Create request with JSON body
	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("%s: response status = %d, want %d", "NewServer", got, want)
	}

	if got, want := string(body), "welcome"; got != want {
		t.Errorf("%s: response body = %s, want %s", "NewServer", got, want)
	}
}


func TestSampleRouteWithoutUrlParamater(t *testing.T) {

	request, _ := http.NewRequest("GET", "http://localhost:9999/clients", strings.NewReader("")) //Create request with JSON body
	response, _ := http.DefaultClient.Do(request)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	if got, want := response.StatusCode, http.StatusOK; got != want {
		t.Errorf("%s: response status = %d, want %d", "NewServer", got, want)
	}

	if got, want := string(body), "welcome"; got != want {
		t.Errorf("%s: response body = %s, want %s", "NewServer", got, want)
	}
}


