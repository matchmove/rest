package rest

import (
	"net/http"
)

// Resource represents information about rest resource.
type Resource struct {
	url string
}

func (resource Resource) execute(out http.ResponseWriter, in *http.Request) {}
