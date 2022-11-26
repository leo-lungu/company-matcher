package main

import (
	"encoding/json"
	"fmt"
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
	Name string `json:"name"`
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

func getData() map[string]interface{} {
	readFile, err := ioutil.ReadFile("scraped.json")
	// println(readFile)
	handleErr(err)

	var dat map[string]interface{}

	_ = json.Unmarshal([]byte(readFile), &dat)
	// fmt.Println(dat)
	// for x := range dat {
	// 	fmt.Println(dat[x].(map[string]interface{})["Name"])
	// }
	handleErr(err)

	// for key := range item {
	// 	fmt.Println(item[key].Name)
	// }

	return dat
}

func main() {

	items := getData()

	itemss := make([]Company, 0, 1)

	for x := range items {
		itemss = append(itemss, Company{
			Name: items[x].(map[string]interface{})["Name"].(string),
		})
	}

	for i := range itemss {
		fmt.Println(itemss[i].Name)
	}

	// http.HandleFunc("/", indexHandler)
	// log.Fatal(http.ListenAndServe(":8080", nil))

	amp()
}
