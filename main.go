package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type Company struct {
	Name          string
	Values        map[string]interface{}
	CompanySize   float64
	RetentionRate float64
}

type UserInputs struct {
	Name                string
	Email               string
	Motivations         []string
	IdeasOrExpand       string
	BigOrSmall          string
	JobHopOrStay        string
	MostImportantValues []string
	Location            string
	JobTitle            string
}

type ReturnData struct {
	Name                string
	Email               string
	Motivations         []string
	IdeasOrExpand       string
	BigOrSmall          string
	JobHopOrStay        string
	MostImportantValues []string
	Location            string
	JobTitle            string
	CompanyTest         Company
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

func getBestSuitedCompany(companies []Company, userInputs UserInputs) (bestCompany Company) {
	bestCompany = companies[0]
	bestScore := 0

	for _, company := range companies {
		score := 0

		for _, value := range userInputs.Motivations {
			if company.Values[value] != nil {
				continue
				// log.Println("Success: ", value, " ", company.Name, "Score: ", score)
			}
		}

		if userInputs.IdeasOrExpand == "Expand" {
			if company.Values["Innovative"] == true {
				score++
				// log.Println("innovative", company.Name, company.Values["Innovative"])
				// log.Println("score", score)
			}
		}

		if userInputs.BigOrSmall == "Big" {
			if company.CompanySize > 88000 {
				// log.Println("big", company.Name, company.CompanySize)
				// log.Println("score", score)
				score++
			}
		} else if userInputs.BigOrSmall == "Small" {
			if company.CompanySize < 88000 {
				// log.Println("small", company.Name, company.CompanySize)
				// log.Println("score", score)
				score++
			}
		}

		if userInputs.JobHopOrStay == "JobHop" {
			if company.RetentionRate < 50 {
				// log.Println("jobhop", company.Name, company.RetentionRate)
				// log.Println("score", score)
				score++
			}
		} else if userInputs.JobHopOrStay == "Stay" {
			if company.RetentionRate > 50 {
				// log.Println("stay", company.Name, company.RetentionRate)
				// log.Println("score", score)
				score++
			}
		}

		for _, value := range userInputs.MostImportantValues {
			if company.Values[value] == true {
				// log.Print("most important matched", value, "\n", company.Name, company.Values[value])
				// log.Println("score", score)
				score++
				score++
			}
			bestScore = score
		}

		fmt.Println("------- score is -------", score)
		fmt.Println("------- best score is -------", bestScore)

		if score > bestScore {
			log.Println("Previous best score", bestScore)
			log.Println("New best score", score)
			bestCompany = company
			bestScore = score
		}
	}

	return bestCompany
}

func companyData() []Company {
	jsonData := getData()
	companies := []Company{}

	for x := range jsonData {
		companies = append(companies, Company{
			Name:          jsonData[x].(map[string]interface{})["Name"].(string),
			Values:        jsonData[x].(map[string]interface{})["Values"].(map[string]interface{}),
			CompanySize:   jsonData[x].(map[string]interface{})["CompanySize"].(float64),
			RetentionRate: jsonData[x].(map[string]interface{})["RetentionRate"].(float64),
		})
	}
	return companies
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	name := r.FormValue("name")
	email := r.FormValue("email")
	ideasOrExpand := r.FormValue("ideasOrExpand")
	motivations := r.Form["motivations"]
	bigOrSmall := r.FormValue("bigOrSmall")
	jobHopOrStay := r.FormValue("jobHopOrStay")
	mostImportantValues := r.Form["mostImportantValues"]
	location := r.FormValue("location")
	jobTitle := r.FormValue("jobTitle")

	d := UserInputs{
		name,
		email,
		motivations,
		ideasOrExpand,
		bigOrSmall,
		jobHopOrStay,
		mostImportantValues,
		location,
		jobTitle,
	}
	companies := companyData()

	companyTest := getBestSuitedCompany(companies, d)

	returnData := ReturnData{
		name,
		email,
		motivations,
		ideasOrExpand,
		bigOrSmall,
		jobHopOrStay,
		mostImportantValues,
		location,
		jobTitle,
		companyTest,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	id := strconv.Itoa(r1.Intn(100))
	print(id)
	// Track("user-"+id, "assessment", &d, &returnData)
	// NEED TO CHANGE BELOW
	// FIXES THE UI NOT UPDATING BY SENDING ANOTHER REQUEST TO THE USER
	// NEED TO EITHER SPLIT THE IDENTIFY PART AND EVENT PART OR CALL THE METHOD TWICE
	Track("user-70", "assessment", &d, &returnData)

	tmpl, err := template.ParseFiles("result.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, returnData)
}

func main() {
	Amp()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", resultHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	log.Fatal(http.ListenAndServe(":8081", nil))
	//amp()
}
