package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}
	data := &Layout{
		Name:        "Sea Battle",
		Lang:        "en",
		Description: "Sea Battle - multiplayer online game",
		Author:      "Oleg Alexandrov",
		Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
		Title:       "About",

		files: []string{
			"./template/page.about.tmpl",
			"./template/partial.header.tmpl",
			"./template/partial.footer.tmpl",
			"./template/layout.base.tmpl",
		},
	}
	/*files := []string{
		"./template/page.about.tmpl",
		"./template/partial.header.tmpl",
		"./template/partial.footer.tmpl",
		"./template/layout.base.tmpl",
	}*/
	tmpl, err := template.ParseFiles(data.files...)
	if err == nil {
		err = tmpl.Execute(w, data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}
