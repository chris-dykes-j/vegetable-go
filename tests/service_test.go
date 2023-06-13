package tests

import (
	s "practical/business"
	"testing"
)

// TestRemoveVegetable ensures that vegetables are actually removed from the list.
// Christopher Dykes, 041013556
func TestRemoveVegetable(t *testing.T) {
	vegetables := s.InitializeVegetables()
	initialCount := len(vegetables.ReadAllVegetables())
	vegetables.DeleteVegetableById(0)
	afterDeleteCount := len(vegetables.ReadAllVegetables())

	if initialCount-1 != afterDeleteCount {
		t.Fatalf("Vegetables length = %d, want %d", afterDeleteCount, initialCount-1)
	}
}
