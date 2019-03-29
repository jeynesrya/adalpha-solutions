package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func (a *Api) InitialiseInstructionRoutes() {
	a.Router.HandleFunc("/instruction/buy", a.handleBuyInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/invest", a.handleInvestInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/sell", a.handleSellInstruction).Methods("POST")
	a.Router.HandleFunc("/instruction/raise", a.handleRaiseInstruction).Methods("POST")
}

func (a *Api) handleBuyInstruction(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) handleInvestInstruction(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var i model.Invest
	err := decoder.Decode(&i)
	if err != nil {
		log.Fatal(err)
	}

	i.NewInvest(a.DB)
}

func (a *Api) handleSellInstruction(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) handleRaiseInstruction(w http.ResponseWriter, r *http.Request) {

}
