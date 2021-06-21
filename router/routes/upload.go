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
	"os"
)


// Upload the upload route
// The entire procedure may be tested via curl using:
// `curl -i -F file=@"$FILE" -F site=site -F user=user -Fpassword=password localhost:8080/upload -v`
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

			_ = r.ParseMultipartForm(32 << 20) // 32 MB, default

			allowedUsername, _ := os.LookupEnv("ADMIN")
			allowedPassword, _ := os.LookupEnv("ADMIN_PASSWORD")

			if r.FormValue("user") != allowedUsername || r.PostForm.Get("password") != allowedPassword {
				w.Header().Add("Content-Type", "application/json")
				w.WriteHeader(400)
				_, _ =w.Write([]byte(util.Stringify(util.JsonObject{Key: "error", Value: "invalid password"})))
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