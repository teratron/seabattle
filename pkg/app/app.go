package app

import (
	//_ "database/sql"
	"fmt"
	"path/filepath"
	"sync"

	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/handler"
	"github.com/teratron/seabattle/pkg/logger"
	"github.com/teratron/seabattle/pkg/server"

	_ "github.com/go-sql-driver/mysql"
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
	return &Application{
		srv:      server.New(),
		cfg:      config.New(),
		log:      logger.New(),
		settings: &settings{},
		mu:       sync.Mutex{},
	}
}

// Server
func (app *Application) Server() {
	app.srv.Config = app.cfg
	app.srv.Logger = app.log
	app.srv.ErrorLog = app.log.Error

	//app.log.Warning =
	_ = app.srv.LoadConfig(filepath.Join("configs", "config.yml"))

	app.handle()
}

// Run
func (app *Application) Run() {
	fmt.Println(app.cfg.Encode(filepath.Join("configs", "config2.yml")))
	app.log.Error.Fatal(app.srv.Run())
}

func (app *Application) Handle(pattern string, handle server.HandlerFunc) {
	//TODO:
}

func (app *Application) handle() {
	/*for _, v := range app.cfg.Entry {
		//fmt.Println(v)
		app.srv.HandleFunc(v, handler.Home)
	}*/
	app.srv.HandleFunc("/", handler.Home)
	app.srv.HandleFunc("/about", handler.About)
	app.srv.HandleFunc("/error", handler.Error)
	app.srv.HandleFile("./web/static")
}

// Theme
func (app *Application) Theme() string {
	return app.settings.theme
}
