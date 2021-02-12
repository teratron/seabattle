package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	layout := &Layout{
		data: Data{
			Name:        "about",
			Lang:        "en",
			Description: "Sea Battle - multiplayer online game",
			Author:      "Oleg Alexandrov",
			Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
			Title:       "About Us",
			Theme:       "default",

			AttrHTML: map[string]string{
				"class": "",
			},
			AttrBody: map[string]string{
				"id":    "about",
				"class": "about",
			},
		},
		files: []string{
			filepath.Join(PathDirTemplate, "page.about.tmpl"),
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
