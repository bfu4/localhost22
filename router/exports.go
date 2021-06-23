package router

import (
	"cdn/router/routes"
	"cdn/structs"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

func GetRoutes(site structs.Site) []structs.Route {
	root := routes.Root(site.Url)
	upload := routes.Upload(site)
	content := routes.Content(site.Url)
	file := routes.File(site.Url)
	remove := routes.Remove(site)
	return []structs.Route{root, upload, content, file, remove}
}

func SetupRoutes(router Router, site structs.Site) {
	_routes := GetRoutes(site)
	_ = os.Mkdir(site.RelativeLocation+"/content", 0755)

	for _, route := range _routes {
		router.Handle(route.Endpoint, func(writer http.ResponseWriter, request *http.Request) {
			if route.Authenticated {
				auth := request.Header.Get("Authorization")

				jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
					token.Valid
				})
			}

			route.Callback(writer, request)
		})
	}
}
