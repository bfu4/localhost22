package router

import (
	"cdn/router/functions"
	"cdn/router/routes"
	"cdn/structs"
	"cdn/util"
	"net/http"
	"os"
)

func GetRoutes(site structs.Site) []structs.Route {
	root := routes.Root(site.Url)
	upload := routes.Upload(site)
	content := routes.Content(site.Url)
	file := routes.File(site.Url)
	remove := routes.Remove(site)
	login := routes.Login(site)
	me := routes.Me(site)

	return []structs.Route{
		root,
		upload,
		content,
		file,
		remove,
		login,
		me,
	}
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

	for _, route := range _routes {

		// note from Ali – Learning golang
		// https://hackernoon.com/dont-make-these-5-golang-mistakes-3l3x3wcw
		// 1.1 Using reference to loop iterator variable
		route := route

		router.Handle(route.Endpoint, func(writer http.ResponseWriter, request *http.Request) {
			writer.Header().Add("Access-Control-Allow-Origin", util.GetEnvironment("FRONTEND_URL"))
			writer.Header().Add("Access-Control-Allow-Credentials", "true")

			if len(route.Methods) > 0 {
				allowed := util.Contains(route.Methods, request.Method)

				if !allowed {
					writer.WriteHeader(405)
					return
				}
			}

			if route.Authenticated {
				cookie := request.Header.Get("cookie")

				wrapper := util.GetJWTWrapper()
				result, err := wrapper.ValidateToken(cookie)

				if err != nil {
					functions.SendError("You are not signed in", 401, writer)
					return
				}

				route.Callback(writer, request, result.UserId)
				return
			}

			route.Callback(writer, request, -1)
		})
	}
}
