package persistence

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"practical/models"
)

// GetAllVegetables reads all the rows from the database and returns an array of delicious Vegetables.
// Christopher Dykes, 041013556
func GetAllVegetables() []models.Vegetable {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:coolpasswordbro@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err2 := conn.Close(ctx)
		if err2 != nil {
			log.Fatal(err2)
		}
	}(conn, context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM vegetable")
	if err != nil {
		log.Fatal(err)
	}

	var vegetables []models.Vegetable
	for rows.Next() {
		var vegetable models.Vegetable
		err = rows.Scan(&vegetable.Id, &vegetable.RefDate, &vegetable.Geo, &vegetable.DguId,
			&vegetable.TypeOfProduct, &vegetable.TypeOfStorage, &vegetable.Uom, &vegetable.UomId,
			&vegetable.ScalarFactor, &vegetable.ScalarId, &vegetable.Vector, &vegetable.Coordinate,
			&vegetable.Value, &vegetable.Status, &vegetable.Symbol, &vegetable.Terminated,
			&vegetable.Decimals)

		if err != nil {
			log.Fatal(err)
		}
		vegetables = append(vegetables, vegetable)
	}
	return vegetables
}
