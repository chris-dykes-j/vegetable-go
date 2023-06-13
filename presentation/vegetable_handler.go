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

// Christopher Dykes, 041013556
type VegetableHandler struct {
	service *s.VegetableService
}

// Christopher Dykes, 041013556
func InitializeHandler(service *s.VegetableService) *VegetableHandler {
	return &VegetableHandler{service}
}

// Christopher Dykes, 041013556
func (vh *VegetableHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("presentation/views/index.gohtml"))
	err := tmpl.Execute(w, vh.service.ReadAllVegetables())
	if err != nil {
		log.Fatal(err)
	}
}

// Christopher Dykes, 041013556
func (vh *VegetableHandler) EditHandler(w http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/edit.gohtml"))
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = editTmpl.Execute(w, vh.service.ReadVegetableById(id))
	if err != nil {
		log.Fatal(err)
	}
}

// Christopher Dykes, 041013556
func (vh VegetableHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/add.gohtml"))
	err := editTmpl.Execute(w, m.Vegetable{})
	if err != nil {
		log.Fatal(err)
	}
}

// Christopher Dykes, 041013556
func (vh *VegetableHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid r method", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	vh.service.DeleteVegetableById(id)
}

// Christopher Dykes, 041013556
func (vh *VegetableHandler) ReloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid r method", http.StatusMethodNotAllowed)
		return
	}
	vh.service.ReloadVegetables()
}

// Christopher Dykes, 041013556
func (vh VegetableHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id >= len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusInternalServerError)
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
	http.Redirect(w, r, "/", http.StatusSeeOther) // 303
}

// Christopher Dykes, 041013556
func (vh VegetableHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusInternalServerError)
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
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (vh *VegetableHandler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vh.service.WriteAsCsv()
	w.Header().Set("Content-Disposition", "attachment; filename="+path.Base("./files/vegetables.csv"))
	w.Header().Set("Content-Type", "text/csv")
	http.ServeFile(w, r, "./files/vegetables.csv")
}
