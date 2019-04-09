package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
)

// Sell struct used to store Isin and Amount
type Sell struct {
	Isin   string
	Amount float64
}

// NewSell used to communicate with the DB to initiate the sell instruction
func (s *Sell) NewSell(db *sql.DB) error {

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT amount FROM portfolio WHERE isin=$1", s.Isin).Scan(&currentValue)

	newValue, err := s.CalculateSell(currentValue)
	if err != nil {
		logger.Error(&es.Log{
			Package:   "model",
			Method:    "NewSell",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return err
	}
	// Alter portfolio based on investment price
	_, err = db.Exec("UPDATE portfolio SET amount=$1 WHERE isin=$2", newValue, s.Isin)
	if err != nil {
		logger.Error(&es.Log{
			Package:   "model",
			Method:    "NewSell",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
	}

	return err
}

// CalculateSell used to take away amount from current value and handle if it's unable to do so
func (s *Sell) CalculateSell(currentValue float64) (float64, error) {

	newValue := currentValue - s.Amount

	if newValue < 0 {
		return 0, fmt.Errorf("unable to sell units as you only have %f in your portfolio", currentValue)
	}

	return newValue, nil
}
