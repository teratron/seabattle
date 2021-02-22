package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/server"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	layout := &server.Layout{
		Data: server.Data{
			Lang:        "en",
			Description: "Sea Battle - multiplayer online game",
			Author:      "Oleg Alexandrov",
			Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
			Theme:       "dark",

			Name:  "Home",
			Title: "Sea Battle - Home",

			AttrHTML: map[string]string{
				"class": "",
			},
			AttrBody: map[string]string{
				"id":    "home",
				"class": "home",
			},
			Path: map[string]string{
				"img": "../static/img/",
				"css": "../static/css/",
				"js":  "../static/js/",
			},
		},
		Files: []string{
			filepath.Join(server.PathTemplateDir, "page.home.tmpl"),
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
