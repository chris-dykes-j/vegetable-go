package persistence

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"practical/models"
)

// GetOneHundredVegetables reads rows from the csv file and returns an array of delicious Vegetables.
// Christopher Dykes, 041013556
func GetOneHundredVegetables() []models.Vegetable {
	file, err := os.Open("files/32100260.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	var vegetables []models.Vegetable
	reader := csv.NewReader(file)
	reader.LazyQuotes = true
	_, _ = reader.Read()
	for i := 0; i <= 100; i++ { // go's version of 'while true'
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		vegetables = append(vegetables, models.Vegetable{
			RefDate:       line[0],
			Geo:           line[1],
			DguId:         line[2],
			TypeOfProduct: line[3],
			TypeOfStorage: line[4],
			Uom:           line[5],
			UomId:         line[6],
			ScalarFactor:  line[7],
			ScalarId:      line[8],
			Vector:        line[9],
			Coordinate:    line[10],
			Value:         line[11],
			Status:        line[12],
			Symbol:        line[13],
			Terminated:    line[14],
			Decimals:      line[15],
		})
	}
	return vegetables
}
