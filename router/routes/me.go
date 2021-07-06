package routes

import (
	"cdn/structs"
	"encoding/json"
	"net/http"
)

type Reply struct {
	UserId int `json:"userId"`
}

func Me(site structs.Site) structs.Route {
	endpoint := structs.Endpoint{
		Name:    "/me",
		HostUrl: site.Url,
	}

	return structs.Route{
		Endpoint:      endpoint,
		Authenticated: true,
		Methods:       []string{"GET"},
		Callback: func(w http.ResponseWriter, r *http.Request, userId int) {
			result, _ := json.Marshal(Reply{userId})
			_, _ = w.Write(result)
		},
	}
}
