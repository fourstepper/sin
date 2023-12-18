package main

import (
	"log"
	"net/http"

	"html/template"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/index.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/about.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	r.HandleFunc("/second-about", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("view/second-about.html", "view/base.html")
		if err != nil {
			log.Fatal(err)
		}
		tmpl.ExecuteTemplate(w, "base", "data")
	})
	log.Println("Listening...")
	http.ListenAndServe(":3000", r)
}
