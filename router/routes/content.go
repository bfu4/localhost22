package routes

import (
	"cdn/db"
	"cdn/structs"
	"cdn/structs/models"
	"encoding/json"
	"net/http"
)

// Content the content route
func Content(hostUrl string) structs.Route {
	point := structs.Endpoint{
		Name:    "/content",
		HostUrl: hostUrl,
	}

	return structs.Route{
		Endpoint: point,
		Callback: func(w http.ResponseWriter, r *http.Request, userId int) {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(200)
			_, _ = w.Write([]byte(getAllContent()))
		},
	}
}

func getAllContent() string {
	database := db.GetGlobalDatabase()

	var contents []models.User
	database.Find(&contents)

	ret, _ := json.Marshal(contents)

	return string(ret)
}
