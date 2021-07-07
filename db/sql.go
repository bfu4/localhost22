package db

import (
	"cdn/structs/models"
	"cdn/util"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var database *gorm.DB

func OpenDatabase() {
	db, err := gorm.Open(
		mysql.Open(util.GetDatabaseURL()),
		&gorm.Config{},
	)

	if err != nil {
		util.Fatal(err.Error())
		return
	}

	if db == nil {
		util.Fatal("What the no db..")
		return
	}

	// AutoMigrate our models so we don't have to use
	// my trashy migrations –ali (7 jul 2021)
	_ = db.AutoMigrate(
		&models.User{},
		&models.Uploaded{},
	)

	database = db
}

// GetGlobalDatabase get the global database
func GetGlobalDatabase() *gorm.DB {
	return database
}
