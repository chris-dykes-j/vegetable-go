package tests

import (
	s "practical/business"
	m "practical/models"
	"testing"
)

// TestRemoveVegetable ensures that vegetables are actually removed from the list.
// Christopher Dykes, 041013556
func TestRemoveVegetable(t *testing.T) {
	vegetables := s.InitializeService()
	initialCount := len(vegetables.ReadAllVegetables())
	vegetables.DeleteVegetableById(1)
	afterDeleteCount := len(vegetables.ReadAllVegetables())

	if initialCount-1 != afterDeleteCount {
		t.Fatalf("Vegetables length = %d, want %d", afterDeleteCount, initialCount-1)
	}

	vegetables.ResetVegetables()
}

// TestAddVegetable tests if vegetables are added to the database.
// Christopher Dykes, 041013556
func TestAddVegetable(t *testing.T) {
	vegetables := s.InitializeService()
	initialCount := len(vegetables.ReadAllVegetables())
	vegetable := m.Vegetable{
		RefDate:       "test",
		Geo:           "test",
		DguId:         "test",
		TypeOfProduct: "test",
		TypeOfStorage: "test",
		Uom:           "test",
		UomId:         "test",
		ScalarFactor:  "test",
		ScalarId:      "test",
		Vector:        "test",
		Coordinate:    "test",
		Value:         "test",
		Status:        "test",
		Symbol:        "test",
		Terminated:    "test",
		Decimals:      "test",
	}
	vegetables.CreateVegetable(vegetable)
	afterCreateCount := len(vegetables.ReadAllVegetables())

	if initialCount+1 != afterCreateCount {
		t.Fatalf("Vegetables length = %d, want %d", afterCreateCount, initialCount+1)
	}

	vegetables.ResetVegetables()
}
