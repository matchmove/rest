package rest

import (
	"net/http"
)

// Resource represents an interface information about a rest resource.
type Resource interface {
	Init()

	Get(http.ResponseWriter, *http.Request)

	Put(http.ResponseWriter, *http.Request)

	Post(http.ResponseWriter, *http.Request)

	Patch(http.ResponseWriter, *http.Request)

	Delete(http.ResponseWriter, *http.Request)

	Deinit()
}
