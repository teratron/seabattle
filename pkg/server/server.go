package server

import (
	"net/http"
	"strconv"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
)

type Server struct {
	http.Server
	router
}

// New initializes a new Server.
func New() *Server {
	srv := &Server{
		router: router{
			ConfServer:  config.NewConfServer(),
			ConfHandler: config.NewConfHandler(),
			Logger:      logger.New(),
		},
	}

	if srv.ConfServer.Err != nil {
		srv.Warning.Printf("load default config: %v", srv.ConfServer.Err)
	}

	srv.Addr = srv.Host + ":" + strconv.Itoa(srv.Port)
	srv.Handler = srv
	srv.ReadHeaderTimeout = srv.ConfServer.Header
	srv.ReadTimeout = srv.ConfServer.Read
	srv.WriteTimeout = srv.ConfServer.Write
	srv.IdleTimeout = srv.ConfServer.Idle

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
