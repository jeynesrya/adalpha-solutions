package utils_test

import (
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/jeynesrya/adalpha-solutions/utils"
)

func TestCalculateGBPReturnsCorrectAmount(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", utils.GenerateExchangeRateURL("USD"),
		httpmock.NewStringResponder(200, `{"base":"GBP","rates":{"USD":1.3071051223},"date":"2019-04-05"}`))

	amount := utils.CalculateGBP("USD", 200)

	if amount != 153.00988159856437 {
		t.Errorf("Amount was %v; expected 153.00988159856437", amount)
	}
}
