package routes

import (
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/structs/models"
	"encoding/json"
	"net/http"
)

type Reply struct {
	UserId   int    `json:"id"`
	Username string `json:"username"`
	Admin    bool   `json:"admin"`
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
			database := db.GetGlobalDatabase()

			var user models.User
			result := database.Take(&user, "id = ?", userId)

			if result.Error != nil {
				functions.SendError(result.Error.Error(), 500, w)
				return
			}

			body, _ := json.Marshal(Reply{
				UserId:   user.Id,
				Username: user.Username,
				Admin:    user.Admin,
			})

			_, _ = w.Write(body)
		},
	}
}
