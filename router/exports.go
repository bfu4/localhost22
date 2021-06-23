package router

import (
	"cdn/router/routes"
	"cdn/structs"
	"cdn/util"
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

	siteDirectory := site.RelativeLocation + "/content"
	_, err := os.Stat(siteDirectory)

	if os.IsNotExist(err) {
		err = os.Mkdir(siteDirectory, 0755)

		if err != nil {
			util.Warn("Could not create content directory {}!.", err.Error())
		}
	}

	jwtSecret, exists := os.LookupEnv("JWT_SECRET")

	if !exists {
		util.Fatal("No JWT_SECRET environment variable found.")
	}

	for _, route := range _routes {

		// note from Ali – Learning golang
		// https://hackernoon.com/dont-make-these-5-golang-mistakes-3l3x3wcw
		// 1.1 Using reference to loop iterator variable
		route := route

		router.Handle(route.Endpoint, func(writer http.ResponseWriter, request *http.Request) {
			if route.Authenticated {
				auth := request.Header.Get("Authorization")

				token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
					return []byte(jwtSecret), nil
				})

				if err != nil || !token.Valid {
					writer.WriteHeader(403)
					return
				}
			}

			route.Callback(writer, request)
		})
	}
}
