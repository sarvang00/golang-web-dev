package controllers

import (
	"net/http"

	"github.com/sarvang00/golang-web-dev/lenslocked/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, nil)
	}
}
