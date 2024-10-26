package db

import (
	"database/sql"
	"log"

	"github.com/ikhmal2/sapusapu/internal/sqlQueries"
	_ "github.com/mattn/go-sqlite3"
)

const database string = "sapusapu.db"

func DBconnect() *sqlQueries.Queries {
	db, err := sql.Open("sqlite3", database)
	if err != nil {
		log.Fatal("error connecting to DB: ", err)
	}

	queries := sqlQueries.New(db)
	return queries
}
