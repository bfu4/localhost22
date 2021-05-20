package routes

import (
	"cdn/structs"
	"cdn/util"
	"net/http"
)

// Upload the upload route
func Upload(hostUrl string) structs.Route {
	point := structs.Endpoint{
		Name:    "/upload",
		HostUrl: hostUrl,
	}
	return structs.Route{
		Endpoint: point,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				// Check
				w.WriteHeader(405)
				return
			}

			// Suppress
			_ = r.ParseForm()

			// Write
			w.WriteHeader(200)
			_, _ = w.Write(getUploadData(r.Form.Get("data")))
		},
	}
}

// Todo actually do this
func getUploadData(formData string) []byte {
	data := util.JsonObject{
		Key:   "posted",
		Value: formData,
	}
	return []byte(util.Stringify(data))
}
