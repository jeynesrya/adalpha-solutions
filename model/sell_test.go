package model_test

import (
	"testing"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func TestSell(t *testing.T) {
	db := GetTestDB()

	s := model.Sell{"IE00B52L4369", 10}

	err := s.NewSell(db)
	if err != nil {
		t.Errorf("Failed to Sell;\n%s", err.Error())
	}
}

func TestCalculateSell(t *testing.T) {
	s := model.Sell{"GB00BQ1YHQ70", 200}

	val, err := s.CalculateSell(250)
	if err != nil {
		t.Fatalf("Calculating Sell failed;\n%s", err.Error())
	}

	if val != 50 {
		t.Errorf("Calculated Sell of 50 expected; got %f", val)
	}
}
