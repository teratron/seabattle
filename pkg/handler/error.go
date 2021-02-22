package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/server"
)

func Error(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Используем метод Header().Set() для добавления заголовка 'Allow: POST' в
		// карту HTTP-заголовков. Первый параметр - название заголовка, а
		// второй параметр - значение заголовка.
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	layout := &server.Layout{
		Data: server.Data{
			Name:  "error",
			Title: "Error",
		},
		Files: []string{
			filepath.Join(server.PathTemplateDir, "page.error.tmpl"),
			filepath.Join(server.PathTemplateDir, "partial.header.tmpl"),
			filepath.Join(server.PathTemplateDir, "partial.footer.tmpl"),
			filepath.Join(server.PathTemplateDir, "layout.base.tmpl"),
		},
	}
	tmpl, err := template.ParseFiles(layout.Files...)
	if err == nil {
		err = tmpl.Execute(w, layout.Data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
