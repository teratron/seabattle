package app

import (
	"sync"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

type Application interface {
}

type App struct {
	srv *server.Server
	cfg *config.ConfApp
	log *logger.Logger
	mu  sync.Mutex
}

// New
func New() *App {
	app := &App{
		srv: server.New(),
		cfg: config.NewConfApp(),
		mu:  sync.Mutex{},
	}

	app.log = app.srv.Logger

	return app
}

// Server
func (app *App) Server() {
	app.srv.HandleEntry()
	app.srv.HandleFile("./web/static")
}

// Run
func (app *App) Run() {
	app.log.Error.Fatal(app.srv.Start())
}

// Theme
/*func (app *Application) Theme() string {
	return app.settings.theme
}*/
