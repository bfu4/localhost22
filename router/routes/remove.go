package routes

import (
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/structs/models"
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
		Callback: func(w http.ResponseWriter, r *http.Request, userId int) {
			_ = r.ParseMultipartForm(util.DefaultFormMaxMem)

			file := r.FormValue("file")

			if file == "" {
				functions.SendError("missing a file to remove!", 400, w)
				return
			}

			database := db.GetGlobalDatabase()

			var resolvedFile models.Uploaded
			result := database.Take(&resolvedFile, "id = ?", file)

			if result.Error != nil {
				functions.SendError(result.Error.Error(), 500, w)
				return
			}

			database.Delete(&models.Uploaded{}, file)

			w.WriteHeader(200)
		},
	}
}
