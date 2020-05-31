package main

import "fmt"
import "net/http"

func main() {
	http.HandleFunc("/", plainHandler)
	http.HandleFunc("/html", htmlHandler)
	http.ListenAndServe(":8020", nil)
}

func plainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	fmt.Fprint(w, "Plain data")
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
	fmt.Fprint(w, "<h1>html data</h1>")
	fmt.Fprint(w, "<h2>Testing tags</h2>")
	fmt.Fprint(w, "<h3>H3 Tag</h3>")
	// Better to use multiline version
	fmt.Fprint(w, `<h4>Hello</h4>
<p> Testing multiline here
<p>..... it is cool </p>
	`)
}
