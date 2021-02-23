package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type router struct {
	http.ServeMux
	http.FileSystem

	*config.Config
	*logger.Logger
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

// GET
func (r *router) GET(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodGet, pattern, handler)
}

// POST
func (r *router) POST(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPost, pattern, handler)
}

// PUT
func (r *router) PUT(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPut, pattern, handler)
}

// PATCH
func (r *router) PATCH(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPatch, pattern, handler)
}

// DELETE
func (r *router) DELETE(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodDelete, pattern, handler)
}

// HEAD
func (r *router) HEAD(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodHead, pattern, handler)
}

// OPTIONS
func (r *router) OPTIONS(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodOptions, pattern, handler)
}

// HandleMethod
func (r *router) HandleMethod(method string, pattern string, handler http.Handler) {
	/*switch handle.(type) {
	case HandlerFunc:
	}*/
	switch method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	case http.MethodHead:
	case http.MethodOptions:
	default:
		r.Error.Printf("wrong method: %s", method)
	}
}

// HandleEntry
func (r *router) HandleEntry() {
	i := 0
	page := make([]*Page, len(r.Entry))
	for key, value := range r.Entry {
		page[i] = &Page{
			pattern: key,
			Page:    value,
		}
		r.Handle(key, page[i])
		i++
	}
}

// HandleFile initializes http.FileServer, that will handle
// HTTP-requests to static files from a folder (for example: "./web/static").
// Use the Handle() function to register a handler for all requests
// that start with the pattern  (for example: "/static/").
func (r *router) HandleFile(path string) {
	r.FileSystem = http.Dir(path)
	pattern := "/" + filepath.Base(path)
	r.Handle(pattern, http.NotFoundHandler())
	r.Handle(pattern+"/", http.StripPrefix(pattern, http.FileServer(r)))
}

// Open makes the Server implement the http.FileSystem interface.
// Check if the file is present index.html in static folders.
func (r *router) Open(path string) (file http.File, err error) {
	if file, err = r.FileSystem.Open(path); err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			if _, err = r.FileSystem.Open(path + "index.html"); err != nil {
				if file.Close() != nil {
					return nil, err
				}
			}
		}
	}
	return
}

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
