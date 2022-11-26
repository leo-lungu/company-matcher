package main

import (
	"log"
	"net/http"
	"os"
	"text/template"
)

type Company struct {
	Name  string `json:"Name"`
	Value map[string]string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, r)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
}

func main() {
	readFile, err := os.ReadFile("scraped.json")
	if err != nil {
		log.Fatal(err)
	}

	strBody := string(readFile)

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
