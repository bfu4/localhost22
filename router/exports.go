package router

import (
	"cdn/router/routes"
	"cdn/structs"
)

func GetRoutes(hostUrl string) []structs.Route {
	root := routes.Root(hostUrl)
	upload := routes.Upload(hostUrl)
	return []structs.Route{root, upload}
}

func SetupRoutes(router Router, hostUrl string) {
	_routes := GetRoutes(hostUrl)
	for _, route := range _routes {
		router.Handle(route.Endpoint, route.Callback)
	}
}