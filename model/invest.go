package model

import (
	"database/sql"

	"github.com/jeynesrya/adalpha-solutions/utils"
)

type Invest struct {
	Isin     string
	Amount   float64
	Currency string
}

func (i *Invest) NewInvest(db *sql.DB) error {
	// TODO : Check the currency (seems important from task brief)
	if i.Currency != "GBP" {
		// Alter amount
	}

	// Convert from currency to units
	isinValue, err := utils.GetIsinPrice(i.Isin)
	if err != nil {
		return err
	}

	// This will be the same everywhere
	additionalUnits := i.Amount / isinValue

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT amount FROM portfolio WHERE isin=$1", i.Isin).Scan(&currentValue)

	newValue := additionalUnits + currentValue

	// Alter portfolio based on investment price
	_, err = db.Exec("UPDATE portfolio SET units=$1 WHERE isin=$2", newValue, i.Isin)

	return err
}
