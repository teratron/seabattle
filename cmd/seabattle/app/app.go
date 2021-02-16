package app

import (
	"github.com/teratron/seabattle/pkg/config"
	"github.com/teratron/seabattle/pkg/logger"
	"os/exec"
	"sync"
)

type Application struct {
	cfg  config.Config
	log  *logger.Logger
	lang string

	settings *settings
	running  bool
	mutex    sync.Mutex
	exec     func(name string, arg ...string) *exec.Cmd
}

func (app *Application) Theme() string {
	return app.settings
}
