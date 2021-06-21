package routes

import (
	"bytes"
	"cdn/structs"
	"io"
	"net/http"
	"os"
)

// File retrieve a file from this route
func File(hostUrl string) structs.Route {
	point := structs.Endpoint{
		Name:    "/file",
		HostUrl: hostUrl,
	}
	return structs.Route{
		Endpoint: point,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			query := r.URL.Query()
			fileName, present := query["name"]
			if !present || len(fileName) != 1 {
				w.Header().Add("content-type", "application/json")
				_, _ = w.Write([]byte("{\"error\": \"invalid arguments\"}"))
				w.WriteHeader(400)
				return
			}
			filePath := "./content/" + fileName[0]
			if doesFileExist(filePath) {
				file, _ := os.Open(filePath)
				bz := bytes.NewBuffer(nil)
				_, _ = io.Copy(bz, file)
				contentType := http.DetectContentType(bz.Bytes())
				w.Header().Add("content-type", contentType)
				w.WriteHeader(200)
				_, _ = w.Write(bz.Bytes())
				return
			}
			w.WriteHeader(404)
		},
	}
}

func doesFileExist(file string) bool {
	_, err := os.Stat(file)
	if err != nil {
		return false
	}
	return true
}
