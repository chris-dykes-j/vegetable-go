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
	repository *r.VegetableRepository
}

// InitializeService creates the VegetableService for performing CRUD operations.
// Initializes the repository for database interaction.
// Returns the VegetableService
// Christopher Dykes, 041013556
func InitializeService() *VegetableService {
	return &VegetableService{r.InitializeRepository()}
}

// ReadAllVegetables reads all vegetables from the application's database.
// Returns an array of Vegetable structs.
// Christopher Dykes, 041013556
func (vs *VegetableService) ReadAllVegetables() []v.Vegetable {
	return vs.repository.ReadAllVegetables()
}

// CreateVegetable adds a new vegetable to database.
// Christopher Dykes, 041013556
func (vs *VegetableService) CreateVegetable(vegetable v.Vegetable) {
	vs.repository.CreateVegetable(vegetable)
}

// ReadVegetableById returns a Vegetable struct given its id
// Christopher Dykes, 041013556
func (vs *VegetableService) ReadVegetableById(id int) v.Vegetable {
	return vs.repository.ReadVegetableById(id)
}

// UpdateVegetableById updates the vegetable given its id, and an instance of a Vegetable.
// Christopher Dykes, 041013556
func (vs *VegetableService) UpdateVegetableById(id int, vegetable v.Vegetable) {
	vs.repository.UpdateVegetableById(id, vegetable)
}

// DeleteVegetableById removes a vegetable from the database permanently, given its id
// Christopher Dykes, 041013556
func (vs *VegetableService) DeleteVegetableById(id int) {
	vs.repository.DeleteVegetableById(id)
}

// ResetVegetables resets the entire database to it's initial state.
// Christopher Dykes, 041013556
func (vs *VegetableService) ResetVegetables() {
	vs.repository.ResetVegetableTable()
}

// WriteAsCsv creates a new csv file from the application database.
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

	for _, vegetable := range vs.repository.ReadAllVegetables() {
		err := writer.Write(vegetable.ToStringArray())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func (vs *VegetableService) SearchVegetables(geo string, veg string, storage string, date string) []v.Vegetable {
	return vs.repository.SearchVegetables(geo, veg, storage, date)
}
