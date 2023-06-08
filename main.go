package main

import (
	"log"
	"net/http"
	s "practical/business"
	h "practical/presentation"
)

func main() {
	vegetables := s.InitializeVegetables()
	handler := h.InitializeHandler(vegetables)

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/edit/", handler.EditHandler)
	http.HandleFunc("/delete/", handler.DeleteHandler)
	http.HandleFunc("/reload/", handler.ReloadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*

package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	s "practical/business"
	"strconv"
)

func main() {
	vegetables := s.InitializeVegetables()
	tmpl := template.Must(template.ParseFiles("presentation/views/index.gohtml"))
	indexHandler := func(writer http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(writer, vegetables.ReadAllVegetables())
		if err != nil {
			log.Fatal(err)
		}
	}

	editTmpl := template.Must(template.ParseFiles("presentation/views/edit.gohtml"))
	editHandler := func(writer http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			http.Error(writer, "Invalid ID", http.StatusBadRequest)
			return
		}
		if id < 0 || id >= len(vegetables.ReadAllVegetables()) {
			http.Error(writer, "Invalid ID", http.StatusBadRequest)
			return
		}

		err = editTmpl.Execute(writer, vegetables.ReadVegetableById(id))
		if err != nil {
			log.Fatal(err)
		}
	}

	deleteHandler := func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		id, err := strconv.Atoi(path.Base(r.URL.Path))
		if err != nil {
			http.Error(writer, "Invalid ID", http.StatusBadRequest)
			return
		}
		if id < 0 || id >= len(vegetables.ReadAllVegetables()) {
			http.Error(writer, "Invalid ID", http.StatusBadRequest)
			return
		}
		vegetables.DeleteVegetableById(id)
	}

	reloadHandler := func(writer http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(writer, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		vegetables.ReloadVegetables()
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/reload/", reloadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

*/
