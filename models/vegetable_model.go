package models

import (
	"fmt"
	"strconv"
)

// Vegetable struct represents a record from the csv file.
// Christopher Dykes, 041013556
type Vegetable struct {
	Id            int
	RefDate       string
	Geo           string
	DguId         string
	TypeOfProduct string
	TypeOfStorage string
	Uom           string
	UomId         string
	ScalarFactor  string
	ScalarId      string
	Vector        string
	Coordinate    string
	Value         string
	Status        string
	Symbol        string
	Terminated    string
	Decimals      string
}

// ToString returns the string representation of the Vegetable struct.
// Christopher Dykes, 041013556
func (v Vegetable) ToString() string {
	return fmt.Sprintf(
		"%d, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s",
		v.Id, v.RefDate, v.Geo, v.DguId, v.TypeOfProduct, v.TypeOfStorage, v.Uom, v.UomId, v.ScalarFactor, v.ScalarId,
		v.Vector, v.Coordinate, v.Value, v.Status, v.Symbol, v.Terminated, v.Decimals)
}

// ToStringArray returns an array of strings, for each field in the Vegetable struct.
// Christopher Dykes, 041013556
func (v Vegetable) ToStringArray() []string {
	return []string{
		strconv.Itoa(v.Id),
		v.RefDate,
		v.Geo,
		v.DguId,
		v.TypeOfProduct,
		v.TypeOfStorage,
		v.Uom,
		v.UomId,
		v.ScalarFactor,
		v.ScalarId,
		v.Vector,
		v.Coordinate,
		v.Value,
		v.Status,
		v.Symbol,
		v.Terminated,
		v.Decimals,
	}
}
