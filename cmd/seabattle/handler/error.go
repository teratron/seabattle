package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
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
	layout := &Layout{
		data: Data{
			Name:  "error",
			Title: "Error",
		},
		files: []string{
			filepath.Join(PathDirTemplate, "page.error.tmpl"),
			filepath.Join(PathDirTemplate, "partial.header.tmpl"),
			filepath.Join(PathDirTemplate, "partial.footer.tmpl"),
			filepath.Join(PathDirTemplate, "layout.base.tmpl"),
		},
	}
	tmpl, err := template.ParseFiles(layout.files...)
	if err == nil {
		err = tmpl.Execute(w, layout.data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
