package file

import (
	"cdn/structs"
	"github.com/thanhpk/randstr"
)

// RandomFile randomly generated file structure
type RandomFile struct {
	oldName   string
	Name      string
	extension string
}

// GenerateFileName generate an internal file Name for the file to upload
func GenerateFileName(file structs.File) RandomFile {
	seq := randstr.Hex(8)
	return RandomFile{file.Name, seq, file.Extension}
}
