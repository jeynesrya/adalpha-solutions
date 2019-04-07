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

// todo: Comment
type Api struct {
	Router *mux.Router
	DB     *sql.DB
}

// todo: Comment
func (a *Api) Initialise(db *sql.DB) {
	a.Router = mux.NewRouter()
	a.DB = db

	// Initialise routes
	a.InitialiseInstructionRoutes()
	a.InitialisePortfolioRoutes()
}

// todo: Comment
func (a *Api) Run(addr string) {
	handler := cors.Default().Handler(a.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}
