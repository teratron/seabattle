package router

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// ServeMux is an HTTP request multiplexer
type ServeMux struct {
	http.ServeMux
}

// NewServeMux() инициализация нового рутера.
func NewServeMux() *ServeMux {
	return new(ServeMux)
}

// HandleFileServer инициализирует FileServer, он будет обрабатывать
// HTTP-запросы к статическим файлам из папки "./web/static".
// Используем функцию mux.Handle() для регистрации обработчика для
// всех запросов, которые начинаются с "/static/".
func (mux *ServeMux) HandleFileServer(path string) {
	p := "/" + filepath.Base(path)
	mux.Handle(p, http.NotFoundHandler())
	mux.Handle(p+"/", http.StripPrefix(p, http.FileServer(FileSystem{http.Dir(path)})))
}

// FileSystem is a customizable struct type that implements the http.FileSystem interface.
type FileSystem struct {
	fs http.FileSystem
}

func (fs FileSystem) Open(path string) (file http.File, err error) {
	fmt.Print(path, " - ")
	file, err = fs.fs.Open(path)
	if err == nil {
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			path = filepath.Join(path, "index.html")
			if _, err = fs.fs.Open(path); err != nil {

			}
			err = file.Close()
		}
	}
	fmt.Println(file, err)
	return
}

// DownloadHandler
func DownloadHandler(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, filepath.Clean(name))
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

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
