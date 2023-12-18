package main

import (
	"log"
	"net/http"

	"html/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	log.Println("Started")
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/index.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.Get("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/about.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.Get("/second-about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/second-about.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	http.ListenAndServe(":3000", r)
}
