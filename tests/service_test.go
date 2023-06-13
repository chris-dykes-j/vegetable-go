package tests

import (
	s "practical/business"
	"testing"
)

func TestRemoveVegetable(t *testing.T) {
	vegetables := s.InitializeVegetables()
	initialCount := len(vegetables.ReadAllVegetables())
	vegetables.DeleteVegetableById(0)
	afterDeleteCount := len(vegetables.ReadAllVegetables())

	if initialCount-1 != afterDeleteCount {
		t.Fatalf("Vegetables length = %d, want %d", afterDeleteCount, initialCount-1)
	}
}
