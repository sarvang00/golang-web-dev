package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Hello from this site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page!</h1>")
	fmt.Fprint(w, "<p>To get in touch, Email: <a href=\"mailto:sdave@dal.ca\">sdave@dal.ca</a></p>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintln(w, r.URL.Path)
// 	// fmt.Fprintln(w, r.URL.RawPath)
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		// Page not found error
// 		// http.NotFound(w, r)
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 		// w.WriteHeader(http.StatusNotFound)
// 		// fmt.Fprint(w, "Page not found")
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func main() {
	// http.HandleFunc("/", pathHandler)
	// http.HandleFunc("/contact", contactHandler)
	var customRouter Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", customRouter)
}
