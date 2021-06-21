package structs

type DatabaseEntry struct {
	OriginalName string `json:"originalName"`
	FileName      string `json:"name"`
	FileExtension string `json:"ext"`
	Site          string `json:"site"`
}
