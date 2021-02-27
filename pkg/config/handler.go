package config

import (
	"path/filepath"

	"github.com/teratron/seabattle/pkg/api"
)

type ConfHandler struct {
	file string
	Err  error `yaml:"-"`

	Common `yaml:"common"`
	Entry  map[string]*Page `yaml:"entry"`
}

type Common struct {
	Lang        string            `yaml:"lang"`
	Description string            `yaml:"description"`
	Author      string            `yaml:"author"`
	Keyword     string            `yaml:"keyword"`
	Theme       string            `yaml:"theme"`
	Path        map[string]string `yaml:"path"` // List of static path
}

type Page struct {
	Data  `yaml:"data"`
	Files []string `yaml:"files,omitempty"`
}

type Data struct {
	*Common `yaml:"-"`

	Name     string            `yaml:"name"`
	Title    string            `yaml:"title"`
	AttrHTML map[string]string `yaml:"attrHTML"` // List of attributes attached to the <html> tag
	AttrBody map[string]string `yaml:"attrBody"` // List of attributes attached to the <body> tag
}

// NewConfHandler
func NewConfHandler() *ConfHandler {
	cfg := &ConfHandler{
		file: filepath.Join("configs", "handler.yml"),
		Common: Common{
			Lang:  "en",
			Theme: "default",
			Path: map[string]string{
				"img": "../static/img/",
				"css": "../static/css/",
				"js":  "../static/js/",
			},
		},
		Entry: make(map[string]*Page),
	}

	file := api.GetFileType(cfg.file)
	if err, ok := file.(*api.FileError); !ok {
		cfg.Err = cfg.Decode(file)
		if cfg.Err == nil {
			for _, value := range cfg.Entry {
				value.Common = &cfg.Common
				for i, file := range value.Files {
					value.Files[i] = filepath.Join("web", "template", file)
				}
			}
		}
	} else {
		cfg.Err = err.Err
	}
	return cfg
}

func (cfg *ConfHandler) Decode(decoder api.Decoder) error {
	return decoder.Decode(cfg)
}
