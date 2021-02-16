package server

import (
	"github.com/teratron/seabattle/pkg/logger"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	http.Server
	http.ServeMux
	fs http.FileSystem

	log *logger.Logger
}

// New инициализация нового Server.
func New(addr string) *Server {
	s := &Server{
		Server: http.Server{
			Addr: addr,
			//ErrorLog: ,
		},
		log: logger.New(),
	}
	s.Server.Handler = s
	s.Server.ErrorLog = s.log.Error
	return s
}

// NewAndRun
func NewAndRun(addr string) error {
	return New(addr).Run()
}

// Run
func (s *Server) Run() error {
	//s.Server.Handler = s
	s.log.Info.Printf("Listening on port %s", s.Server.Addr)
	s.log.Info.Printf("Open http://%s in the browser", s.Server.Addr)
	return s.Server.ListenAndServe()
}

// HandleFileServer инициализирует http.FileServer, который будет обрабатывать
// HTTP-запросы к статическим файлам из папки (например "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (например "/static/").
func (s *Server) HandleFileServer(path string) {
	s.fs = http.Dir(path)
	p := "/" + filepath.Base(path)
	s.Handle(p, http.NotFoundHandler())
	s.Handle(p+"/", http.StripPrefix(p, http.FileServer(s)))
}

// Open implements the Server to http.FileSystem interface.
// Проверяем присутсвует файл index.html в статических папках.
func (s *Server) Open(path string) (file http.File, err error) {
	file, err = s.fs.Open(path)
	if err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			if _, err = s.fs.Open(path + "index.html"); err != nil {
				if file.Close() != nil {
					return nil, err
				}
			}
		}
	}
	return
}
