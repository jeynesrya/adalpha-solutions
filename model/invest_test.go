package model_test

import (
	"testing"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func TestInvest(t *testing.T) {
	db := GetTestDB()

	i := model.Invest{"GB00BQ1YHQ70", 2000, "GBR"}

	err := i.NewInvest(db)
	if err != nil {
		t.Errorf("Failed to Invest;\n%s", err.Error())
	}

}

func TestCalculateInvest(t *testing.T) {
	i := model.Invest{"GB00BQ1YHQ70", 200, "GBR"}

	val := i.CalculateInvest(1000, 2)

	if val != 1100 {
		t.Errorf("Calculated invest of 1100 expected; got %f", val)
	}
}
