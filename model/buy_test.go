package model_test

import (
	"testing"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func TestBuy(t *testing.T) {
	db := GetTestDB()

	b := model.Buy{"GB00BQ1YHQ70", 2000}

	// Price in test data is 37931.0345
	err := b.NewBuy(db)
	if err != nil {
		t.Errorf("Failed to Buy;\n%s", err.Error())
	}

}

func TestCalculateNewBuy(t *testing.T) {
	b := model.Buy{"TestISIN", 1000}
	val := b.CalculateNewBuy(250.00)

	if val != 1250 {
		t.Errorf("Calculated buy of 1250 expected; got %f", val)
	}
}
