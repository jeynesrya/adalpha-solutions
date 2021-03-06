package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jeynesrya/adalpha-solutions/es"
)

var (
	exchangerateURL = "https://api.exchangeratesapi.io/latest?base=GBP&symbols=%s"
)

// CalculateGBP used to gather exchange rate and convert to GBP
func CalculateGBP(currency string, amount float64) float64 {
	// get exchange rates
	rate, err := GetExchangeRate(currency)
	if err != nil {
		return 0
	}

	gbp := amount / rate
	return gbp
}

// GetExchangeRate used to communicate with the exchange rate API and return the value
func GetExchangeRate(currency string) (float64, error) {
	exchangerateURL = GenerateExchangeRateURL(currency)

	res, err := http.Get(exchangerateURL)
	if err != nil {
		logger.Error(&es.Log{
			Package:   "utils",
			Method:    "GetExchangeRate",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})

		return 0, err
	}
	defer res.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		logger.Error(&es.Log{
			Package:   "utils",
			Method:    "GetExchangeRate",
			Message:   err.Error(),
			Timestamp: time.Now(),
		})

		return 0, err
	}

	// Get rates
	rates := result["rates"].(map[string]interface{})
	rate := rates[currency].(float64)

	return rate, nil
}

// GenerateExchangeRateURL used to build the exchange rate from the currency entered
func GenerateExchangeRateURL(currency string) string {
	return fmt.Sprintf(exchangerateURL, currency)
}
