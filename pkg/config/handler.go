package config

import (
	"path/filepath"

	"github.com/teratron/seabattle/pkg/api"
)

type ConfHandler struct {
	file string
	Err  error `json:"-" yaml:"-"`

	Common `json:"common" yaml:"common"`
	Entry  map[string]*Page `json:"entry" yaml:"entry"`
}

type Common struct {
	Lang  string            `json:"lang" yaml:"lang"`
	Theme string            `json:"theme" yaml:"theme"`
	Meta  map[string]string `json:"meta" yaml:"meta"`
	Path  map[string]string `json:"path" yaml:"path"` // List of static path
}

type Page struct {
	Data  `json:"data" yaml:"data"`
	Files []string `json:"files,omitempty" yaml:"files,omitempty"`
}

type Data struct {
	*Common `json:"-" yaml:"-"`

	Name     string            `json:"name" yaml:"name"`
	Title    string            `json:"title" yaml:"title"`
	AttrHTML map[string]string `json:"attrHTML" yaml:"attrHTML"` // List of attributes attached to the <html> tag
	AttrBody map[string]string `json:"attrBody" yaml:"attrBody"` // List of attributes attached to the <body> tag
}

// NewConfHandler
func NewConfHandler() *ConfHandler {
	cfg := &ConfHandler{
		file: filepath.Join("configs", "handler.yml"),
		Common: Common{
			Lang:  "en",
			Theme: "default",
			Meta: map[string]string{
				"description": "",
				"author":      "",
				"keyword":     "",
				"generator":   "",
				"robots":      "index, follow",
			},
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
