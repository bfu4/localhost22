package router

import (
	"cdn/router/routes"
	"cdn/structs"
	"os"
)

func GetRoutes(site structs.Site) []structs.Route {
	root := routes.Root(site.Url)
	upload := routes.Upload(site)
	content := routes.Content(site.Url)
	file := routes.File(site.Url)
	return []structs.Route{root, upload, content, file}
}

func SetupRoutes(router Router, site structs.Site) {
	_routes := GetRoutes(site)
	_ = os.Mkdir(site.RelativeLocation + "/content", 0755)
	for _, route := range _routes {
		router.Handle(route.Endpoint, route.Callback)
	}
}