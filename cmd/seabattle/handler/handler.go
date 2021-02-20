package handler

import (
	"path/filepath"
)

var (
	PathWebDir      = filepath.Join(".", "web")
	PathStaticDir   = filepath.Join(PathWebDir, "static")
	PathTemplateDir = filepath.Join(PathWebDir, "template")
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

/*func (l *Layout) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, l.data, r)
}
*/
/*func init() {
	mux := http.NewServeMux()

	h1 := Layout{Name: "Index"}
	h2 := Layout{Name: "About"}

	mux.Handle("/123", &h1)
	mux.Handle("/1234", &h2)

	_ = http.ListenAndServe(":8282", mux)
}*/
