package router

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

// Router is an HTTP request multiplexer
type Router struct {
	http.ServeMux
	fs http.FileSystem
}

// NewRouter() инициализация нового рутера.
func NewRouter() *Router {
	return new(Router)
}

// HandleFileServer инициализирует FileServer, он будет обрабатывать
// HTTP-запросы к статическим файлам из папки "./web/static".
// Используем функцию mux.Handle() для регистрации обработчика для
// всех запросов, которые начинаются с "/static/".
func (r *Router) HandleFileServer(path string) {
	r.fs = http.Dir(path)
	p := "/" + filepath.Base(path)
	r.Handle(p, http.NotFoundHandler())
	r.Handle(p+"/", http.StripPrefix(p, http.FileServer(r)))
}

// Проверяем присутсвует файл index.html
// Open implements the Router to http.FileSystem interface.
func (r *Router) Open(path string) (file http.File, err error) {
	fmt.Println(path, " - ")
	file, err = r.fs.Open(path)
	if err == nil {
		fmt.Println("++++ 1", file, err)
		var info os.FileInfo
		if info, err = file.Stat(); err == nil && info.IsDir() {
			fmt.Println("++++ 2", file, err)
			path = path + "index.html" //filepath.Join(path, "index.html")
			if _, err = r.fs.Open(path); err != nil {
				fmt.Println("++++ 3", file, err)
				return file, file.Close()
			}
		}
	}
	fmt.Println("++++ 4", file, err)
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
