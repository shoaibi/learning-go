package main

import (
	"learning-go/16_net/homepage"
	"log"
	"net/http"
	"os"
)


// Example from: https://www.youtube.com/watch?v=bM6N-vgPlyQ

func main()  {
	logger := log.New(os.Stdout, "[app] ", log.LstdFlags | log.Lshortfile)
	mux := http.NewServeMux()

	h := homepage.New(logger)
	h.SetupRoutes(mux)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		logger.Fatalf("Could not start HTTP Server on :8080 : %v", err)
	}
}

