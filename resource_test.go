package rest_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/matchmove/rest"
)

func TestResourceDefaultGET(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "GET")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting GET status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceDefaultPOST(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "POST")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting POST status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceDefaultPUT(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "PUT")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting PUT status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceDefaultPATCH(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "PATCH")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting PATCH status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceDefaultOPTIONS(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "OPTIONS")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting OPTIONS status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceDefaultDELETE(t *testing.T) {
	resp := GetHandlerResponse(new(rest.Resource), "DELETE")
	if body, err := ioutil.ReadAll(resp.Body); 405 != resp.StatusCode || err != nil || "" != string(body) {
		t.Errorf(
			"Expecting DELETE status to be `%d` got, `%s` and error `%v` with status `%d`",
			http.StatusMethodNotAllowed,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestResourceJSONContentType(t *testing.T) {
	s, _ := rest.NewServer("http://0.0.0.0:8899")
	s.Env = rest.ServerEnvTesting

	resp := GetServerHandlerResponse(new(MockJSONResource), "GET", s)
	contentType := resp.Header.Get("Content-Type")
	shouldBeBody := `{"foo":"bar"}`

	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || shouldBeBody != string(body) {
		t.Errorf(
			"Expecting GET contentType to be `%s` and body `%s` got, `%s` and error `%v` with status `%d`",
			contentType,
			shouldBeBody,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}
