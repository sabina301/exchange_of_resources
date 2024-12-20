package db

import (
	"database/sql"
	"log"
)

var Db *sql.DB

func InitDB() {
	var err error
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=exchange_of_resources sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
