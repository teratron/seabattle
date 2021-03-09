package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/teratron/seabattle/pkg/logger"
)

type Router struct {
	http.Server
	http.ServeMux
	http.FileSystem

	*Config
	*logger.Logger
}

// New
func New() *Router {
	r := &Router{
		Config: NewConfig(),
		Logger: logger.New(),
	}
	r.Logger.File = filepath.Join("logs", "server.log")

	if r.Config.Err != nil {
		r.Logger.Error.Printf("load default config: %v", r.Config.Err)
	}

	r.Server.Addr = r.Config.Host + ":" + strconv.Itoa(r.Config.Port)
	r.Server.ReadHeaderTimeout = r.Config.Header
	r.Server.ReadTimeout = r.Config.Read
	r.Server.WriteTimeout = r.Config.Write
	r.Server.IdleTimeout = r.Config.Idle
	r.Server.ErrorLog = r.Logger.Error
	r.Server.Handler = r

	return r
}

// Start
func (r *Router) Start() error {
	r.Logger.Info.Print("Start server")
	r.Logger.Info.Printf("Listening on port %d", r.Config.Port)
	r.Logger.Info.Printf("Open http://%s in the browser", r.Server.Addr)

	err := r.ListenAndServe()
	r.Logger.Error.Print(err)

	return err
}

// Stop
func (r *Router) Stop() {
	r.Logger.Info.Print("Stop server")
}

// Address
func (r *Router) Address() string {
	return r.Server.Addr
}

// SetAddress
func (r *Router) SetAddress(addr string) {
	r.Server.Addr = addr
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

// HandleEntry
func (r *Router) HandleEntry() {
	for key, value := range r.Entry {
		r.Handle(key, &page{key, value})
	}
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
		r.Error.Printf("wrong method: %s", method)
	}
}
