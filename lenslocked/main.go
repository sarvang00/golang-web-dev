package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Unix-only:
	// tpl, err := template.ParseFiles("templates/home.gohtml")
	// Universal:
	tplPath := filepath.Join("templates", "home.gohtml")
	renderTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	renderTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func renderTemplate(w http.ResponseWriter, templatePath string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		log.Println("Error parsing template: ", err)
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to execute template: %v. Maybe check the data passed.\n", err)
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		return
	}
}

func paramHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "hi %v", userID)
}

func main() {
	rtr := chi.NewRouter()
	rtr.Use(middleware.Logger)
	rtr.Get("/", homeHandler)
	rtr.Get("/contact", contactHandler)
	rtr.Get("/faq", faqHandler)

	rtr.Route("/params/", func(subRtr chi.Router) {
		subRtr.Get("/{userID}", paramHandler)
	})

	rtr.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", rtr)
}
