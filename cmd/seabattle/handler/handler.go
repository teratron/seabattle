package handler

import (
	"fmt"
	"net/http"
)

type Layout struct {
	Name        string
	Lang        string
	Description string
	Author      string
	Keyword     string
	Title       string

	files []string
}

func (l *Layout) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(resp, l.Name)
}

/*func init() {
	mux := http.NewServeMux()

	h1 := Layout{Name: "Index"}
	h2 := Layout{Name: "About"}

	mux.Handle("/123", &h1)
	mux.Handle("/1234", &h2)

	_ = http.ListenAndServe(":8282", mux)
}*/
