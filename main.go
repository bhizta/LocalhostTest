package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Test web server</h1>")
    fmt.Fprintf(w, "<p>1st Paragraph</p>")
    fmt.Fprintf(w, "<p>I'm a beginner Go developer.</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>About Me</h1>")
    fmt.Fprintf(w, "<p>This is the About page.</p>")
}

func main() {
    http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
    fmt.Println("Server started at http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}