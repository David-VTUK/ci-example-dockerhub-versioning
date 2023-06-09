package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		if err != nil {
			return
		}

		// Add an update
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Hi")
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
