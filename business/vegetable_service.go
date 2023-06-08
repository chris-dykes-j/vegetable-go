package business

import (
	"encoding/csv"
	"log"
	"os"
	v "practical/models"
	r "practical/persistence"
)

type VegetableService struct {
	vegetables []v.Vegetable
}

// InitializeVegetables initializes the application's in memory data for CRUD operations
func InitializeVegetables() *VegetableService {
	vegetables := r.GetOneHundredVegetables()
	return &VegetableService{vegetables}
}

// ReloadVegetables reloads the in memory data from the repository
func (vs *VegetableService) ReloadVegetables() {
	vs.vegetables = r.GetOneHundredVegetables()
}

func (vs *VegetableService) CreateVegetable(vegetable v.Vegetable) {
	vs.vegetables = append(vs.vegetables, vegetable)
}

func (vs VegetableService) ReadAllVegetables() []v.Vegetable {
	return vs.vegetables
}

func (vs VegetableService) ReadVegetableById(id int) v.Vegetable {
	return vs.vegetables[id]
}

func (vs VegetableService) UpdateVegetableById(id int, vegetable v.Vegetable) {
	vs.vegetables[id] = vegetable
}

func (vs VegetableService) DeleteVegetableById(id int) {
	vs.vegetables = append(vs.vegetables[:id], vs.vegetables[id+1:]...) // slicing wizardry...
}

func (vs VegetableService) WriteAsCsv() {
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
