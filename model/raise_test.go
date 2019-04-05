package model_test

import (
	"testing"

	"github.com/jeynesrya/adalpha-solutions/model"
)

func TestRaise(t *testing.T) {
	db := GetTestDB()

	r := model.Raise{"IE00B52L4369", 10, "GBR"}

	err := r.NewRaise(db)
	if err != nil {
		t.Errorf("Failed to Raise;\n%s", err.Error())
	}
}

func TestCalculateRaise(t *testing.T) {
	// todo: multiple with failures
	r := model.Raise{"GB00BQ1YHQ70", 1500, "GBR"}

	val, err := r.CalculateRaise(2000, 2)
	if err != nil {
		t.Fatalf("Calculating Raise failed;\n%s", err.Error())
	}

	if val != 1250 {
		t.Errorf("Calculated Raise of 1250 expected; got %f", val)
	}
}
