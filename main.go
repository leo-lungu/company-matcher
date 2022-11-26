package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

// type Company struct {
// 	Name          string
// 	Values        map[string]string
// 	CompanySize   int
// 	RetentionRate int
// }

type Company struct {
	Name string
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
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func getData() (data map[string]interface{}) {
	readFile, err := ioutil.ReadFile("scraped.json")
	handleErr(err)

	err = json.Unmarshal(readFile, &data)
	handleErr(err)

	return data
}

func main() {

	data := getData()

	// var companyData []Company

	for key, _ := range data {
		// fmt.Println(data[key])
		for _, value := range val.(map[string]interface{}) {

			// fmt.Println(data[key])
			// 	companyData = append(companyData, Company{
			// 		Name: value["Name"].(string),
			// 	})
		}
	}

	// for i := 0; i < len(strBody); i++ {
	// companyData = append(companyData, Company{
	// 	Name: string(strBody[i].Name),
	// 	Value: string(strBody[i].Value),
	// })
	// }

	// fmt.Println(strBody)

	// http.HandleFunc("/", indexHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	amp()
}
