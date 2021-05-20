package router

import (
	"cdn/structs"
	"github.com/go-chi/chi"
	"net/http"
)

type Router struct {
	chi.Router
}

// Ease of access

func (router Router) Use(middlewares ...func(http.Handler) http.Handler) {
	router.Router.Use(middlewares...)
}

func (router Router) Get(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Get(endpoint.Name, function)
}

func (router Router) Connect(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Connect(endpoint.Name, function)
}

func (router Router) Delete(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Delete(endpoint.Name, function)
}

func (router Router) Head(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Head(endpoint.Name, function)
}

func (router Router) Options(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Options(endpoint.Name, function)
}

func (router Router) Post(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Post(endpoint.Name, function)
}

func (router Router) Put(endpoint structs.Endpoint, function http.HandlerFunc) {
	router.Router.Put(endpoint.Name, function)
}

