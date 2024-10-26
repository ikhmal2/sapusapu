package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const database string = "sapusapu.db"

func DBconnect() {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal("error connecting to DB: ", err)
	}

	rows, err := db.QueryContext()
}
