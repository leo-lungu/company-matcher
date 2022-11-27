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
	CompanySize   int
	RetentionRate int
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
			for _, companyValue := range company.Values {
				if value == companyValue {
					score++
					score++
					score++
				}
			}
		}

		if userInputs.IdeasOrExpand == "Ideas" {
			for _, companyValue := range company.Values {
				if companyValue == "Innovative" {
					score++
					score++
				}
			}
		}

		if userInputs.BigOrSmall == "Big" {
			if company.CompanySize >= 80000 {
				score++
			}
		} else {
			if company.CompanySize < 80000 {
				score++
			}
		}

		if userInputs.JobHopOrStay == "Stay" {
			if company.RetentionRate >= 75 {
				score++
				score++
			}
		} else {
			if company.RetentionRate < 75 {
				score++
			}
		}

		for _, value := range userInputs.MostImportantValues {
			for _, companyValue := range company.Values {
				if value == companyValue {
					score++
					score++
					score++
				}
			}
		}

		if score > bestScore {
			bestScore = score
			bestCompany = company
		}

		log.Println("Overall: ", company.Name, score)
	}

	fmt.Println("--------=-=-------?>>", bestScore)
	fmt.Println("--------=-=-------?>>", bestCompany)

	return bestCompany
}

func companyData() []Company {
	jsonData := getData()
	companies := []Company{}

	for x := range jsonData {
		companies = append(companies, Company{
			Name:          jsonData[x].(map[string]interface{})["Name"].(string),
			Values:        jsonData[x].(map[string]interface{})["Values"].(map[string]interface{}),
			CompanySize:   int(jsonData[x].(map[string]interface{})["CompanySize"].(float64)),
			RetentionRate: int(jsonData[x].(map[string]interface{})["RetentionRate"].(float64)),
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

	print(companies)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	id := strconv.Itoa(r1.Intn(100))
	Track("user-"+id, "assessment", &d, &returnData)

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
}
