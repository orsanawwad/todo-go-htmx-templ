package database

import (
	"log"

	// "gorm.io/driver/sqlite"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func New() *Database {
	db, err := gorm.Open(sqlite.Open("todosdemo.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return &Database{DB: db}
}