package routes

import (
	"bytes"
	"cdn/db"
	cdnFile "cdn/file"
	"cdn/structs"
	"cdn/util"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// Upload the upload route
// The entire procedure may be tested via curl using:
// `curl -i -F file=@"$FILE" -F site=site -F user=user -Fpassword=password localhost:8080/upload -v`
func Upload(site structs.Site)	 structs.Route {
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
				_, _ = w.Write([]byte(util.Stringify(util.JsonObject{Key: "error", Value: "invalid password"})))
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
			split := strings.Split(handler.Filename, ".")
			var extension string
			// Make sure there is not a failure for files absent of extensions
			if len(split) == 1 {
				extension = " "
			} else {
				extension = "." + split[len(split) - 1]
			}
			uploadFile := structs.File{
				Name:      handler.Filename,
				Size:      uint16(bz.Len()),
				Type:      contentType,
				Extension: extension,
				Contents:  bz.Bytes(),
			}

			// Write file into uploaded content folder
			randomFile := cdnFile.Upload(uploadFile, db.GetGlobalDatabase(), site)
			w.WriteHeader(200)

			_, _ = w.Write([]byte(util.Stringify(util.JsonObject{
				Key:   "file",
				Value: randomFile.Name + extension,
			})))

			defer func(file multipart.File) {
				_ = file.Close()
			}(file)
		},
	}
}
