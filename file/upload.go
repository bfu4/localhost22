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
	oldName   string
	Name      string
	extension string
}

// Upload upload a file to a database
func Upload(file structs.File, database db.SqlDatabase, site structs.Site) RandomFile {
	name := GenerateFileName(file)
	doUpload(database, site, name, file.Contents)
	return name
}

func doUpload(database db.SqlDatabase, site structs.Site, newFileName RandomFile, contents []byte) {
	fileName := newFileName.Name + newFileName.extension
	// Write file
	err := os.WriteFile(util.Format("{}/content/{}", site.RelativeLocation, fileName), contents, 0755)
	if err != nil {
		util.Info("Failed to create file [{}] because of {}!", fileName, err.Error())
	} else {
		queryValues := util.Format("(\"{}\",\"{}\", \"{}\", \"{}\");", newFileName.oldName, newFileName.Name, newFileName.extension, site.Name)
		database.Query("insert into uploaded values" + queryValues)
	}
}

// GenerateFileName generate an internal file Name for the file to upload
func GenerateFileName(file structs.File) RandomFile {
	seq := randstr.Hex(8)
	return RandomFile{file.Name, seq, file.Extension}
}
