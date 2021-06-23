package routes

import (
	"cdn/db"
	cdnFile "cdn/file"
	"cdn/structs"
	"cdn/util"
	"net/http"
	"os"
)

// Remove the remove route
// Delete a file
// `curl -i -F file=name -F user=user -Fpassword=password localhost:8080/delete -v`
func Remove(site structs.Site) structs.Route {
	point := structs.Endpoint{
		Name:    "/remove",
		HostUrl: site.Url,
	}

	return structs.Route{
		Endpoint: point,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")

			// Check for post request
			if r.Method != "POST" {
				w.WriteHeader(405)
				return
			}

			_ = r.ParseMultipartForm(32 << 20) // 32 MB, default

			allowedUsername, _ := os.LookupEnv("ADMIN")
			allowedPassword, _ := os.LookupEnv("ADMIN_PASSWORD")

			if r.FormValue("user") != allowedUsername || r.PostForm.Get("password") != allowedPassword {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(400)
				_, _ = w.Write([]byte(util.Stringify(util.JsonObject{Key: "error", Value: "invalid password"})))
				return
			}

			file := r.FormValue("file")

			if file == "" {
				w.WriteHeader(400)
				w.Header().Add("Content-Type", "application/json")
				_, _ = w.Write([]byte(util.Stringify(util.JsonObject{Key: "error", Value: "missing a file to remove"})))
				return
			}

			// Write the file into the byte buffer

			// Write file into uploaded content folder
			w.WriteHeader(200)
			cdnFile.Remove(file, site, db.GetGlobalDatabase())
		},
	}
}
