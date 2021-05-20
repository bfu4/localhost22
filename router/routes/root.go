package routes

import (
	"cdn/structs"
	"cdn/util"
	"net/http"
)

// Root the root route
func Root(hostUrl string) structs.Route {
	point := structs.Endpoint{
		Name:    "/",
		HostUrl: hostUrl,
	}
	return structs.Route{
		Endpoint: point,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(200)
			_, _ = w.Write(getData())
		},
	}
}

func getData() []byte {
	helloWorld := util.JsonObject{
		Key:   "hello",
		Value: "world",
	}

	testObject := util.JsonObject{
		Key:   "test",
		Value: "object",
	}

	return []byte(util.Stringify(helloWorld, testObject))
}
