package router

import (
	"cdn/router/routes"
	"cdn/structs"
)

func GetRoutes(hostUrl string) []structs.Route {
	root := routes.Root(hostUrl)
	return []structs.Route{root}
}