package model_test

import (
	"testing"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func TestGetPortfolio(t *testing.T) {
	db := GetTestDB()

	portfolio, err := model.GetPortfolio(db)
	if err != nil {
		t.Fatalf("Error getting portfolio elements;\n%s", err.Error())
	}

	if len(portfolio) != 6 {
		t.Errorf("Number of portfolio items should be 6; got %d", len(portfolio))
	}
}
