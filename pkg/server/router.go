package server

import (
	"net/http"
)

type router struct {
}

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

// HandleMethod
func (srv *Server) HandleMethod(method string, pattern string, handler http.Handler) {
	/*switch handle.(type) {
	case HandlerFunc:

	}*/
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
func (srv *Server) GET(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodGet, pattern, handler)
}

// POST
func (srv *Server) POST(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodPost, pattern, handler)
}

// PUT
func (srv *Server) PUT(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodPut, pattern, handler)
}

// PATCH
func (srv *Server) PATCH(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodPatch, pattern, handler)
}

// DELETE
func (srv *Server) DELETE(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodDelete, pattern, handler)
}

// HEAD
func (srv *Server) HEAD(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodHead, pattern, handler)
}

// OPTIONS
func (srv *Server) OPTIONS(pattern string, handler http.Handler) {
	srv.HandleMethod(http.MethodOptions, pattern, handler)
}

// DownloadHandler
/*func DownloadHandler(w http.ResponseWriter, r *http.Request, name string) {
	http.ServeFile(w, r, filepath.Clean(name))
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
