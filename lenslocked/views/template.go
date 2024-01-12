package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Must(tpl Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return tpl
}

func Parse(templatePath string) (Template, error) {
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return Template{}, fmt.Errorf("Template parsing error %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl, err := template.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("FS Template parsing error %w", err)
	}
	return Template{htmlTpl: tpl}, nil
}

func (tpl Template) ExecuteTemplate(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tpl.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("Failed to execute template: %v. Maybe check the data passed.\n", err)
		http.Error(w, "Template execution failed", http.StatusInternalServerError)
		return
	}
}
