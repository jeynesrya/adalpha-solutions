package model_test

import (
	"database/sql"

	"github.com/jeynesrya/adalpha-solutions/db"
)

func GetTestDB() *sql.DB {
	db := db.Postgres{}

	db.Initialise(
		"localhost",
		54321,
		"test-user",
		"test-pw",
		"test-db",
	)

	return db.DB
}
