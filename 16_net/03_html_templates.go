package main

import (
	"html/template"
	"log"
)
import "net/http"

type Post struct {
	Title          string
	CreatedAt      string
	LastModifiedAt string
}

type About struct {
	Name  string
	Email string
	Posts []Post
}

func main() {
	http.HandleFunc("/", postsHandler)
	if err := http.ListenAndServe(":8020", nil); err != nil {
		log.Fatalf("HTTP server failed on 8020: %v", err)
	}
}

func postsHandler(w http.ResponseWriter, r *http.Request) {
	p := []Post{
		{
			Title:          "Hello World",
			CreatedAt:      "2006-01-01",
			LastModifiedAt: "2006-01-03",
		},
		{
			Title:          "Second Post",
			CreatedAt:      "2006-01-05",
			LastModifiedAt: "2006-01-06",
		},
		{
			Title:          "Third Post",
			CreatedAt:      "2006-01-08",
			LastModifiedAt: "2006-01-10",
		},
		{
			Title:          "Fourth Post",
			CreatedAt:      "2006-01-22",
			LastModifiedAt: "2006-03-21",
		},
	}
	d := About{
		Name:  "Shoaib",
		Email: "shoaibi@mailinator.com",
		Posts: p,
	}

	t, _ := template.ParseFiles("03_posts.html")
	// This line returns an error if there are placeholders that aren't translated
	t.Execute(w, d)
}
