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
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(tmpl, "views/index.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/about.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.HandleFunc("/second-about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("views/second-about.html", "views/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", r)
}
