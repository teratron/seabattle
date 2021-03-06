package app

import (
	"path/filepath"

	"github.com/teratron/seabattle/pkg/util"
)

type Config struct {
	file string
	Err  error `json:"-" yaml:"-"`

	Application string `json:"application" yaml:"application"`
	Version     string `json:"version" yaml:"version"`
	Runtime     string `json:"runtime" yaml:"runtime"`
	ApiVersion  string `json:"api_version" yaml:"api_version"`

	*Settings `json:"settings" yaml:"settings"`
}

type Settings struct {
	Language string `json:"language" yaml:"language"`
	Theme    string `json:"theme" yaml:"theme"`
}

// NewApp
func NewConfig() *Config {
	cfg := &Config{
		file:        filepath.Join("configs", "app.yml"),
		Application: "seabattle",
		Version:     "0.0.1",
		Runtime:     "go116",
		ApiVersion:  "go1",
		Settings: &Settings{
			Language: "en",
			Theme:    "default",
		},
	}

	file := util.GetFileType(cfg.file)
	if err, ok := file.(*util.FileError); !ok {
		cfg.Err = cfg.Decode(file)
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *Config) Decode(decoder util.Decoder) error {
	return decoder.Decode(cfg)
}
