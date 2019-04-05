package model

import (
	"database/sql"
	"fmt"

	"github.com/jeynesrya/adalpha-solutions/utils"
)

type Raise struct {
	Isin     string
	Amount   float64
	Currency string
}

func (r *Raise) NewRaise(db *sql.DB) error {
	// TODO : Check the currency (seems important from task brief)
	if r.Currency != "GBP" {
		// Alter amount
	}

	// Convert from currency to units
	isinValue, err := utils.GetIsinPrice(r.Isin)
	if err != nil {
		return err
	}

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT units FROM portfolio WHERE isin=$1", r.Isin).Scan(&currentValue)

	newValue, err := r.CalculateRaise(currentValue, isinValue)
	if err != nil {
		return err
	}

	// Alter portfolio based on investment price
	_, err = db.Exec("UPDATE portfolio SET units=$1 WHERE isin=$2", newValue, r.Isin)

	return err
}

func (r *Raise) CalculateRaise(currentValue, isinValue float64) (float64, error) {
	// This will be the same everywhere
	unitsToRaise := r.Amount / isinValue

	newValue := currentValue - unitsToRaise
	if newValue < 0 {
		currentValueWorth := currentValue * isinValue
		return 0, fmt.Errorf("unable to sell units as you only have %f in your portfolio", currentValueWorth)
	}

	return newValue, nil
}
