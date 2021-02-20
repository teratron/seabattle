package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
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
	//fmt.Println(srv.Config.File)
	err = srv.Decode(path)
	if err == nil {
		srv.Addr = srv.Config.Server.Host + ":" + strconv.Itoa(srv.Config.Server.Port)
	}
	return
}

func (srv *Server) handle() {
	for _, v := range srv.Entry {
		fmt.Println(v)
	}
	srv.HandleFunc("/", handler.Home)
	srv.HandleFunc("/about", handler.About)
	srv.HandleFunc("/error", handler.Error)
	srv.HandleFileServer("./web/static")
}

// Run
func (srv *Server) Run() error {
	srv.Info.Printf("Listening on port %d", srv.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Addr)

	//fmt.Println(srv.Info)
	//srv.Warning.F()

	return srv.ListenAndServe()
}

// HandleFileServer initializes http.FileServer, that will handle
// HTTP-requests to static files from a folder (for example: "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (for example: "/static/").
func (srv *Server) HandleFileServer(path string) {
	srv.FileSystem = http.Dir(path)
	pattern := "/" + filepath.Base(path)
	srv.Handle(pattern, http.NotFoundHandler())
	srv.Handle(pattern+"/", http.StripPrefix(pattern, http.FileServer(srv)))
}

// Open implements the Server to http.FileSystem interface.
// Проверяем присутсвует файл index.html в статических папках.
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
