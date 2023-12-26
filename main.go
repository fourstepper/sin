package main

import (
	"embed"
	"log"
	"net/http"

	"html/template"
)

//go:embed views/*.html
var tmpl embed.FS

func main() {
	mux := http.NewServeMux()

	// https://lets-go.alexedwards.net/sample/02.08-serving-static-files.html#the-http-fileserver-handler
	staticDir := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static", staticDir))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		tmpl, err := template.ParseFS(tmpl, "views/index.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/about.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	mux.HandleFunc("/second-about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/second-about.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", mux)
}
