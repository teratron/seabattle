package server

import (
	"net/http"
	"strconv"
	//"github.com/teratron/seabattle/pkg/config"
	//"github.com/teratron/seabattle/pkg/logger"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type Server struct {
	http.Server

	router
}

// New initializes a new Server.
func New(addr ...string) *Server {
	srv := &Server{
		router: router{
			Server:  config.Server{}.New(),
			Handler: config.Handler{}.New(),
			Logger:  logger.New(),
		},
	}
	if len(addr) > 0 {
		srv.Addr = addr[0]
	} else {
		if err := srv.LoadConfig(srv.router.Server.File); err != nil {
			srv.Addr = "localhost:8080"
		}
	}
	srv.Server.Handler = srv
	return srv
}

// LoadConfig
func (srv *Server) LoadConfig(path string) (err error) {
	if err = srv.Decode(path); err == nil {
		srv.Addr = srv.Host + ":" + strconv.Itoa(srv.Port)
	}
	return
}

// Run
func (srv *Server) Run() error {
	srv.Info.Printf("Listening on port %d", srv.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Addr)
	return srv.ListenAndServe()
}
