package main

import (
	"fmt"
	"log"
	"time"
)
import "net/http"

func main() {
	// We can create a custom server to override any of the options
	mux := http.NewServeMux()
	mux.HandleFunc("/", customIndexHandler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("HTTP server failed on 8080: %v", err)
	}
}

func customIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Index page")
}
