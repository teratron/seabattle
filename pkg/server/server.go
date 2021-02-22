package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type Server struct {
	http.Server
	http.ServeMux
	http.FileSystem

	*config.Config
	*logger.Logger
}

// New initializes a new Server.
func New(addr ...string) *Server {
	srv := new(Server)
	if len(addr) > 0 {
		srv.Addr = addr[0]
	} else {
		srv.Addr = "localhost:8080"
	}
	srv.Server.Handler = srv
	return srv
}

// LoadConfig
func (srv *Server) LoadConfig(path string) (err error) {
	err = srv.Decode(path)
	if err == nil {
		srv.Addr = srv.Config.Server.Host + ":" + strconv.Itoa(srv.Config.Server.Port)
	}
	return
}

// Run
func (srv *Server) Run() error {
	srv.Info.Printf("Listening on port %d", srv.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Addr)
	return srv.ListenAndServe()
}

// HandleFile initializes http.FileServer, that will handle
// HTTP-requests to static files from a folder (for example: "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (for example: "/static/").
func (srv *Server) HandleFile(path string) {
	srv.FileSystem = http.Dir(path)
	pattern := "/" + filepath.Base(path)
	srv.Handle(pattern, http.NotFoundHandler())
	srv.Handle(pattern+"/", http.StripPrefix(pattern, http.FileServer(srv)))
}

// Open makes the Server implement the http.FileSystem interface.
// Check if the file is present index.html in static folders.
func (srv *Server) Open(path string) (file http.File, err error) {
	if file, err = srv.FileSystem.Open(path); err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			if _, err = srv.FileSystem.Open(path + "index.html"); err != nil {
				if file.Close() != nil {
					return nil, err
				}
			}
		}
	}
	return
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

/*type HandlerString string

func (h HandlerString) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//h(w, r)
	_, _ = fmt.Fprintf(w, "Welcome %s", h)
}*/

func (srv *Server) HandleEntry() {
	for key, value := range srv.Entry {
		fmt.Println(key, value)
		//srv.Handle(key, HandlerString(value))
	}
	/*var r HandlerFunc
	srv.Handle("/", r)
	srv.HandleFunc("/", r)
	srv.HandleFunc("/", handler.Home)
	srv.HandleFunc("/about", handler.About)
	srv.HandleFunc("/error", handler.Error)
	srv.HandleFile("./web/static")*/
}

// HandleMethod
func (srv *Server) HandleMethod(method string, pattern string, handle HandlerFunc) {
	switch method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodPatch:
	case http.MethodDelete:
	case http.MethodHead:
	case http.MethodOptions:
	default:
		srv.Error.Printf("wrong method: %s", method)
	}
}

// GET
func (srv *Server) GET(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodGet, pattern, handle)
}

// POST
func (srv *Server) POST(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodPost, pattern, handle)
}

// PUT
func (srv *Server) PUT(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodPut, pattern, handle)
}

// PATCH
func (srv *Server) PATCH(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodPatch, pattern, handle)
}

// DELETE
func (srv *Server) DELETE(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodDelete, pattern, handle)
}

// HEAD
func (srv *Server) HEAD(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodHead, pattern, handle)
}

// OPTIONS
func (srv *Server) OPTIONS(pattern string, handle HandlerFunc) {
	srv.HandleMethod(http.MethodOptions, pattern, handle)
}
