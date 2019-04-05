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
	db.QueryRow("SELECT units FROM portfolio WHERE isin=$1", s.Isin).Scan(&currentValue)

	newValue, err := s.CalculateSell(currentValue)
	if err != nil {
		return err
	}
	// Alter portfolio based on investment price
	_, err = db.Exec("UPDATE portfolio SET units=$1 WHERE isin=$2", newValue, s.Isin)

	return err
}

func (s *Sell) CalculateSell(currentValue float64) (float64, error) {

	newValue := currentValue - s.Units

	if newValue < 0 {
		return 0, fmt.Errorf("unable to sell units as you only have %f in your portfolio", currentValue)
	}

	return newValue, nil
}
