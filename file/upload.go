package file

import (
	"cdn/db"
	"cdn/structs"
	"cdn/util"
	"github.com/thanhpk/randstr"
	"os"
)

type RandomFileName string

// Upload upload a file to a database
func Upload(file structs.File, database db.SqlDatabase, site structs.Site) {
	upload(file, GenerateFileName(file), database, site)
}

// upload internal function to upload the file
func upload(file structs.File, newFileName RandomFileName, database db.SqlDatabase, site structs.Site) {
	// do upload of the file contents, with the generated new name
	// TODO: contents
	doUpload(database, site, newFileName, file.Contents)
}

func doUpload(database db.SqlDatabase, site structs.Site, newFileName RandomFileName, contents []byte) {
	fileName := string(newFileName)
	file, err := os.Create(util.Format("{}/{}", site.RelativeLocation, fileName))
	if err != nil {
		util.Info("Failed to create file [{}] because of {}!", fileName, err.Error())
	} else {
		// TODO: write the file name to the respective database here as well
		_, err := file.Write(contents)
		if err != nil {
			util.Info("Could not write contents to file [{}] at {}!", fileName, site.Name)
		} else {
			util.Info("Successfully wrote file [{}] to {}!", fileName, site.Name)
		}
	}
}

// GenerateFileName generate an internal file name for the file to upload
func GenerateFileName(file structs.File) RandomFileName {
	seq := randstr.Hex(8)
	name := util.Format("{}.{}", seq, file.Extension)
	return RandomFileName(name)
}