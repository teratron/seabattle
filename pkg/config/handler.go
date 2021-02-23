package config

type Handler struct {
	Common `yaml:"common"`

	Entry map[string]Page `yaml:"entry"`
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
	Data `yaml:"data"`

	Files []string `yaml:"files,omitempty"`
}

type Data struct {
	common *Common

	Name     string            `yaml:"name"`
	Title    string            `yaml:"title"`
	AttrHTML map[string]string `yaml:"attrHTML"` // List of attributes attached to the <html> tag
	AttrBody map[string]string `yaml:"attrBody"` // List of attributes attached to the <body> tag
}
