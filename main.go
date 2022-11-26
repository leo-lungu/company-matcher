package main

import (
	"log"
	"net/http"
	"text/template"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, r)
	amp()
}

func main() {
	log.Println("Listening on :5500")
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":5500", nil))
}
