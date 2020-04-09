package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB(psqlInfo string) {
	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {

		log.Panic(err)
	}
}
