package db

import (
	"database/sql"
	"fmt"
	"log"

	// Used to inject postgres into sql db
	_ "github.com/lib/pq"
)

// Postgres struct used to carry around the sql.DB pointer
type Postgres struct {
	DB *sql.DB
}

// Initialise used to set up connection to Postgres
func (p *Postgres) Initialise(host string, port int, user string, password string, dbname string) {

	// todo: return error.

	connectionString :=
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, "disable")

	var err error
	// Open connection pool
	p.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	// Check connection has been made
	err = p.DB.Ping()
	if err != nil {
		log.Fatalf("Err: %s", err.Error())
	}
}
