package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var Db *sql.DB

func InitDB() {
	var err error
	db, err := sql.Open("postgres", "user=sabina301 password=1 dbname=exchange_of_resources sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
