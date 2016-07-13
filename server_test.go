package rest

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"fmt"
)

func TestMake(t *testing.T) {



	server := Server{}
	go func (server Server) {
		routes := Routes{
			Route{
				"index",
				http.MethodGet,
				"/",
				Index,
			},
		}

		config := Config{
			path:"",
			port:"8088",
			environment:"dev",
		}

		server.Make(config, routes)

	}(server)


	// TEST 200
	r, _ := http.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	InvokeHandler(http.HandlerFunc(Index), "/test", w, r)

	// valid test
	if got, want := w.Code, http.StatusOK; got != want {
		t.Errorf("%s: response code = %d, want %d", "TestMake", got, want)
	}

	// TEST 404
	r, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()

	InvokeHandler(http.HandlerFunc(Index), "/test", w, r)

	if got, want := w.Code, http.StatusNotFound; got != want {
		t.Errorf("%s: response code = %d, want %d", "TestMake", got, want)
	}
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test Server!\n")
}
