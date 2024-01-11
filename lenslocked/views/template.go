package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}

func Parse(templatePath string) (Template, error) {
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return Template{}, fmt.Errorf("Template parsing error %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

func ParseFS(fs fs.FS, pattern string) (Template, error) {
	tpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("FS Template parsing error %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

type Template struct {
	htmlTpl *template.Template
}

func (t Template) ExecuteTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("Failed to execute template: %v. Maybe check the data passed.\n", err)
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		return
	}
}
