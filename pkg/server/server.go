package server

import (
	"net/http"
	"strconv"

	"github.com/teratron/seabattle/pkg/config"
)

type Server struct {
	http.Server
	*Router
	*config.ConfServer
}

// New initializes a new Server.
func New() *Server {
	srv := &Server{
		Router:     NewRouter(),
		ConfServer: config.NewConfServer(),
	}

	if srv.ConfServer.Err != nil {
		srv.Error.Printf("load default config: %v", srv.ConfServer.Err)
	}
	if srv.ConfHandler.Err != nil {
		srv.Error.Printf("load default config: %v", srv.ConfHandler.Err)
	}

	srv.Server.Addr = srv.Host + ":" + strconv.Itoa(srv.Port)
	srv.Server.Handler = srv
	srv.Server.ReadHeaderTimeout = srv.ConfServer.Header
	srv.Server.ReadTimeout = srv.ConfServer.Read
	srv.Server.WriteTimeout = srv.ConfServer.Write
	srv.Server.IdleTimeout = srv.ConfServer.Idle
	srv.Server.ErrorLog = srv.Logger.Error

	return srv
}

// Start
func (srv *Server) Start() error {
	srv.Info.Print("Start server")
	srv.Info.Printf("Listening on port %d", srv.Port)
	srv.Info.Printf("Open http://%s in the browser", srv.Addr)

	err := srv.ListenAndServe()
	srv.Error.Print(err)

	return err
}

// Stop
func (srv *Server) Stop() {
	srv.Info.Print("Stop server")
}

// Restart
func (srv *Server) Restart() {
	srv.Info.Print("Restart server")
}

// Address
func (srv *Server) Address() string {
	return srv.Server.Addr
}

// SetAddress
func (srv *Server) SetAddress(addr string) {
	srv.Server.Addr = addr
}
