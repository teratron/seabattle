package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		//http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	layout := &Layout{
		data: Data{
			Name:        "home",
			Lang:        "en",
			Description: "Sea Battle - multiplayer online game",
			Author:      "Oleg Alexandrov",
			Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
			Title:       "Sea Battle - Home",
			Theme:       "dark",

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
			/*Path: Path{
				Img: "../static/img/",
				CSS: "../static/css/",
				JS:  "../static/js/",
			},*/
		},
		files: []string{
			filepath.Join(PathDirTemplate, "page.home.tmpl"),
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
