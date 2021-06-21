package router

import (
	"cdn/router/routes"
	"cdn/structs"
)

func GetRoutes(site structs.Site) []structs.Route {
	root := routes.Root(site.Url)
	upload := routes.Upload(site)
	return []structs.Route{root, upload}
}

func SetupRoutes(router Router, site structs.Site) {
	_routes := GetRoutes(site)
	for _, route := range _routes {
		router.Handle(route.Endpoint, route.Callback)
	}
}