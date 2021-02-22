package server

import (
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

// HandleEntry
func (srv *Server) HandleEntry() {
	i := 0
	page := make([]*Page, len(srv.Entry))
	for key, value := range srv.Entry {
		page[i] = &Page{
			pattern: key,
			Page:    value,
		}
		srv.Handle(key, page[i])
		i++
	}
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
