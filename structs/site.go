package structs

import (
	"cdn/util"
	"github.com/go-chi/chi"
	"net/http"
)

type Site struct {
	Name             string
	RelativeLocation string
	Url              string
	Port		     string
}

// Listen start listening on the site's port with the specified router router
func (site Site) Listen(router chi.Router) {
	defer func(handler http.Handler) {
		err := http.ListenAndServe(":" + site.Port, handler)
		if err != nil {
			util.Fatal("Failed to start server on port {}! {}", site.Port, err.Error())
		} else {
			util.Info("Started the server on port {}!", site.Port)
		}
	}(router)
}
