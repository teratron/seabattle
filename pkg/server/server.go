package server

import (
	"github.com/teratron/seabattle/pkg/logger"
	"log"
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
	return &Server{
		Server: http.Server{
			Addr:     addr,
			ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		},
	}
}

// NewAndRun
func NewAndRun(addr string) error {
	return New(addr).Run()
}

// Run
func (s *Server) Run() error {
	s.Server.Handler = s
	logger.Info.Printf("Listening on port %s", s.Server.Addr)
	logger.Info.Printf("Open http://%s in the browser", s.Server.Addr)
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

// DownloadHandler
/*func DownloadHandler(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, filepath.Clean(name))
}*/

// HandlerFunc is a function type that implements the http.Handler interface.
/*type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
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
