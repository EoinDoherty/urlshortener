package urlshortener

import (
	"fmt"
)

var urls map[string]string = make(map[string]string)
var counter int

// LookupURL Look up the long URL (if it exists) from the database
func LookupURL(id string) (string, error) {

	longURL, found := urls[id]

	if !found {
		return "", fmt.Errorf("LookupURL: URL not found")
	}

	return longURL, nil
}

// AddURL Add the long URL to the database
func AddURL(url string) (string, error) {

	newID := fmt.Sprintf("%d", counter)
	urls[newID] = url
	counter++

	return newID, nil
}
