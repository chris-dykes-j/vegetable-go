package business

import (
	"encoding/csv"
	"log"
	"os"
	v "practical/models"
	r "practical/persistence"
)

// Christopher Dykes, 041013556
type VegetableService struct {
	vegetables []v.Vegetable
}

// InitializeVegetables initializes the application's in memory data for CRUD operations
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

// Christopher Dykes, 041013556
func (vs *VegetableService) CreateVegetable(vegetable v.Vegetable) {
	vegetable.Id = len(vs.vegetables)
	vs.vegetables = append(vs.vegetables, vegetable)
}

// Christopher Dykes, 041013556
func (vs *VegetableService) ReadAllVegetables() []v.Vegetable {
	return vs.vegetables
}

// Christopher Dykes, 041013556
func (vs *VegetableService) ReadVegetableById(id int) v.Vegetable {
	return vs.vegetables[id]
}

// Christopher Dykes, 041013556
func (vs *VegetableService) UpdateVegetableById(id int, vegetable v.Vegetable) {
	vs.vegetables[id] = vegetable
}

// Christopher Dykes, 041013556
func (vs *VegetableService) DeleteVegetableById(id int) {
	vs.vegetables = append(vs.vegetables[:id], vs.vegetables[id+1:]...)
}

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
