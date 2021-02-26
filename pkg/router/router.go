package router

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type Router struct {
	http.Server
	http.ServeMux
	http.FileSystem

	*config.ConfRouter
	*config.ConfHandler
	*logger.Logger
}

// New initializes a new Router.
func New() *Router {
	r := &Router{
		ConfRouter:  config.NewConfRouter(),
		ConfHandler: config.NewConfHandler(),
		Logger:      logger.New(),
	}

	if r.ConfRouter.Err != nil {
		r.Error.Printf("load default config: %v", r.ConfRouter.Err)
	}

	r.Server.Addr = r.Host + ":" + strconv.Itoa(r.Port)
	r.Server.Handler = r
	r.Server.ReadHeaderTimeout = r.ConfRouter.Header
	r.Server.ReadTimeout = r.ConfRouter.Read
	r.Server.WriteTimeout = r.ConfRouter.Write
	r.Server.IdleTimeout = r.ConfRouter.Idle
	r.Server.ErrorLog = r.Logger.Error

	return r
}

// LoadConfig
func (r *Router) LoadConfig(path string) (err error) {
	// TODO:
	return
}

// Run
func (r *Router) Run() error {
	r.Info.Printf("Listening on port %d", r.Port)
	r.Info.Printf("Open http://%s in the browser", r.Addr)

	return r.ListenAndServe()
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
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

// HandleEntry
func (r *Router) HandleEntry() {
	for key, value := range r.Entry {
		r.Handle(key, &Page{key, value})
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
