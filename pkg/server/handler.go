package server

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teratron/seabattle/pkg/config"
)

// Page
type Page struct {
	pattern string
	*config.Page
}

func (p *Page) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != p.pattern {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(p.Files...)
	if err == nil {
		err = tmpl.Execute(w, p.Data)
	}

	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
	}
}
