package main

import (
	"fmt"
	"log"
	"time"
)
import "net/http"

// Most examples are from: https://www.alexedwards.net/blog/a-recap-of-request-handling
// Really good blog, love the detailed and easy explanations
// Buy Alex a beer if you find these useful.

type CounterHandler struct {
	counter int
}

func (ct *CounterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ct.counter++
	fmt.Fprint(w, "Counter:", ct.counter)
}

type timeHandler struct {
	format string
}

func (th *timeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tm := time.Now().Format(th.format)
	// We can directly write a bytes array to the responseWriter too
	w.Write([]byte("The time is: " + tm))
}

func main() {
	http.HandleFunc("/", funcHandler)

	// There is also Handle on http and ServeMux which takes a type that has Handler interface implemented
	ch := &CounterHandler{counter: 0}
	http.Handle("/count", ch)

	// Another one to tell us the time
	th1123 := &timeHandler{format: time.RFC1123}
	http.Handle("/time/rfc1123", th1123)

	// Reusing timeHandle to produce different formats
	th3339 := &timeHandler{format: time.RFC3339}
	http.Handle("/time/rfc3339", th3339)

	// We can avoid creating a whole struct for one value
	// and use a function that returns Handler interface
	// complying function instead
	// This could be useful in real world where we might want
	// to pass around database connection string, template map
	// or anything else to avoid global variables
	thf822Z := timeHandlerFunc(time.RFC822Z)
	http.Handle("/time/rfc822z", thf822Z)

	// We can also wrap our function as HandleFunc and still use Handle() but it isn't as useful
	http.Handle("/function", http.HandlerFunc(funcHandler))

	// http comes with other useful handlers such as FileServer, NotFoundHandler, RedirectHandler

	log.Println("Starting the default one on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP server failed on 8080: %v", err)
	}
}

func funcHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "I am default one")
}

func timeHandlerFunc(format string) http.Handler {
	// The trick here is to create a closure that closes over
	//  the parameters to the outer function and then return
	//  a converted handler function from it
	fn := func(w http.ResponseWriter, r *http.Request) {
		tm := time.Now().Format(format)
		w.Write([]byte("The time is: " + tm))
	}
	return http.HandlerFunc(fn)

	// We could make it bit more generic by doing:
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	tm := time.Now().Format(format)
	//	w.Write([]byte("The time is: " + tm))
	//})

	// or even an implicit conversion to the HandlerFunc type on return:
	// func timeHandler(format string) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		tm := time.Now().Format(format)
	//		w.Write([]byte("The time is: " + tm))
	//	}

}
