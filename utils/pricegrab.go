package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

var (
	url        = "https://markets.ft.com/data/funds/tearsheet/summary?s=%s"
	parent     = ".mod-tearsheet-overview__quote__bar"
	childLabel = ".mod-ui-data-list__label"
	childValue = ".mod-ui-data-list__value"
)

func GetIsinPrice(isin string) (float64, error) {

	c := colly.NewCollector(colly.AllowedDomains("markets.ft.com"))
	var price float64
	var err error

	c.OnHTML(parent, func(e *colly.HTMLElement) {
		var isGBX bool

		// Check for GBX in label
		e.ForEachWithBreak(childLabel, func(i int, elem *colly.HTMLElement) bool {
			if i == 0 {
				if strings.Contains(elem.Text, "GBX") {
					isGBX = true
				}

				return false
			}
			return true
		})

		e.ForEachWithBreak(childValue, func(i int, elem *colly.HTMLElement) bool {
			// The first element is the price
			if i == 0 {
				price, err = strconv.ParseFloat(elem.Text, 64)
				if err != nil {
					err = fmt.Errorf("Could not find the isin price.\n%s", err.Error())
				}

				if isGBX {
					price = price / 100
				}

				return false
			}

			return true
		})
	})

	err = c.Visit(fmt.Sprintf(url, isin))
	if err != nil {
		fmt.Println(err)
	}

	return price, err
}
