package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/jeynesrya/adalpha-solutions/es"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

var logger = *es.NewLogger()

// API struct used for passing around the Router and Postgres
type API struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialise method used to attach PostgresDB to API struct and initialise routes
func (a *API) Initialise(db *sql.DB) {
	a.Router = mux.NewRouter()
	a.DB = db

	// Initialise routes
	a.InitialiseInstructionRoutes()
	a.InitialisePortfolioRoutes()
}

// Run method used to set up cors for UI and run the http web server
func (a *API) Run(addr string) {
	handler := cors.Default().Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}
