package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type Company struct {
	Name          string
	Values        map[string]interface{}
	CompanySize   float64
	RetentionRate float64
}

type Results struct {
	Email string
}

const (
	integrity = iota
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, r)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("results.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, r)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getData() (data map[string]interface{}) {
	readFile, err := ioutil.ReadFile("scraped.json")
	handleErr(err)

	err = json.Unmarshal([]byte(readFile), &data)
	handleErr(err)

	return data
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	email := r.FormValue("email")

	d := Results{
		email,
	}

	tmpl, err := template.ParseFiles("result.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, d)
}

func main() {

	data := getData()

	companies := []Company{}

	for x := range data {
		companies = append(companies, Company{
			Name:          data[x].(map[string]interface{})["Name"].(string),
			Values:        data[x].(map[string]interface{})["Values"].(map[string]interface{}),
			CompanySize:   data[x].(map[string]interface{})["CompanySize"].(float64),
			RetentionRate: data[x].(map[string]interface{})["RetentionRate"].(float64),
		})
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", resultHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
	// amp()
}
