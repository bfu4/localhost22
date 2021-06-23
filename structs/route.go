package structs

import "net/http"

type Route struct {
	Endpoint      Endpoint
	Callback      func(w http.ResponseWriter, r *http.Request)
	Authenticated bool
}
