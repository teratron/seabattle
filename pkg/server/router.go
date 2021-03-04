package server

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type Router struct {
	http.ServeMux
	http.FileSystem

	cfg *config.Handler
	log *logger.Logger
}

// NewRouter
func NewRouter() *Router {
	r := &Router{
		cfg: config.NewHandler(),
		log: logger.New(),
	}

	r.log.File = filepath.Join("logs", "server.log")

	if r.cfg.Err != nil {
		r.log.Error.Printf("load default config: %v", r.cfg.Err)
	}
	return r
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

// HandleEntry
func (r *Router) HandleEntry() {
	for key, value := range r.cfg.Entry {
		r.Handle(key, &Page{key, value})
	}
}

// HandlePage
func (r *Router) HandlePage(pattern string) {
	if value, exist := r.cfg.Entry[pattern]; exist {
		r.Handle(pattern, &Page{pattern, value})
	}
}

// HandleFile initializes http.FileServer, that will handle
// HTTP-requests to static files from a folder (for example: "./web/static").
// Use the Handle() function to register a handler for all requests
// that start with the pattern  (for example: "/static/").
func (r *Router) HandleFile(path string) {
	r.FileSystem = http.Dir(path)
	pattern := "/" + filepath.Base(path)
	r.Handle(pattern, http.NotFoundHandler())
	r.Handle(pattern+"/", http.StripPrefix(pattern, http.FileServer(r)))
}

// Open makes the Server implement the http.FileSystem interface.
// Check if the file is present index.html in static folders.
func (r *Router) Open(path string) (file http.File, err error) {
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

// GET
func (r *Router) GET(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodGet, pattern, handler)
}

// HEAD
func (r *Router) HEAD(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodHead, pattern, handler)
}

// POST
func (r *Router) POST(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPost, pattern, handler)
}

// PUT
func (r *Router) PUT(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPut, pattern, handler)
}

// PATCH
func (r *Router) PATCH(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodPatch, pattern, handler)
}

// DELETE
func (r *Router) DELETE(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodDelete, pattern, handler)
}

// CONNECT
func (r *Router) CONNECT(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodConnect, pattern, handler)
}

// OPTIONS
func (r *Router) OPTIONS(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodOptions, pattern, handler)
}

// TRACE
func (r *Router) TRACE(pattern string, handler http.Handler) {
	r.HandleMethod(http.MethodTrace, pattern, handler)
}

// HandleMethod
func (r *Router) HandleMethod(method string, pattern string, handler http.Handler) {
	switch method {
	case http.MethodGet:
		r.Handle(pattern, handler)
	case http.MethodHead:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	case http.MethodConnect:
	case http.MethodOptions:
	case http.MethodTrace:
	default:
		r.log.Error.Printf("wrong method: %s", method)
	}
}
