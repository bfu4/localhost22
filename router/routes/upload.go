package routes

import (
	"bytes"
	"cdn/db"
	cdnFile "cdn/file"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/structs/models"
	"cdn/util"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// Upload the upload route
// The entire procedure may be tested via curl using:
// `curl -i -F file=@"$FILE" -F site=site -H 'Authorization: jwtToken' localhost:8080/upload -v`
func Upload(site structs.Site) structs.Route {
	point := structs.Endpoint{
		Name:    "/upload",
		HostUrl: site.Url,
	}

	return structs.Route{
		Endpoint:      point,
		Authenticated: true,
		Methods:       []string{"POST"},
		Callback: func(w http.ResponseWriter, r *http.Request, userId int) {
			_ = r.ParseMultipartForm(util.DefaultFormMaxMem)

			file, handler, err := r.FormFile("file")
			if err != nil {
				functions.SendError(err.Error(), 400, w)
				return
			}

			database := db.GetGlobalDatabase()

			// Write the file into the byte buffer
			bz := bytes.NewBuffer(nil)
			_, err = io.Copy(bz, file)

			// If it fails, return internal server error
			if err != nil {
				functions.SendError(err.Error(), 500, w)
				return
			}

			contentType := http.DetectContentType(bz.Bytes())
			split := strings.Split(handler.Filename, ".")
			splitLen := len(split)
			var extension string

			// Make sure there is not a failure for files absent of extensions
			if splitLen == 1 {
				extension = " "
			} else {
				extension = "." + split[splitLen-1]
			}

			uploadFile := structs.File{
				Name:      handler.Filename,
				Size:      uint16(bz.Len()),
				Type:      contentType,
				Extension: extension,
				Contents:  bz.Bytes(),
			}

			newFileName := cdnFile.GenerateFileName(uploadFile)

			fileName := newFileName.Name + uploadFile.Extension
			err = os.WriteFile(
				util.Format("{}/content/{}", site.RelativeLocation, fileName),
				uploadFile.Contents,
				0755,
			)

			if err != nil {
				msg := err.Error()
				util.Info("Failed to create file [{}] because of {}!", fileName, msg)
				functions.SendError(msg, 500, w)
				return
			}

			uploaded := models.Uploaded{
				OriginalName: uploadFile.Name,
				Name:         newFileName.Name,
				Extension:    uploadFile.Extension,
				Site:         site.Name,
				UploaderID:   userId,
			}

			database.Create(&uploaded)

			w.WriteHeader(200)

			_, _ = w.Write([]byte(util.Stringify(
				util.JsonObject{
					Key:   "file",
					Value: newFileName.Name + extension,
				}),
			))

			defer func(file multipart.File) {
				_ = file.Close()
			}(file)
		},
	}
}
