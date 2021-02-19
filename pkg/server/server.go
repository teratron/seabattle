package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

const (
	// DefaultAddr
	DefaultAddr = "localhost:8080"

	// DefaultHost
	DefaultHost = "localhost"

	// DefaultAddr
	DefaultPort = "8080"
)

type Server struct {
	http.Server
	http.ServeMux
	http.FileSystem

	*config.Config
	*logger.Logger
}

// New инициализация нового Server.
func New(addr ...string) *Server {
	srv := new(Server)
	srv.Config = new(config.Config)
	srv.Logger = new(logger.Logger)

	if len(addr) > 0 {
		srv.Addr = addr[0]
	} else {
		fmt.Println(srv.Config.File)
		if err := srv.LoadConfig(filepath.Join(".", "configs", "config.yml")); err != nil {
			srv.Addr = DefaultAddr
		}
		//srv.Addr = DefaultAddr
	}
	srv.Server.Handler = srv
	return srv
}

// LoadConfig
func (srv *Server) LoadConfig(path string) (err error) {
	err = srv.Decode(path)
	return
}

// Run
func (srv *Server) Run() error {
	srv.Info.Printf("Listening on port %s", srv.Config.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Config.Host+":"+srv.Config.Port)
	return srv.ListenAndServe()
}

// HandleFileServer initializes http.FileServer, that will handle
// HTTP-requests to static files from a folder (for example "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (например "/static/").
func (srv *Server) HandleFileServer(path string) {
	srv.FileSystem = http.Dir(path)
	pattern := "/" + filepath.Base(path)
	srv.Handle(pattern, http.NotFoundHandler())
	srv.Handle(pattern+"/", http.StripPrefix(pattern, http.FileServer(srv)))
}

// Open implements the Server to http.FileSystem interface.
// Проверяем присутсвует файл index.html в статических папках.
func (srv *Server) Open(path string) (file http.File, err error) {
	file, err = srv.FileSystem.Open(path)
	if err == nil {
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

// DividePortFromAddr
/*func DividePortFromAddr(addr string) (port string) {
	index := strings.Index(addr, ":")
	if index > -1 && index < len(addr)-1 {
		split := strings.Split(addr, ":")
		index = len(split)
		if index > 0 {
			port = split[index-1]
		}
	}
	return
}*/
