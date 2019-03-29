package model

import (
	"database/sql"

	"github.com/jeynesrya/adalpha-solutions/utils"
)

type Invest struct {
	Isin     string
	Amount   float64
	Currency float64
}

func (i *Invest) NewInvest(db *sql.DB) {
	// Check the currency
	if i.Currency != "GB" {
		// Alter amount
	}

	// Convert from currency to units
	isinValue := utils.GetIsinPrice(i.Isin)

	// Get portfolio for isin
	currentValue := db.QueryRow("SELECT amount FROM portfolio WHERE id=$1", i.Isin)

	newValue := isinValue + currentValue

	// Alter portfolio based on investment price
	db.Exec("UPDATE portfolio SET $1 WHERE id=$2", newValue, i.Isin)
}
