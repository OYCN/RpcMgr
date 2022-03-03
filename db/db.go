package db

import (
	"webbk/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() {
	// db, err := gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})
	var h gorm.Dialector
	if config.DB_type == "sqlite" {
		h = sqlite.Open(config.DB_uri)
	} else {
		panic("DB Type Not Fount")
	}
	db, err := gorm.Open(h, &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Node{})
	db.AutoMigrate(&RpcCtx{})

	Db = db
}
