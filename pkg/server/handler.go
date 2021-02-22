package server

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
	Data  Data
	Files []string
}

// Data
type Data struct {
	Lang        string
	Description string
	Author      string
	Keyword     string
	Theme       string

	Name  string
	Title string

	// List of attributes attached to the <html> tag
	AttrHTML map[string]string

	// List of attributes attached to the <body> tag
	AttrBody map[string]string

	// List of static path
	Path map[string]string
}

/*func (h *HandlerFunc) ServHTTP(w http.ResponseWriter, r *http.Request, layout *Layout) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles(layout.Files...)
	if err == nil {
		err = tmpl.Execute(w, layout.Data)
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}
}*/

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

// DownloadHandler
/*func DownloadHandler(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, filepath.Clean(name))
}*/

/*var (
	NotFoundHandler       = StatusHandler(http.StatusNotFound)
	NotLegalHandler       = StatusHandler(451)
	NotImplementedHandler = StatusHandler(501)
)

// HandlerStatus is a function type that implements the Handler interface.
type StatusHandler int

//
func (s StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	code := int(s)
	w.WriteHeader(code)
	if _, err := io.WriteString(w, http.StatusText(code)); err != nil {
		log.Println(err)
	}
}*/

// ClientFunc is a function type that implements the Client interface.
//type ClientFunc func(*http.Request) (*http.Response, error)

// Do does the request
/*func (c ClientFunc) Do(r *http.Request) (*http.Response, error) {
	return c(r)
}*/
