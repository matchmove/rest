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

func (c *MockResource) Get() (status int, body interface{}) {
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
	return status, body
}

func (c *MockResource) Post() (status int, body interface{}) {
	fmt.Fprintf(c.Response, ResponseMockPOST)
	return status, body
}

func (c *MockResource) Put() (status int, body interface{}) {
	fmt.Fprintf(c.Response, ResponseMockPUT)
	return status, body
}

func (c *MockResource) Patch() (status int, body interface{}) {
	fmt.Fprintf(c.Response, ResponseMockPATCH)
	return status, body
}

func (c *MockResource) Delete() (status int, body interface{}) {
	fmt.Fprintf(c.Response, ResponseMockDELETE)
	return status, body
}

func (c *MockResource) Options() (status int, body interface{}) {
	fmt.Fprintf(c.Response, ResponseMockOPTIONS)
	return status, body
}

// Another Mock Resource
type Mock2Resource struct {
	rest.Resource
}

func (c *Mock2Resource) Get() (status int, body interface{}) {
	c.Response.WriteHeader(http.StatusOK)
	fmt.Fprintf(c.Response, ResponseMock2)
	return status, body
}

type MockJSONResource struct {
	rest.Resource
}

func (c *MockJSONResource) Get() (status int, body interface{}) {
	c.Response.WriteHeader(http.StatusOK)
	c.SetContentType(rest.ContentTypeTextPlain)
	fmt.Fprintf(c.Response, `{"foo":"bar"}`)
	return status, body
}

func GetHandlerResponse(resource rest.ResourceType, method string) *http.Response {
	return GetServerHandlerResponse(resource, method, &rest.Server{Env: rest.ServerEnvTesting})
}

func GetServerHandlerResponse(resource rest.ResourceType, method string, s *rest.Server) *http.Response {
	route := rest.Route{Resource: resource}
	handler := route.GetHandler(s)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, "http://0.0.0.0/foo", nil)
	handler(w, req)
	return w.Result()
}
