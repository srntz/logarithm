package db

import (
	"database/sql"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"os"
)

var connection *sql.DB

func Connect() *sql.DB {
	if connection == nil {
		db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION_URL"))

		if err != nil {
			panic(err)
		}

		connection = db
	}

	return connection
}
