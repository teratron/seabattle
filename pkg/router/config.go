package router

import (
	"html/template"
	"path/filepath"
	"time"

	"github.com/teratron/seabattle/pkg/util"
)

type Config struct {
	file string
	Err  error `json:"-" yaml:"-"`

	// Server
	Host     string `json:"host" yaml:"host"`
	Port     int    `json:"port" yaml:"port"`
	*Timeout `json:"timeout" yaml:"timeout"`

	// Handlers
	*Commons `json:"commons" yaml:"commons"`
	Entry    map[string]*Page `json:"entry" yaml:"entry"`
}

type Timeout struct {
	Header time.Duration `json:"header" yaml:"header"`
	Read   time.Duration `json:"read" yaml:"read"`
	Write  time.Duration `json:"write" yaml:"write"`
	Idle   time.Duration `json:"idle" yaml:"idle"`
}

type Commons struct {
	Lang  string                  `json:"lang" yaml:"lang"`
	Theme string                  `json:"theme" yaml:"theme"`
	Meta  map[string]string       `json:"meta" yaml:"meta"`
	Path  map[string]template.URL `json:"path" yaml:"path"` // List of static path (img, css, js & etc)
}

type Page struct {
	*Data   `json:"data" yaml:"data"`
	Files   []string `json:"files" yaml:"files"`
	pattern string
}

type Data struct {
	*Commons `json:"-" yaml:"-"`
	//params   Parameter

	Title    string            `json:"title" yaml:"title"`
	AttrHTML map[string]string `json:"attrHTML" yaml:"attrHTML"` // List of attributes attached to the <html> tag
	AttrBody map[string]string `json:"attrBody" yaml:"attrBody"` // List of attributes attached to the <body> tag
}

// NewConfig
func NewConfig() *Config {
	cfg := &Config{
		file: filepath.Join("configs", "router.yml"),
		Host: "localhost",
		Port: 8080,
		/*Timeout: &Timeout{
			Header: 30 * time.Second,
			Read:   15 * time.Second,
			Write:  10 * time.Second,
			Idle:   5 * time.Second,
		},
		Commons: &Commons{
			Lang:  "en",
			Theme: "light",
			Meta: map[string]string{
				"robots": "index, follow",
			},
			Path: map[string]template.URL{
				"img": "../static/img/",
				"css": "../static/css/",
				"js":  "../static/js/",
			},
		},*/
		//Entry: make(map[string]*Page),
	}

	file := util.GetFileType(cfg.file)
	if err, ok := file.(*util.FileError); !ok {
		cfg.Err = cfg.Decode(file)
		if cfg.Err == nil {
			for key, value := range cfg.Entry {
				if value != nil {
					value.pattern = key
					value.Commons = cfg.Commons
					for i, file := range value.Files {
						value.Files[i] = filepath.Join("web", "template", file)
					}
				}
			}
		}
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *Config) Decode(data interface{}) error {
	return data.(util.Decoder).Decode(cfg)
}
