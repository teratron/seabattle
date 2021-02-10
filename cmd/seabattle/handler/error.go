package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Error(w http.ResponseWriter, r *http.Request) {
	data := &Layout{
		Title: "Error",
	}
	files := []string{
		"./template/page.error.tmpl",
		"./template/layout.base.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err == nil {
		err = tmpl.Execute(w, data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
