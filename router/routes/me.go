package routes

import (
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
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

			scan, _ := database.DB.Query("SELECT * FROM users WHERE id = ?", userId)

			var user structs.User
			err := scan.Scan(&user)

			if err != nil {
				functions.SendError(err.Error(), 500, w)
				return
			}

			reply := Reply{
				UserId:   user.Id,
				Username: user.Username,
				Admin:    user.Admin,
			}

			body, _ := json.Marshal(reply)
			_, _ = w.Write(body)
		},
	}
}
