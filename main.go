package main

import (
	"github.com/integrii/flaggy"
	"github.com/MysterGoN/worktime-recording/db"
	"github.com/MysterGoN/worktime-recording/dto"
)

func main() {
	dbInitFlag := false

	flaggy.Bool(&dbInitFlag, "i", "init", "Init SQLite DB")

	flaggy.Parse()

	if dbInitFlag {
		conn := db.Connect("sqlite3", "main.db")
		defer conn.Close()

		conn.AutoMigrate(&dto.Task{}, &dto.WorkDay{})
	} else {
		println("Database not created!")
	}
}
