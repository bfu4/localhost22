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
		Callback: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("content-type", "application/json")
			w.Header().Add("Access-Control-Allow-Origin", "*")
			w.WriteHeader(200)
			_, _ = w.Write([]byte(getAllContent()))
		},
	}
}

type databaseEntry struct {
	FileName      string `json:"name"`
	FileExtension string `json:"ext"`
	Site          string `json:"site"`
}

func getAllContent() string {
	database := db.GetGlobalDatabase()
	rows, err := database.DB.Query("select * from uploaded;")
	if err != nil {
		return util.Stringify(util.JsonObject{Key: "values", Value: "none"})
	}
	var sites = make(map[string][]databaseEntry)
	var curr string
	for rows.Next() {
		entry := databaseEntry{}
		_ = rows.Scan(&entry.FileName, &entry.FileExtension, &entry.Site)
		if curr != entry.Site {
			curr = entry.Site
		}
		sites[curr] = append(sites[curr], entry)
	}

	ret, err := json.Marshal(sites)
	if err != nil {
		return util.Stringify(util.JsonObject{Key:"values", Value: "unreadable"})
	}
	return string(ret)
}