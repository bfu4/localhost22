package routes

import (
	"cdn/db"
	cdnFile "cdn/file"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/util"
	"net/http"
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
		Endpoint:      point,
		Methods:       []string{"POST"},
		Authenticated: true,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			_ = r.ParseMultipartForm(util.DefaultFormMaxMem)

			file := r.FormValue("file")

			if file == "" {
				functions.SendError("missing a file to remove!", 400, w)
				return
			}

			// Write the file into the byte buffer

			// Write file into uploaded content folder
			w.WriteHeader(200)
			cdnFile.Remove(file, site, db.GetGlobalDatabase())
		},
	}
}
