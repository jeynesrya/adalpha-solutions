package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jeynesrya/adalpha-solutions/db"

	"github.com/jeynesrya/adalpha-solutions/api"
)

func main() {
	fmt.Println("Starting application...")

	// Initialise application
	db := db.Postgres{}
	port, err := strconv.Atoi(os.Getenv("PG_PORT"))
	if err != nil {
		log.Fatal("Unable to convert port to integer.")
	}

	// todo add logging
	db.Initialise(
		os.Getenv("PG_HOST"),
		port,
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DB_NAME"),
	)

	a := api.API{}
	a.Initialise(db.DB)

	// Run application
	a.Run(":8080")
}
