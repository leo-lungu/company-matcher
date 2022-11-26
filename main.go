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

	//for i := range companies {
	//	fmt.Println(companies[i])
	//}

	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

	amp()
}
