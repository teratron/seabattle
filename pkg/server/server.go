package server

import (
	"net/http"
	"strconv"

	"github.com/teratron/seabattle/pkg/config"
)

type Server struct {
	http.Server
	*Router

	cfg *config.Server
}

// New initializes a new Server.
func New() *Server {
	srv := &Server{
		Router: NewRouter(),
		cfg:    config.NewServer(),
	}

	if srv.cfg.Err != nil {
		srv.log.Error.Printf("load default config: %v", srv.cfg.Err)
	}

	srv.Addr = srv.cfg.Host + ":" + strconv.Itoa(srv.cfg.Port)
	srv.ReadHeaderTimeout = srv.cfg.Header
	srv.ReadTimeout = srv.cfg.Read
	srv.WriteTimeout = srv.cfg.Write
	srv.IdleTimeout = srv.cfg.Idle
	srv.ErrorLog = srv.log.Error
	srv.Handler = srv

	return srv
}

// Start
func (srv *Server) Start() error {
	srv.log.Info.Print("Start server")
	srv.log.Info.Printf("Listening on port %d", srv.cfg.Port)
	srv.log.Info.Printf("Open http://%s in the browser", srv.Addr)

	err := srv.ListenAndServe()
	srv.log.Error.Print(err)

	return err
}

// Stop
func (srv *Server) Stop() {
	srv.log.Info.Print("Stop server")
}

// Restart
func (srv *Server) Restart() {
	srv.log.Info.Print("Restart server")
}

// Address
func (srv *Server) Address() string {
	return srv.Server.Addr
}

// SetAddress
func (srv *Server) SetAddress(addr string) {
	srv.Server.Addr = addr
}
