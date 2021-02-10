package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := &Layout{
		Name:        "Sea Battle",
		Lang:        "en",
		Description: "Sea Battle - multiplayer online game",
		Author:      "Oleg Alexandrov",
		Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
		Title:       "Sea Battle - Home",
	}
	files := []string{
		"./template/page.home.tmpl",
		"./template/partial.header.tmpl",
		"./template/partial.footer.tmpl",
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
