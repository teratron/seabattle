package server

import (
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	http.Server
	http.ServeMux
	fs http.FileSystem
}

// New инициализация нового Server.
func New(addr string) *Server {
	srv := &Server{
		Server: http.Server{
			Addr: addr,
		},
	}
	srv.Server.Handler = srv
	return srv
}

// Run
func (srv *Server) Run() error {
	return srv.Server.ListenAndServe()
}

func (srv *Server) Address() string {
	return srv.Addr
}

// HandleFileServer инициализирует http.FileServer, который будет обрабатывать
// HTTP-запросы к статическим файлам из папки (например "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (например "/static/").
func (srv *Server) HandleFileServer(path string) {
	srv.fs = http.Dir(path)
	p := "/" + filepath.Base(path)
	srv.Handle(p, http.NotFoundHandler())
	srv.Handle(p+"/", http.StripPrefix(p, http.FileServer(srv)))
}

// Open implements the Server to http.FileSystem interface.
// Проверяем присутсвует файл index.html в статических папках.
func (srv *Server) Open(path string) (file http.File, err error) {
	file, err = srv.fs.Open(path)
	if err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			if _, err = srv.fs.Open(path + "index.html"); err != nil {
				if file.Close() != nil {
					return nil, err
				}
			}
		}
	}
	return
}
