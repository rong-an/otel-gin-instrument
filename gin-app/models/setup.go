package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"github.com/glebarez/sqlite"  // for windows
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})
	DB = db
}
