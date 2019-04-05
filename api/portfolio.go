package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
	"github.com/jeynesrya/adalpha-solutions/model"
)

func (a *Api) InitialisePortfolioRoutes() {
	a.Router.HandleFunc("/portfolio", a.handleGetPortfolio).Methods("GET")
}

func (a *Api) handleGetPortfolio(w http.ResponseWriter, r *http.Request) {
	portfolio, err := model.GetPortfolio(a.DB)

	if err != nil {
		// This should be a utility
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleGetPortfolio",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(portfolio)
}