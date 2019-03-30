package model

import (
	"database/sql"
)

type Buy struct {
	Isin  string
	Units float64
}

func (b *Buy) NewBuy(db *sql.DB) error {

	// Get portfolio for isin
	var currentValue float64
	db.QueryRow("SELECT amount FROM portfolio WHERE isin=$1", b.Isin).Scan(&currentValue)

	newValue := b.Units + currentValue

	// Alter portfolio based on investment price
	_, err := db.Exec("UPDATE portfolio SET amount=$1 WHERE isin=$2", newValue, b.Isin)

	return err
}
