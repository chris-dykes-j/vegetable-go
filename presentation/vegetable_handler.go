package presentation

import (
	"html/template"
	"log"
	"net/http"
	"path"
	s "practical/business"
	m "practical/models"
	"strconv"
)

type VegetableHandler struct {
	service *s.VegetableService
}

func InitializeHandler(service *s.VegetableService) *VegetableHandler {
	return &VegetableHandler{service}
}

func (vh *VegetableHandler) IndexHandler(writer http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("presentation/views/index.gohtml"))
	err := tmpl.Execute(writer, vh.service.ReadAllVegetables())
	if err != nil {
		log.Fatal(err)
	}
}

func (vh *VegetableHandler) EditHandler(writer http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/edit.gohtml"))
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = editTmpl.Execute(writer, vh.service.ReadVegetableById(id))
	if err != nil {
		log.Fatal(err)
	}
}

func (vh VegetableHandler) AddHandler(writer http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/add.gohtml"))
	err := editTmpl.Execute(writer, m.Vegetable{})
	if err != nil {
		log.Fatal(err)
	}
}

func (vh *VegetableHandler) DeleteHandler(writer http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}
	vh.service.DeleteVegetableById(id)
}

func (vh *VegetableHandler) ReloadHandler(writer http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	vh.service.ReloadVegetables()
}

func (vh VegetableHandler) UpdateHandler(writer http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(writer, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(writer, "Can't parse form", http.StatusInternalServerError)
		return
	}

	vegetable := m.Vegetable{
		RefDate:       r.FormValue("RefDate"),
		Geo:           r.FormValue("Geo"),
		DguId:         r.FormValue("DguId"),
		TypeOfProduct: r.FormValue("TypeOfProduct"),
		TypeOfStorage: r.FormValue("TypeOfStorage"),
		Uom:           r.FormValue("Uom"),
		UomId:         r.FormValue("UomId"),
		ScalarFactor:  r.FormValue("ScalarFactor"),
		ScalarId:      r.FormValue("ScalarId"),
		Vector:        r.FormValue("Vector"),
		Coordinate:    r.FormValue("Coordinate"),
		Value:         r.FormValue("Value"),
		Status:        r.FormValue("Status"),
		Symbol:        r.FormValue("Symbol"),
		Terminated:    r.FormValue("Terminated"),
		Decimals:      r.FormValue("Decimals"),
	}

	vh.service.UpdateVegetableById(id, vegetable)
	vh.IndexHandler(writer, r)
}

func (vh VegetableHandler) CreateHandler(writer http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(writer, "Can't parse form", http.StatusInternalServerError)
		return
	}

	vegetable := m.Vegetable{
		RefDate:       r.FormValue("RefDate"),
		Geo:           r.FormValue("Geo"),
		DguId:         r.FormValue("DguId"),
		TypeOfProduct: r.FormValue("TypeOfProduct"),
		TypeOfStorage: r.FormValue("TypeOfStorage"),
		Uom:           r.FormValue("Uom"),
		UomId:         r.FormValue("UomId"),
		ScalarFactor:  r.FormValue("ScalarFactor"),
		ScalarId:      r.FormValue("ScalarId"),
		Vector:        r.FormValue("Vector"),
		Coordinate:    r.FormValue("Coordinate"),
		Value:         r.FormValue("Value"),
		Status:        r.FormValue("Status"),
		Symbol:        r.FormValue("Symbol"),
		Terminated:    r.FormValue("Terminated"),
		Decimals:      r.FormValue("Decimals"),
	}

	vh.service.CreateVegetable(vegetable)
	vh.IndexHandler(writer, r)
}

func (vh *VegetableHandler) DownloadHandler(writer http.ResponseWriter, request *http.Request) {

}
