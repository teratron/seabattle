package server

import (
	"net/http"
	"strconv"
	//"github.com/teratron/seabattle/pkg/config"
	//"github.com/teratron/seabattle/pkg/logger"
)

type Server struct {
	http.Server
	router
	/*http.ServeMux
	http.FileSystem

	*config.Config
	*logger.Logger*/
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
	err = srv.Decode(path)
	if err == nil {
		srv.Addr = srv.Config.Server.Host + ":" + strconv.Itoa(srv.Config.Server.Port)
	}
	return
}

// Run
func (srv *Server) Run() error {
	srv.Info.Printf("Listening on port %d", srv.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Addr)
	return srv.ListenAndServe()
}
