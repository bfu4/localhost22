package routes

import (
	"bytes"
	"cdn/db"
	cdnFile "cdn/file"
	"cdn/structs"
	"cdn/util"
	"io"
	"mime"
	"net/http"
)

// Upload the upload route
// The entire procedure may be tested via curl using:
// `curl -i -F file=@"$FILE".png localhost:8080/upload -v`
func Upload(site structs.Site) structs.Route {
	point := structs.Endpoint{
		Name:    "/upload",
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

			// Get the file from key `file`
			file, handler, err := r.FormFile("file")
			if err != nil {
				util.Info(err.Error())
				w.WriteHeader(400)
				return
			}

			// Write the file into the byte buffer
			bz := bytes.NewBuffer(nil)
			_, err = io.Copy(bz, file)

			// If it fails, return internal server error
			if err != nil {
				print(err.Error())
				w.WriteHeader(500)
				return
			}
			contentType := http.DetectContentType(bz.Bytes())
			ext, _ := mime.ExtensionsByType(contentType)

			uploadFile := structs.File{
				Name:      handler.Filename,
				Size:      uint16(bz.Len()),
				Type:      contentType,
				Extension: ext[0],
				Contents:  bz.Bytes(),
			}
			// Write file into uploaded content folder
			cdnFile.Upload(uploadFile, db.GetGlobalDatabase(), site)
			w.WriteHeader(200)

			defer file.Close()
		},
	}
}