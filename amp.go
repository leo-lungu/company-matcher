package main

import (
	"github.com/amplitude/analytics-go/amplitude"
)

var analytics amplitude.Client

func Amp() {
	config := amplitude.NewConfig("ac610d38b345c833f241e1dc353c3691")
	config.FlushQueueSize = 200
	config.FlushInterval = 1
	x := amplitude.NewClient(config)

	analytics = x
}

func Track(userID, eventType string, d *UserInputs, r *ReturnData) {
	identify(userID, d, r)

	analytics.Track(amplitude.Event{
		UserID:          userID,
		EventType:       eventType,
		EventProperties: map[string]interface{}{},
		EventOptions:    amplitude.EventOptions{},
	})

}

func identify(userID string, d *UserInputs, r *ReturnData) {

	identifyObj := amplitude.Identify{}

	identifyObj.Set("UserInputs-name", d.Name)
	identifyObj.Set("UserInputs-email", d.Email)
	identifyObj.Set("UserInputs-ideasOrExpand", d.IdeasOrExpand)
	identifyObj.Set("UserInputs-motivations", d.Motivations)
	identifyObj.Set("UserInputs-bigOrSmall", d.BigOrSmall)
	identifyObj.Set("UserInputs-jobHopOrStay", d.JobHopOrStay)
	identifyObj.Set("UserInputs-mostImportantValues", d.MostImportantValues)
	identifyObj.Set("UserInputs-location", d.Location)
	identifyObj.Set("UserInputs-jobTitle", d.JobTitle)
	identifyObj.Set("Result-companyName", r.CompanyTest.Name)
	identifyObj.Set("Result-values", r.CompanyTest.Values)
	identifyObj.Set("Result-companySize", r.CompanyTest.CompanySize)
	identifyObj.Set("Result-retentionRate", r.CompanyTest.RetentionRate)

	analytics.Identify(identifyObj, amplitude.EventOptions{UserID: userID})
}
