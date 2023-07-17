package main

import (
	"log"
	"net/http"
	s "practical/business"
	h "practical/presentation"
)

// main is the entry point for the application. Contains the service, handler, manages routes.
// Christopher Dykes, 041013556
func main() {
	vegetableService := s.InitializeService()
	handler := h.InitializeHandler(vegetableService)

	http.HandleFunc("/", handler.IndexHandler)
	http.HandleFunc("/edit/", handler.EditHandler)
	http.HandleFunc("/add/", handler.AddHandler)
	http.HandleFunc("/delete/", handler.DeleteHandler)
	http.HandleFunc("/reload/", handler.ReloadHandler)
	http.HandleFunc("/update/", handler.UpdateHandler)
	http.HandleFunc("/create/", handler.CreateHandler)
	http.HandleFunc("/download/", handler.DownloadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
