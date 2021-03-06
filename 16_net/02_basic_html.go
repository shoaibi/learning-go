package main

import (
	"fmt"
	"log"
)
import "net/http"

func main() {
	http.HandleFunc("/", plainHandler)
	http.HandleFunc("/html", htmlHandler)
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatalf("HTTP server failed on 8020: %v", err)
	}
}

func plainHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Plain data")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "<h1>html data</h1>")
	fmt.Fprint(w, "<h2>Testing tags</h2>")
	fmt.Fprint(w, "<h3>H3 Tag</h3>")
	// Better to use multiline version
	fmt.Fprint(w, `<h4>Hello</h4>
<p> Testing multiline here
<p>..... it is cool </p>
	`)
}
