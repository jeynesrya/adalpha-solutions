package model

import (
	"database/sql"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
)

type Buy struct {
	Isin   string
	Amount float64
}

func (b *Buy) NewBuy(db *sql.DB) error {

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT amount FROM portfolio WHERE isin=$1", b.Isin).Scan(&currentValue)

	newValue := b.CalculateNewBuy(currentValue)

	// Alter portfolio based on investment price
	_, err := db.Exec("UPDATE portfolio SET amount=$1 WHERE isin=$2", newValue, b.Isin)
	if err != nil {
		logger.Error(&es.Log{
			Package:   "model",
			Method:    "NewBuy",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
	}

	return err
}

func (b *Buy) CalculateNewBuy(currentValue float64) float64 {
	return b.Amount + currentValue
}
