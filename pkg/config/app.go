package config

import (
	"path/filepath"
)

type ConfApp struct {
	file string
	Err  error `yaml:"-"`

	Application string `yaml:"application"`
	Version     string `yaml:"version"`
	Runtime     string `yaml:"runtime"`
	ApiVersion  string `yaml:"api_version"`

	//ConfServer  `yaml:"server"`
	//ConfHandler `yaml:"handler"`
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
	//cfg.Err = cfg.Decode(cfg.file)
	return cfg
}
