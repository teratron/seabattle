package config

import (
	"path/filepath"

	"github.com/teratron/seabattle/pkg/api"
)

type ConfApp struct {
	file string
	Err  error `yaml:"-"`

	Application string `yaml:"application"`
	Version     string `yaml:"version"`
	Runtime     string `yaml:"runtime"`
	ApiVersion  string `yaml:"api_version"`
}

// NewConfApp
func NewConfApp() *ConfApp {
	cfg := &ConfApp{
		file:        filepath.Join("configs", "app.yml"),
		Application: "seabattle",
		Version:     "0.0.1",
		Runtime:     "go116",
		ApiVersion:  "go1",
	}

	file := api.GetFileType(cfg.file)
	if err, ok := file.(*api.FileError); !ok {
		cfg.Err = cfg.Decode(file)
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *ConfApp) Decode(decoder api.Decoder) error {
	return decoder.Decode(cfg)
}
