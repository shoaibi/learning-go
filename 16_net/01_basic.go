package main

import (
	"fmt"
	"log"
)
import "net/http"

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server failed on 8080: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// Generally recommended to write the StatusCode and
	// content-type headers to save the language from adding
	// it which could be a costly operation on huge responses
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Index page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "About page")
}
