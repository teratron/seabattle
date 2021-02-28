package app

import (
	"sync"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/server"

	_ "github.com/go-sql-driver/mysql"
)

type Application struct {
	srv      *server.Server
	cfg      *config.ConfApp
	log      *logger.Logger
	settings *settings
	mu       sync.Mutex
}

// New
func New() *Application {
	return &Application{
		srv:      server.New(),
		cfg:      config.NewConfApp(),
		log:      logger.New(),
		settings: &settings{},
		mu:       sync.Mutex{},
	}
}

// Server
func (app *Application) Server() {
	//app.srv.Config = app.cfg
	app.srv.Logger = app.log
	app.srv.ErrorLog = app.log.Error

	app.srv.HandleEntry()
	app.srv.HandleFile("./web/static")
}

// Run
func (app *Application) Run() {
	//_ = app.cfg.Encode(filepath.Join("configs", "config2.yml"))
	app.log.Error.Fatal(app.srv.Start())
}

/*func (app *Application) handle() {
	app.srv.HandleFunc("/", handler.Home)
	app.srv.HandleFunc("/about", handler.About)
	app.srv.HandleFunc("/error", handler.Error)
	app.srv.HandleFile("./web/static")
}*/

// Theme
func (app *Application) Theme() string {
	return app.settings.theme
}
