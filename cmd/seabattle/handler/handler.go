package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
)

var (
	PathDirWeb      = filepath.Join("web")
	PathDirStatic   = filepath.Join(PathDirWeb, "static")
	PathDirTemplate = filepath.Join(PathDirWeb, "template")
)

// Layout
type Layout struct {
	data  Data
	files []string
}

// Data
type Data struct {
	Name string

	// Header
	Lang        string
	Description string
	Author      string
	Keyword     string
	Title       string

	// The web theme name
	Theme string

	// List of attributes attached to the <html> tag
	AttrHTML map[string]string

	// List of attributes attached to the <body> tag
	AttrBody map[string]string

	// List of static path
	Path map[string]string
}

// Path static
/*type Path struct {
	Img, CSS, JS string
}*/

// Theme
type Theme struct {
	name string
}

func (l *Layout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, l.data, r)
}

// Index
func Index() *Layout {
	return &Layout{
		data: Data{
			Name:        "Sea Battle",
			Lang:        "en",
			Description: "Sea Battle - multiplayer online game",
			Author:      "Oleg Alexandrov",
			Keyword:     "SeaBattle,Sea,Battle,Multiplayer,Online,Game",
			Title:       "Sea Battle - Home",
			Theme:       "default",

			AttrHTML: map[string]string{
				"class": "home",
			},
			AttrBody: map[string]string{
				"id":    "home",
				"class": "home",
			},
		},
		files: []string{
			filepath.Join(PathDirTemplate, "page.home.tmpl"),
			filepath.Join(PathDirTemplate, "partial.header.tmpl"),
			filepath.Join(PathDirTemplate, "partial.footer.tmpl"),
			filepath.Join(PathDirTemplate, "layout.base.tmpl"),
		},
	}
}

/*func init() {
	mux := http.NewServeMux()

	h1 := Layout{Name: "Index"}
	h2 := Layout{Name: "About"}

	mux.Handle("/123", &h1)
	mux.Handle("/1234", &h2)

	_ = http.ListenAndServe(":8282", mux)
}*/
