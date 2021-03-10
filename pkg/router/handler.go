package router

import (
	"fmt"
	"net/http"
)

// HandlerFunc is a function type that implements the http.Handler interface.
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func (p Pattern) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != p["pattern"] {
		http.NotFound(w, r)
		return
	}
	//var f = p["files"].([]interface{})[0]
	//fmt.Println(f[0])
	//fmt.Printf("%T %v", p["files"].([]interface{})[0], p["files"].([]interface{})[0])
	//fmt.Printf("%T %v", p["data"].(map[interface{}]interface{})["title"], p["data"].(map[interface{}]interface{})["title"])
	if files, ok := p["files"].([]interface{}); ok {
		fmt.Println(files)
	}

	/*funcMap := template.FuncMap{
		"attrMap": func(m map[string]string) template.HTMLAttr {
			var s string
			for k, v := range m {
				if len(v) > 0 {
					s += fmt.Sprintf(" %s=%s", k, template.HTMLEscapeString(v))
				}
			}
			return (template.HTMLAttr)(s)
		},
		"attr": func(s string) template.HTMLAttr {
			return (template.HTMLAttr)(s)
		},
		"safe": func(s string) template.HTML {
			return (template.HTML)(s)
		},
		"url": func(s string) template.URL {
			return (template.URL)(s)
		},
		"css": func(s string) template.CSS {
			return (template.CSS)(s)
		},
		"js": func(s string) template.JS {
			return (template.JS)(s)
		},
	}

	err := template.Must(template.New(filepath.Base(p["files"].([]interface{})[0].(string))).Funcs(funcMap).ParseFiles(p["files"]...)).Execute(w, p["data"].(map[interface{}]interface{}))
	if err != nil {
		_, err = fmt.Fprintf(w, err.Error())
	}*/
}
