package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var dbHandler Database

func InitDatabase() {
	db, err := sql.Open("sqlite3", "./chat.db")
	if err != nil {
		log.Fatal(err)
	}
	dbHandler = Database{db}
	dbHandler.RunMigrations()
	return
}
