package urlshortener

// InitServer Initialize URL shortener server
func InitServer() {
	r := mux.NewRouter()

	r.HandleFunc("/s/{id}", redirectShortened)
	r.HandleFunc("/shorten", shortenURL)

	http.Handle("/", r)

	log.Printf("Starting server")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
