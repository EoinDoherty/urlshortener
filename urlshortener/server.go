package urlshortener

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// InitServer Initialize URL shortener server
func InitServer() {
	r := mux.NewRouter()

	r.HandleFunc("/s/{id}", redirectShortened)
	r.HandleFunc("/shorten", shortenURL)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./urlshortener/static/")))

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
