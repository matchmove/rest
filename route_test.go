package rest_test

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"bitbucket.org/matchmove/rest"
)

func TestDefaultNotFoundRoute(t *testing.T) {
	var (
		ts   *httptest.Server
		res  *http.Response
		body []byte
		err  error
	)
	ts = httptest.NewServer(http.HandlerFunc(rest.DefaultNotFoundRouteHandler))
	defer ts.Close()

	if res, err = http.Get(ts.URL); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()
	if body, err = ioutil.ReadAll(res.Body); err != nil {
		log.Fatal(err)
	}

	if "" != string(body) {
		t.Errorf(
			"DefaultNotFoundRouteHandler must return an EMPTY response; but got `%s`",
			body)
	}

	if 404 != res.StatusCode {
		t.Errorf(
			"DefaultNotFoundRouteHandler must return an `404` response; but got `%d`",
			res.StatusCode)
	}

	if rest.ContentTypeTextPlain != res.Header.Get("Content-Type") {
		t.Errorf(
			"DefaultNotFoundRouteHandler must return an `%s` response; but got `%d`",
			rest.ContentTypeTextPlain,
			res.StatusCode)
	}
}

func TestGetHandlerWithNilServer(t *testing.T) {
	defer func() {
		errMsg := "(s *Server) cannot be `nil`"
		if r := recover(); r != errMsg {
			t.Errorf("Expecting `%s`, but got, `%v`", errMsg, r)
		}
	}()

	route := rest.Route{Resource: new(MockResource)}
	route.GetHandler(nil)
}

func TestGetHandlerGET(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "GET")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMock != string(body) {
		t.Errorf(
			"Expecting GET response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMock,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerPOST(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "POST")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMockPOST != string(body) {
		t.Errorf(
			"Expecting POST response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMockPOST,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerPUT(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "PUT")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMockPUT != string(body) {
		t.Errorf(
			"Expecting PUT response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMockPUT,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerPATCH(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "PATCH")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMockPATCH != string(body) {
		t.Errorf(
			"Expecting PATCH response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMockDELETE,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerDELETE(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "DELETE")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMockDELETE != string(body) {
		t.Errorf(
			"Expecting DELETE response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMockDELETE,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerOPTIONS(t *testing.T) {
	resp := GetHandlerResponse(new(MockResource), "OPTIONS")
	if body, err := ioutil.ReadAll(resp.Body); 200 != resp.StatusCode || err != nil || ResponseMockOPTIONS != string(body) {
		t.Errorf(
			"Expecting OPTIONS response to be `%s` got, `%s` and error `%v` with status `%d`",
			ResponseMockOPTIONS,
			string(body),
			err,
			resp.StatusCode,
		)
	}
}

func TestGetHandlerDevelopmentEnv(t *testing.T) {

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	GetServerHandlerResponse(new(MockResource), "GET", &rest.Server{
		Env: rest.ServerEnvDev,
	})

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
		t.Error("Logs should be dumped but was empty")
	}
}

func TestGetSimplePattern(t *testing.T) {
	const (
		ipv4   = "(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)"
		api    = "acm|wallet|card|risk"
		hash   = "[a-f0-9]{32}"
		method = "get|put|patch|post|delete|options|header"
	)

	r := rest.Route{
		Pattern: fmt.Sprintf(
			"/source/{ipv4:%s}/validation/{api:%s}/{hash:%s}/{method:%s}/encryption",
			ipv4,
			api,
			hash,
			method),
	}

	shouldBe := "/source/{ipv4}/validation/{api}/{hash}/{method}/encryption"

	if got := r.GetSimplePattern(); shouldBe != got {
		t.Errorf("Expected `%s` got `%s`", shouldBe, got)
	}
}
