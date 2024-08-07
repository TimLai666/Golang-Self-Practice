package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DBconnect *gorm.DB
var err error

func DD() {
	DBconnect, err = gorm.Open(sqlite.Open("database/gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
