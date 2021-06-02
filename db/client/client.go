package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// InitDB initializes the database
func InitDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", "user=postgres dbname=simple_bank password=adeleye sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	return db
}
