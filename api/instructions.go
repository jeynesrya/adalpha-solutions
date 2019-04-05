package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
	"github.com/jeynesrya/adalpha-solutions/model"
)

func (a *Api) InitialiseInstructionRoutes() {
	a.Router.HandleFunc("/instruction/buy", a.handleBuyInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/invest", a.handleInvestInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/sell", a.handleSellInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/raise", a.handleRaiseInstruction).Methods("POST")
}

func (a *Api) handleBuyInstruction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var buy model.Buy
	err := decoder.Decode(&buy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleBuyInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
	}

	err = buy.NewBuy(a.DB)
	if err != nil {
		// This should be a utility
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleBuyInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Buy was successful."))

}

func (a *Api) handleInvestInstruction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var invest model.Invest
	err := decoder.Decode(&invest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleInvestInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	err = invest.NewInvest(a.DB)
	if err != nil {
		// This should be a utility
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleInvestInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Invest was successful."))
}

func (a *Api) handleSellInstruction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var sell model.Sell
	err := decoder.Decode(&sell)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleSellInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	err = sell.NewSell(a.DB)
	if err != nil {
		// This should be a utility
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleSellInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sell was successful."))
}

func (a *Api) handleRaiseInstruction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var raise model.Raise
	err := decoder.Decode(&raise)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleRaiseInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	err = raise.NewRaise(a.DB)
	if err != nil {
		// This should be a utility
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Error(&es.Log{
			Package:   "api",
			Method:    "handleRaiseInstruction",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Raise was successful."))
}
