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

	funcMap := template.FuncMap{
		"attrMap": func(m map[string]string) template.HTMLAttr {
			var s string
			for k, v := range m {
				s += fmt.Sprintf(` %s="%s"`, k, v)
			}
			return template.HTMLAttr(s)
		},
		"attr": func(s string) template.HTMLAttr { return template.HTMLAttr(s) },
		"safe": func(s string) template.HTML { return template.HTML(s) },
		"url":  func(s string) template.URL { return template.URL(s) },
		"css":  func(s string) template.CSS { return template.CSS(s) },
		"js":   func(s string) template.JS { return template.JS(s) },
	}
	//Sea := [10]int{2, 3, 5, 6, 8, 9}

	err := template.Must(template.New(p.Name).Funcs(funcMap).ParseFiles(p.Files...)).Execute(w, p.Data)
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
	}
}
