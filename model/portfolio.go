package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
	"github.com/jeynesrya/adalpha-solutions/utils"
)

// Portfolio struct used to store Isin, Units and Currency
type Portfolio struct {
	Isin         string
	Units        float64
	CurrentWorth float64
}

// GetPortfolio used to get all the PortfolioItems from the DB
func GetPortfolio(db *sql.DB) ([]Portfolio, error) {
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
		// Establish currentWorth
		isinValue, err := utils.GetIsinPrice(singleRow.Isin)
		if err != nil {
			logger.Error(&es.Log{
				Package:   "model",
				Method:    "GetPortfolio",
				Message:   err.Error(),
				Timestamp: time.Now(),
			})

			isinValue = 0
		}

		singleRow.CurrentWorth = isinValue * singleRow.Units
		investorPortfolio = append(investorPortfolio, singleRow)
	}

	return investorPortfolio, nil
}
