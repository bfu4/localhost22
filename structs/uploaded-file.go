package structs

// UploadedFile defines a file that has been uploaded and it's initial location
type UploadedFile struct {

	File
	Timestamp string
	Location string

}