package main

import (
	"fmt"
	"log"
)
import "net/http"

func main() {
	// By default go has a defaultServeMux so we don't need to define our own.
	// This is also the 2nd param to ListenAndServer by default.
	// ServeMux can be thought of a special implementation of Handler interface
	//	which doesn't handle the request on its own but hands it over to
	//	the most suitable handle; think chaining.
	// Sometimes we want to initialize 2 different http servers, listening
	// on different ports. This is where the mux concept because quite useful.
	// The http.HandleFunc() actually proxies to mux.HandleFunc

	// Using the default ServeMux is discouraged as it is stored as global and
	// is accessible to any package in the application; poses a security risk

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "I am custom one")
		})
		log.Println("Starting the default one on :8282")
		if err := http.ListenAndServe(":8282", mux); err != nil {
			log.Fatalf("HTTP server failed on 8282: %v", err)
		}
	}()

	http.HandleFunc("/", defaultHandler)
	log.Println("Starting the default one on :8080")
	// this will use the default ServeMux
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server failed on 8080: %v", err)
	}
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am default one")
}
