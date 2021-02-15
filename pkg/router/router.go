package router

import (
	"github.com/teratron/seabattle/pkg/logger"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Router is an HTTP request multiplexer
type Router struct {
	http.Server
	http.ServeMux
	fs http.FileSystem
}

// New инициализация нового рутера.
func New() *Router {
	return new(Router)
}

func NewS(addr string) *Router {
	return &Router{
		Server: http.Server{
			Addr:     addr,
			ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		},
	}
}

func (r *Router) Run() error {
	r.Server.Handler = r
	logger.Info.Printf("Listening on port %s", r.Server.Addr)
	logger.Info.Printf("Open http://%s in the browser", r.Server.Addr)
	return r.Server.ListenAndServe()
}

// HandleFileServer инициализирует FileServer, который будет обрабатывать
// HTTP-запросы к статическим файлам из папки (например "./web/static").
// Используем функцию Handle() для регистрации обработчика для
// всех запросов, которые начинаются с паттерна (например "/static/").
func (r *Router) HandleFileServer(path string) {
	r.fs = http.Dir(path)
	p := "/" + filepath.Base(path)
	r.Handle(p, http.NotFoundHandler())
	r.Handle(p+"/", http.StripPrefix(p, http.FileServer(r)))
}

// Open implements the Router to http.FileSystem interface.
// Проверяем присутсвует файл index.html в статических папках.
func (r *Router) Open(path string) (file http.File, err error) {
	file, err = r.fs.Open(path)
	if err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			if _, err = r.fs.Open(path + "index.html"); err != nil {
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
