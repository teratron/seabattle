package app

import (
	"log"
	"sync"

	"github.com/teratron/seabattle/cmd/seabattle/handler"
	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/server"
)

type Application struct {
	srv      *server.Server
	cfg      *config.Config
	log      *logger.Logger
	settings *settings
	mu       sync.Mutex
}

// New
func New() *Application {
	app := &Application{
		srv:      server.New(),
		cfg:      config.New(),
		log:      logger.New(),
		settings: &settings{},
		mu:       sync.Mutex{},
	}
	app.srv.Config = app.cfg
	app.srv.Logger = app.log
	app.srv.ErrorLog = app.log.Error
	return app
}

// Server
func (app *Application) Server() {
	//app.srv.Creat()
	app.handle()
	app.log.Info.Printf("Listening on port %s", app.cfg.Server.Port)
	app.log.Info.Printf("Open http://%s in the browser", app.cfg.Server.Addr)
	log.Fatal(app.srv.Run())
}

func (app *Application) handle() {
	app.srv.HandleFunc("/", handler.Home)
	app.srv.HandleFunc("/about", handler.About)
	app.srv.HandleFunc("/error", handler.Error)
	app.srv.HandleFileServer("./web/static")
}

// Theme
func (app *Application) Theme() string {
	return app.settings.theme
}
