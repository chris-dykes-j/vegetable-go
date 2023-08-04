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

// VegetableHandler provides methods to handle url requests, returning web pages or handling CRUD operations
// Christopher Dykes, 041013556
type VegetableHandler struct {
	service *s.VegetableService
}

// InitializeHandler initializes the handler given a VegetableService struct.
// Christopher Dykes, 041013556
func InitializeHandler(service *s.VegetableService) *VegetableHandler {
	return &VegetableHandler{service}
}

// IndexHandler provides the index page for the web application.
// Christopher Dykes, 041013556
func (vh *VegetableHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("presentation/views/index.gohtml"))
	err := tmpl.Execute(w, vh.service.ReadAllVegetables())
	if err != nil {
		log.Fatal(err)
	}
}

func (vh VegetableHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("presentation/views/index.gohtml"))
	geography := r.URL.Query().Get("geography")
	vegetable := r.URL.Query().Get("vegetable")
	storage := r.URL.Query().Get("storage")
	date := r.URL.Query().Get("date")
	err := tmpl.Execute(w, vh.service.SearchVegetables(geography, vegetable, storage, date))
	if err != nil {
		log.Fatal(err)
	}
}

// EditHandler provides the edit page to make changes to Vegetables
// Christopher Dykes, 041013556
func (vh *VegetableHandler) EditHandler(w http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/edit.gohtml"))
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id > len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = editTmpl.Execute(w, vh.service.ReadVegetableById(id))
	if err != nil {
		log.Fatal(err)
	}
}

// AddHandler provides the add vegetable page to well... add a Vegetable.
// Christopher Dykes, 041013556
func (vh VegetableHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	editTmpl := template.Must(template.ParseFiles("presentation/views/add.gohtml"))
	err := editTmpl.Execute(w, m.Vegetable{})
	if err != nil {
		log.Fatal(err)
	}
}

// DeleteHandler removes a Vegetable given a specified id from the url.
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
	if id < 0 || id > len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	vh.service.DeleteVegetableById(id)
}

// ReloadHandler resets changes made to the changes made to the in memory vegetable list, and reloads from the vegetable service.
// Christopher Dykes, 041013556
func (vh *VegetableHandler) ReloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid r method", http.StatusMethodNotAllowed)
		return
	}
	vh.service.ResetVegetables()
}

// UpdateHandler updates vegetable fields given an id. Redirects user to the index page.
// Christopher Dykes, 041013556
func (vh VegetableHandler) UpdateHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if id < 0 || id > len(vh.service.ReadAllVegetables()) {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Can't parse form", http.StatusInternalServerError)
		return
	}

	vegetable := m.Vegetable{
		Id:            id,
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

// CreateHandler creates a new vegetable and adds it to the in memory list. Redirects user to the index page.
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

// DownloadHandler creates the csv file and serves it to the client.
// Christopher Dykes, 041013556
func (vh *VegetableHandler) DownloadHandler(w http.ResponseWriter, r *http.Request) {
	vh.service.WriteAsCsv()
	w.Header().Set("Content-Disposition", "attachment; filename="+path.Base("./files/vegetables.csv"))
	w.Header().Set("Content-Type", "text/csv")
	http.ServeFile(w, r, "./files/vegetables.csv")
}
