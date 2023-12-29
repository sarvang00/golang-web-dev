package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
<ul>
  <li>
    <b>Is there a free version?</b>
    Yes! We offer a free trial for 30 days on any paid plans.
  </li>
  <li>
    <b>What are your support hours?</b>
    We have support staff answering emails 24/7, though response
    times may be a bit slower on weekends.
  </li>
  <li>
    <b>How do I contact support?</b>
    Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>
  </li>
</ul>
`)
}

// Exercise 1 - URL Params
func paramHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// the above working shows just that it is perfectly fine working with the native http package
	// It was defaulting to text/plain type so setting it to content/html just for sakeof it
	fmt.Fprintf(w, "hi %v", userID)
	// We will get output "hi 1010"
}

func main() {
	rtr := chi.NewRouter()
	// Exercise 2.1 - Setting logger middleware to our application
	// r.Use(middleware.Logger)
	rtr.Get("/", homeHandler)
	rtr.Get("/contact", contactHandler)
	rtr.Get("/faq", faqHandler)
	// Exercise 1 : Using URL Params
	rtr.Route("/params/", func(subRtr chi.Router) {
		// Exercise 2.2 - Adding logger middleware to a single route
		subRtr.Use(middleware.Logger)
		// Assuming path should look like "/params/1010"
		subRtr.Get("/{userID}", paramHandler)
	})
	rtr.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", rtr)
}
