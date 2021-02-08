package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	//_ "github.com/teratron/seabattle/pkg/router"
)

type page struct {
	Name   string
	Lang   string
	Title  string
	Author string
}

func init() {
	err := os.Setenv("PORT", "8080")
	if err != nil {
		log.Println(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	home := &page{
		Name:   "Sea Battle",
		Lang:   "en",
		Title:  "Sea Battle",
		Author: "Oleg Alexandrov",
	}
	tmpl, err := template.ParseFiles("template/page.html",
		"template/index.html",
		"template/header.html",
		"template/footer.html")
	if err == nil {
		err = tmpl.ExecuteTemplate(w, "page", home)
	}
	if err != nil {
		//w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
	}

}

func main() {
	fmt.Println("Sea Battle")

	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
