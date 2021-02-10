package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := &Layout{
		Name:   "Sea Battle",
		Lang:   "en",
		Title:  "Sea Battle",
		Author: "Oleg Alexandrov",
	}
	files := []string{
		"./template/home.page.tmpl",
		"./template/header.partial.tmpl",
		"./template/footer.partial.tmpl",
		"./template/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
	/*tmpl, err := template.ParseFiles("template/base.layout.tmpl",
		"template/home.page.tmpl",
		"template/header.partial.tmpl",
		"template/footer.partial.tmpl")
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "base", data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}*/
}
