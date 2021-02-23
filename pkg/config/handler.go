package config

import (
	"path/filepath"
)

type Handler struct {
	File string `yaml:"-"`

	Base `yaml:"base,flow"`

	Entry map[string]Page `yaml:"entry"`
}

type Base struct {
	Lang        string            `yaml:"lang"`
	Description string            `yaml:"description"`
	Author      string            `yaml:"author"`
	Keyword     string            `yaml:"keyword"`
	Theme       string            `yaml:"theme"`
	Path        map[string]string `yaml:"path"` // List of static path
}

type Page struct {
	Data `yaml:"data"`

	Files []string `yaml:"files,omitempty"`
}

type Data struct {
	base *Base

	Name     string            `yaml:"name"`
	Title    string            `yaml:"title"`
	AttrHTML map[string]string `yaml:"attrHTML"` // List of attributes attached to the <html> tag
	AttrBody map[string]string `yaml:"attrBody"` // List of attributes attached to the <body> tag
}

// New
func (h Handler) New() *Handler {
	return &Handler{
		File: filepath.Join("configs", "handler.yml"),
	}
}
