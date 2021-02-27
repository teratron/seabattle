package handler

/*import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/server"
)

func About(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		http.NotFound(w, r)
		return
	}
	page := &server.Page{
		Data: server.Data{
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
		Files: []string{
			filepath.Join(server.PathTemplateDir, "page.about.tmpl"),
			filepath.Join(server.PathTemplateDir, "partial.header.tmpl"),
			filepath.Join(server.PathTemplateDir, "partial.footer.tmpl"),
			filepath.Join(server.PathTemplateDir, "layout.base.tmpl"),
		},
	}
	tmpl, err := template.ParseFiles(page.Files...)
	if err == nil {
		err = tmpl.Execute(w, page.Data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}*/
