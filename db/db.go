package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

type DB struct {
	*gorm.DB
}

func Connect(driver string, conn string) *DB {
	db, err := gorm.Open(driver, conn)
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}
