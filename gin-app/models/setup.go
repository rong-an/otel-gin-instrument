package models

import (
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	//"github.com/glebarez/sqlite" // for windows
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})
	DB = db
	if err := DB.Use(otelgorm.NewPlugin()); err != nil {
		panic(err)
	}
}
