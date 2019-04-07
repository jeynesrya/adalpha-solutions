package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
)

type Portfolio struct {
	Isin         string
	Units        float64
	CurrentWorth float64
}

func GetPortfolio(db *sql.DB) ([]Portfolio, error) {
	// get all from portfolio table (improvement - add investor to db table)
	rows, err := db.Query("SELECT isin, amount FROM portfolio")
	if err != nil {
		logger.Error(&es.Log{
			Package:   "model",
			Method:    "GetPortfolio",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})
		return nil, fmt.Errorf("unable to find portfolio\n%+v", err.Error())
	}

	var investorPortfolio []Portfolio
	for rows.Next() {
		var singleRow Portfolio
		if err := rows.Scan(&singleRow.Isin, &singleRow.Units); err != nil {
			logger.Error(&es.Log{
				Package:   "model",
				Method:    "GetPortfolio",
				Message:   err.Error(),
				Timestamp: time.Now(),
			})
			return nil, err
		}
		investorPortfolio = append(investorPortfolio, singleRow)
	}

	return investorPortfolio, nil
}
