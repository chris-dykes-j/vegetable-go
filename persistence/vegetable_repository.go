package persistence

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"practical/models"
)

// VegetableRepository contains a connection to the database for CRUD operations.
type VegetableRepository struct {
	conn *pgx.Conn
}

// InitializeRepository Creates a connection to the database, and initializes the repository struct.
func InitializeRepository() *VegetableRepository {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:coolpasswordbro@localhost:5432/postgres")
	if err != nil {
		log.Fatal(err)
	}

	return &VegetableRepository{
		conn: conn,
	}
}

// ReadAllVegetables reads all the rows from the database and returns an array of delicious Vegetables.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) ReadAllVegetables() []models.Vegetable {
	rows, err := vr.conn.Query(context.Background(), "SELECT * FROM vegetable ORDER BY id DESC")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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

// CreateVegetable adds a vegetable to the database of nutritious Vegetables.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) CreateVegetable(vegetable models.Vegetable) {
	_, err := vr.conn.Exec(context.Background(), `
		INSERT INTO vegetable(
			ref_date, geo, dguid, type_of_product, type_of_storage, uom, uom_id, 
			scalar_factor, scalar_id, vector, coordinate, value, status, symbol, 
			terminated, decimals
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)`,
		vegetable.RefDate, vegetable.Geo, vegetable.DguId, vegetable.TypeOfProduct,
		vegetable.TypeOfStorage, vegetable.Uom, vegetable.UomId, vegetable.ScalarFactor,
		vegetable.ScalarId, vegetable.Vector, vegetable.Coordinate, vegetable.Value,
		vegetable.Status, vegetable.Symbol, vegetable.Terminated, vegetable.Decimals)
	if err != nil {
		log.Fatal(err)
	}
}

// ReadVegetableById reads a vegetable from the database and returns that vegetable as a Vegetable struct
// Christopher Dykes, 041013556
func (vr *VegetableRepository) ReadVegetableById(id int) models.Vegetable {
	row, err := vr.conn.Query(context.Background(), "SELECT * FROM vegetable WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	var vegetable models.Vegetable
	for row.Next() {
		err = row.Scan(&vegetable.Id, &vegetable.RefDate, &vegetable.Geo, &vegetable.DguId,
			&vegetable.TypeOfProduct, &vegetable.TypeOfStorage, &vegetable.Uom, &vegetable.UomId,
			&vegetable.ScalarFactor, &vegetable.ScalarId, &vegetable.Vector, &vegetable.Coordinate,
			&vegetable.Value, &vegetable.Status, &vegetable.Symbol, &vegetable.Terminated,
			&vegetable.Decimals)
		if err != nil {
			log.Fatal(err)
		}
	}

	return vegetable
}

// UpdateVegetableById updates a vegetable in the database given an id.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) UpdateVegetableById(id int, vegetable models.Vegetable) {
	_, err := vr.conn.Exec(context.Background(), `
		UPDATE vegetable
		SET ref_date = $1, geo = $2, dguid = $3, type_of_product = $4, type_of_storage = $5, uom = $6, 
		    uom_id = $7, scalar_factor = $8, scalar_id = $9, vector = $10, coordinate = $11, value = $12,
		    status = $13, symbol = $14, terminated = $15, decimals = $16
		WHERE id = $17`,
		vegetable.RefDate, vegetable.Geo, vegetable.DguId, vegetable.TypeOfProduct,
		vegetable.TypeOfStorage, vegetable.Uom, vegetable.UomId, vegetable.ScalarFactor,
		vegetable.ScalarId, vegetable.Vector, vegetable.Coordinate, vegetable.Value,
		vegetable.Status, vegetable.Symbol, vegetable.Terminated, vegetable.Decimals, id)
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteVegetableById removes a vegetable from the database given an id.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) DeleteVegetableById(id int) {
	_, err := vr.conn.Exec(context.Background(), "DELETE FROM vegetable WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
}

// ResetVegetableTable resets the database back to its original format, from the csv file with SQL magic.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) ResetVegetableTable() {
	_, err := vr.conn.Exec(context.Background(), `
		DROP TABLE IF EXISTS vegetable;

		CREATE TABLE vegetable(
			id SERIAL PRIMARY KEY,
			ref_date VARCHAR(255),
			geo VARCHAR(255),
			dguid VARCHAR(255),
			type_of_product VARCHAR(255),
			type_of_storage VARCHAR(255),
			uom VARCHAR(255),
			uom_id VARCHAR(255),
			scalar_factor VARCHAR(255),
			scalar_id VARCHAR(255),
			vector VARCHAR(255),
			coordinate VARCHAR(255),
			value VARCHAR(255),
			status VARCHAR(255),
			symbol VARCHAR(255),
			terminated VARCHAR(255),
			decimals VARCHAR(255)
		);

		COPY vegetable(ref_date, geo, dguid, type_of_product, type_of_storage, uom, uom_id, scalar_factor, 
		    scalar_id, vector, coordinate, value, status, symbol, terminated, decimals)
			FROM 'C:\Users\Public\32100260.csv'
			DELIMITER ','
			CSV HEADER;
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// SearchVegetables Searches the database for any entries that match the given search parameters.
// Parameters: All strings, geo is the Geography of the entry, veg the type of vegetable, storage it's storage type.
// The parameter date is expected to follow the structure "yyyy-mm" or be numeric.
// Christopher Dykes, 041013556
func (vr *VegetableRepository) SearchVegetables(geo string, veg string, storage string, date string) []models.Vegetable {
	rows, err := vr.conn.Query(context.Background(), `
		SELECT * FROM vegetable
		WHERE UPPER(geo) LIKE UPPER($1) 
		  AND UPPER(type_of_product) LIKE UPPER($2) 
		  AND UPPER(type_of_storage) LIKE UPPER($3)
		  AND ref_date LIKE $4
		ORDER BY id DESC`,
		"%"+geo+"%", "%"+veg+"%", "%"+storage+"%", "%"+date+"%")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

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
