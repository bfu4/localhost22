package models

import "gorm.io/gorm"

type Uploaded struct {
	gorm.Model
	OriginalName string
	Name         string
	Extension    string
	Site         string
	Id           int `gorm:"primaryKey;autoIncrement:true"`
	Uploader     User
	UploaderID   int
}
