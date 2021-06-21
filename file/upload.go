package file

import (
	"cdn/db"
	"cdn/structs"
	"cdn/util"
	"github.com/thanhpk/randstr"
	"os"
)

// RandomFile randomly generated file structure
type RandomFile struct {
	oldName string
	name string
	extension string
}

// Upload upload a file to a database
func Upload(file structs.File, database db.SqlDatabase, site structs.Site) {
	upload(file, GenerateFileName(file), database, site)
}

// upload internal function to upload the file
func upload(file structs.File, newFileName RandomFile, database db.SqlDatabase, site structs.Site) {
	doUpload(database, site, newFileName, file.Contents)
}

func doUpload(database db.SqlDatabase, site structs.Site, newFileName RandomFile, contents []byte) {
	fileName := newFileName.name + newFileName.extension
	// Write file
	_ = os.Mkdir(site.RelativeLocation + "/content/", 0755)
	err := os.WriteFile(util.Format("{}/content/{}", site.RelativeLocation, fileName), contents, 0755)
	if err != nil {
		util.Info("Failed to create file [{}] because of {}!", fileName, err.Error())
	} else {
		queryValues := util.Format("(\"{}\",\"{}\", \"{}\", \"{}\");", newFileName.oldName, newFileName.name, newFileName.extension, site.Name)
		database.Query("insert into uploaded values" + queryValues)
	}
}

// GenerateFileName generate an internal file name for the file to upload
func GenerateFileName(file structs.File) RandomFile {
	seq := randstr.Hex(8)
	return RandomFile{file.Name, seq, file.Extension}
}