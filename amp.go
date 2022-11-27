package main

// Import amplitude package
import (
	"time"

	"github.com/amplitude/analytics-go/amplitude"
)

type c struct {
	Name string
}

var analytics amplitude.Client

// amp func
func Amp() {
	config := amplitude.NewConfig("ac610d38b345c833f241e1dc353c3691")
	config.FlushQueueSize = 200
	x := amplitude.NewClient(config)

	analytics = x
}

func Track(userID, eventType string, d *UserInputs, r *ReturnData) {

	analytics.Track(amplitude.Event{
		UserID:          userID,
		EventType:       eventType,
		EventProperties: map[string]interface{}{},
		EventOptions:    amplitude.EventOptions{},
	})

	time.Sleep(1)

	identify(userID, d, r)

}

func identify(userID string, d *UserInputs, r *ReturnData) {

	identifyObj := amplitude.Identify{}

	// for UserInputs
	identifyObj.Set("UserInputs-name", d.Name)
	identifyObj.Set("UserInputs-email", d.Email)
	identifyObj.Set("UserInputs-ideasOrExpand", d.IdeasOrExpand)
	identifyObj.Set("UserInputs-motivations", d.Motivations)
	identifyObj.Set("UserInputs-bigOrSmall", d.BigOrSmall)
	identifyObj.Set("UserInputs-jobHopOrStay", d.JobHopOrStay)
	identifyObj.Set("UserInputs-mostImportantValues", d.MostImportantValues)
	identifyObj.Set("UserInputs-location", d.Location)
	identifyObj.Set("UserInputs-jobTitle", d.JobTitle)

	// for resultant data
	identifyObj.Set("Result-companyName", r.CompanyTest.Name)
	identifyObj.Set("Result-values", r.CompanyTest.Values)
	identifyObj.Set("Result-companySize", r.CompanyTest.CompanySize)
	identifyObj.Set("Result-retentionRate", r.CompanyTest.RetentionRate)

	analytics.Identify(identifyObj, amplitude.EventOptions{UserID: userID})
}

func Trial() {
	cc := []c{}
	cc = append(cc, c{
		Name: "asdas",
	})
}
