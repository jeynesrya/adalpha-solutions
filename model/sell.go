package model

import (
	"database/sql"
	"fmt"
)

type Sell struct {
	Isin  string
	Units float64
}

func (s *Sell) NewSell(db *sql.DB) error {

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT amount FROM portfolio WHERE isin=$1", s.Isin).Scan(&currentValue)

	// Not below zero...?
	newValue := currentValue - s.Units
	if newValue < 0 {
		return fmt.Errorf("unable to sell units as you only have %f in your portfolio", currentValue)
	}

	// Alter portfolio based on investment price
	_, err := db.Exec("UPDATE portfolio SET amount=$1 WHERE isin=$2", newValue, s.Isin)

	return err
}
