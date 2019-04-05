package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Postgres struct {
	DB *sql.DB
}

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
