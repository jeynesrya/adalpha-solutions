package model

import (
	"database/sql"
	"fmt"
)

type Portfolio struct {
	Isin         string
	Units        float64
	CurrentWorth float64
}

func GetPortfolio(db *sql.DB) ([]Portfolio, error) {
	// get all from portfolio table (improvement - add investor to db table)
	rows, err := db.Query("SELECT isin, units FROM portfolio")
	if err != nil {
		return nil, fmt.Errorf("unable to find portfolio\n%+v", err.Error())
	}

	var investorPortfolio []Portfolio
	for rows.Next() {
		var singleRow Portfolio
		if err := rows.Scan(&singleRow.Isin, &singleRow.Units); err != nil {
			return nil, err
		}
		investorPortfolio = append(investorPortfolio, singleRow)
	}

	return investorPortfolio, nil
}
