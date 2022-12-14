package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"text/template"
	"time"
)

type Company struct {
	Name          string
	Values        map[string]interface{}
	CompanySize   int
	RetentionRate int
	Image         string
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
	CompanyData Company
}

type UserIDTracker struct {
	ID int `json:"id"`
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

var idTracker UserIDTracker

func (i *UserIDTracker) increment() int {
	i.ID++
	return i.ID
}

func (i *UserIDTracker) createIDString() string {
	id := i.increment()
	i.updateIDJSON()
	return "user-" + strconv.Itoa(id)
}

func (i *UserIDTracker) updateIDJSON() {
	marsh, err := json.Marshal(i)
	handleErr(err)
	_ = ioutil.WriteFile("id.json", marsh, 0644)
}

func loadIDJSON() UserIDTracker {
	readFile, err := ioutil.ReadFile("id.json")
	handleErr(err)

	x := UserIDTracker{}
	err = json.Unmarshal([]byte(readFile), &x)
	handleErr(err)

	return x
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
			Image:         jsonData[x].(map[string]interface{})["Image"].(string),
		})
	}
	return companies
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, r)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("result.html")
	handleErr(err)

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
	company := getBestSuitedCompany(companies, d)

	returnData := ReturnData{
		company,
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	id := strconv.Itoa(r1.Intn(1000))
	Track("user-"+id, "collection", &d, &returnData)

	tmpl.Execute(w, returnData)

	content := "From: " + "leo.lungu13@gmail.com" + "\n" + "To: " + email + "\n" + "Subject: Your Perfect Company\n\n" + "Hello " + name + ",\n\n" + "Thank you for taking the time to complete the assessment. We have found that the best company for you to work is: " + company.Name + "!\n\n" + "We hope you have a great day!\n\n" + "Best regards,\n" + "The team at Coman.py"

	err = smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", "leo.lungu13@gmail.com", "qebrqfgwpjntyceb", "smtp.gmail.com"),
		"leo.lungu13@gmail.com",
		[]string{email},
		[]byte(content))
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	Amp()
	idTracker = loadIDJSON()
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/result", resultHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("./img"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	log.Fatal(http.ListenAndServe(":8099", nil))
}
