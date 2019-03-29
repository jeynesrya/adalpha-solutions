package utils

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

var (
	url    = "https://markets.ft.com/data/funds/tearsheet/summary?s=%s"
	parent = ".mod-tearsheet-overview__quote__bar"
	child  = ".mod-ui-data-list__value"
)

func GetIsinPrice(isin string) float64 {

	c := colly.NewCollector(colly.AllowedDomains("markets.ft.com"))
	var price float64
	var err error

	c.OnHTML(parent, func(e *colly.HTMLElement) {

		e.ForEach(child, func(i int, elem *colly.HTMLElement) {
			// The first element is the price
			if i == 0 {
				price, err = strconv.ParseFloat(elem.Text, 64)
				if err != nil {
					log.Fatal("Broken")
				}
			}
		})
	})

	err = c.Visit(fmt.Sprintf(url, isin))
	if err != nil {
		fmt.Println(err)
	}

	return price
}
