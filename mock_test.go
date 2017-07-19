package rest_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/matchmove/rest"
)

const (
	ResponseRoot           = "FooBarTest"
	ResponseMock           = "FooBar"
	ResponseMockPOST       = "FooBarPOST"
	ResponseMockPUT        = "FooBarPUT"
	ResponseMockPATCH      = "FooBarPATCH"
	ResponseMockOPTIONS    = "FooBarOPTIONS"
	ResponseMockDELETE     = "FooBarDELETE"
	ResponseMock2          = "FooBarSub"
	ResponseMockWithParams = "FooBar1"
)

// Mock Resource
type MockResource struct {
	rest.Resource
	Param string
}

func (c *MockResource) Get() {
	c.Response.WriteHeader(http.StatusOK)
	if o := c.Request.URL.Query().Get("out"); "" != o {
		c.Param = o
	}

	if c.Param != "" {
		fmt.Fprintf(c.Response, c.Param)
		return
	}

	if c.Vars["id"] != "" {
		fmt.Fprintf(c.Response, ResponseMockWithParams)
		return
	}

	fmt.Fprintf(c.Response, ResponseMock)
}

func (c *MockResource) Post() {
	fmt.Fprintf(c.Response, ResponseMockPOST)
}

func (c *MockResource) Put() {
	fmt.Fprintf(c.Response, ResponseMockPUT)
}

func (c *MockResource) Patch() {
	fmt.Fprintf(c.Response, ResponseMockPATCH)
}

func (c *MockResource) Delete() {
	fmt.Fprintf(c.Response, ResponseMockDELETE)
}

func (c *MockResource) Options() {
	fmt.Fprintf(c.Response, ResponseMockOPTIONS)
}

// Another Mock Resource
type Mock2Resource struct {
	rest.Resource
}

func (c *Mock2Resource) Get() {
	c.Response.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Response, ResponseMock2)
}

type MockJSONResource struct {
	rest.Resource
}

func (c *MockJSONResource) Get() {
	c.Response.WriteHeader(http.StatusOK)
	c.SetContentType(rest.ContentTypeTextPlain)
	fmt.Fprintf(c.Response, `{"foo":"bar"}`)
}

func GetHandlerResponse(resource rest.ResourceType, method string) *http.Response {
	return GetServerHandlerResponse(resource, method, &rest.Server{Env: rest.ServerEnvTesting})
}

func GetServerHandlerResponse(resource rest.ResourceType, method string, s *rest.Server) *http.Response {
	route := rest.Route{ResourceInstantiator: func() rest.ResourceType { return resource }}
	handler := route.GetHandler(s)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, "http://0.0.0.0/foo", nil)
	handler(w, req)
	return w.Result()
}
