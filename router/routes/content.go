package routes

import (
	"cdn/db"
	"cdn/structs"
	"cdn/util"
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
			w.Header().Add("content-type", "application/json")
			w.WriteHeader(200)
			_, _ = w.Write([]byte(getAllContent()))
		},
	}
}

func getAllContent() string {
	database := db.GetGlobalDatabase()
	rows, err := database.DB.Query("select * from uploaded;")

	if err != nil {
		return util.Stringify(util.JsonObject{Key: "values", Value: "none"})
	}

	var sites = make(map[string][]structs.DatabaseEntry)
	var curr string

	for rows.Next() {
		entry := structs.DatabaseEntry{}
		_ = rows.Scan(&entry.OriginalName, &entry.FileName, &entry.FileExtension, &entry.Site)

		if curr != entry.Site {
			curr = entry.Site
		}

		sites[curr] = append(sites[curr], entry)
	}

	ret, err := json.Marshal(sites)

	if err != nil {
		return util.Stringify(util.JsonObject{Key: "values", Value: "unreadable"})
	}

	return string(ret)
}
