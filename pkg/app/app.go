package app

import (
	"sync"

	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/router"

	_ "github.com/go-sql-driver/mysql"
)

type Application interface {
}

type App struct {
	mu  sync.Mutex
	srv *router.Router
	log *logger.Logger
	cfg *Config
}

// New
func New() *App {
	app := &App{
		mu:  sync.Mutex{},
		srv: router.New(),
		log: logger.New(),
		cfg: NewConfig(),
	}
	//app.log = app.srv.log

	return app
}

// Server
func (app *App) Server() {
	app.srv.HandleEntry()
	app.srv.HandleStatic("./web/static")
}

// Run
func (app *App) Run() {
	app.log.Error.Fatal(app.srv.Start())
}

// Theme
/*func (app *Application) Theme() string {
	return app.settings.theme
}*/
