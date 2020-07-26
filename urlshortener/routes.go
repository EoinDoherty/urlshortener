package urlshortener

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func redirectShortened(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	newURL, err := LookupURL(id)

	if err != nil {
		log.Printf("redirectShortened: %v", err)
		http.NotFound(w, r)
	} else {
		http.Redirect(w, r, newURL, 307)
	}
}

// shortenURL get URL from body, shorten it, and reply
func shortenURL(w http.ResponseWriter, r *http.Request) {
	resp, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Printf("shortenURL: %v", err)
	}

	defer r.Body.Close()

	statusCode := 200
	var message string

	if len(resp) == 0 {
		log.Printf("shortenURL: url empty")
		message = "This URL empty; yeet."
	} else {
		url := string(resp)

		message, err = AddURL(url)

		if err != nil {
			log.Printf("shortenedURL: %v", err)
			message = "error adding url"
		}
	}

	w.WriteHeader(statusCode)
	w.Write([]byte(message))
}
