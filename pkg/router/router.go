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
