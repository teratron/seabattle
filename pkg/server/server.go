package server

import (
	"github.com/teratron/seabattle/pkg/logger"
	"log"
	"os"

	"net/http"
)

type Server struct {
	http.Server
}

// New
func New(addr string, handler ...http.Handler) *Server {
	/*for range handler {
		fmt.Println("*")
	}*/
	//mux := http.ServeMux{}
	srv := &Server{
		Server: http.Server{
			Addr:    addr,
			Handler: handler[0],
			//Handler: &http.ServeMux{},
			//Handler: &mux,
			//ErrorLog: logger.Error,
			ErrorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		},
	}
	return srv
}

// Run
func (s *Server) Run() (err error) {
	logger.Info.Printf("Listening on port %s", s.Addr)
	logger.Info.Printf("Open http://%s in the browser", s.Addr)
	return s.ListenAndServe()
}

// NewAndRun
func NewAndRun(addr string, handler ...http.Handler) error {
	return New(addr).Run()
}
