package business

import (
	"encoding/csv"
	"log"
	"os"
	v "practical/models"
	r "practical/persistence"
)

// VegetableService is a service layer struct that provides methods to interact with a list of Vegetables.
// Christopher Dykes, 041013556
type VegetableService struct {
	vegetables []v.Vegetable
}

// InitializeVegetables initializes the application's in memory data for CRUD operations.
// Returns the list of vegetables
// Christopher Dykes, 041013556
func InitializeVegetables() *VegetableService {
	vegetables := r.GetOneHundredVegetables()
	return &VegetableService{vegetables}
}

// ReloadVegetables reloads the in memory data from the repository
// Christopher Dykes, 041013556
func (vs *VegetableService) ReloadVegetables() {
	vs.vegetables = r.GetOneHundredVegetables()
}

// CreateVegetable adds a new vegetable to the list of vegetables.
// Christopher Dykes, 041013556
func (vs *VegetableService) CreateVegetable(vegetable v.Vegetable) {
	vegetable.Id = len(vs.vegetables)
	vs.vegetables = append(vs.vegetables, vegetable)
}

// ReadAllVegetables reads all vegetables currently in the application's memory.
// Returns an array of Vegetable structs.
// Christopher Dykes, 041013556
func (vs *VegetableService) ReadAllVegetables() []v.Vegetable {
	return vs.vegetables
}

// ReadVegetableById returns a Vegetable struct given its id
// Christopher Dykes, 041013556
func (vs *VegetableService) ReadVegetableById(id int) v.Vegetable {
	return vs.vegetables[id]
}

// UpdateVegetableById replaces the vegetable given its id, and an instance of a Vegetable.
// Christopher Dykes, 041013556
func (vs *VegetableService) UpdateVegetableById(id int, vegetable v.Vegetable) {
	vs.vegetables[id] = vegetable
}

// DeleteVegetableById removes a vegetable from the list in memory given its id
// Christopher Dykes, 041013556
func (vs *VegetableService) DeleteVegetableById(id int) {
	vs.vegetables = append(vs.vegetables[:id], vs.vegetables[id+1:]...)
}

// WriteAsCsv creates a new csv file from the list of vegetables in memory.
// Christopher Dykes, 041013556
func (vs *VegetableService) WriteAsCsv() {
	file, err := os.Create("files/vegetables.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err2 := file.Close()
		if err2 != nil {
			log.Fatal(err2)
		}
	}(file)

	// Christopher Dykes, 041013556
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, vegetable := range vs.vegetables {
		err := writer.Write(vegetable.ToStringArray())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
