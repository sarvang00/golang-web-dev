package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/sarvang00/golang-web-dev/lenslocked/controllers"
	"github.com/sarvang00/golang-web-dev/lenslocked/templates"
	"github.com/sarvang00/golang-web-dev/lenslocked/views"
)

func paramHandler(w http.ResponseWriter, r *http.Request) {
	userID := chi.URLParam(r, "userID")
	fmt.Fprintf(w, "hi %v", userID)
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

func main() {
	rtr := chi.NewRouter()

	rtr.Use(middleware.Logger)

	rtr.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))
	rtr.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml"))))
	rtr.Get("/faq", controllers.FAQ(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))
	rtr.Route("/params/", func(subRtr chi.Router) {
		subRtr.Get("/{userID}", paramHandler)
	})

	rtr.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", rtr)
}
