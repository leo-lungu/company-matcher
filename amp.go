package main

// Import amplitude package
import (
	"github.com/amplitude/analytics-go/amplitude"
)

type c struct {
	Name string
}

var analytics amplitude.Client

// amp func
func amp() {
	config := amplitude.NewConfig("ac610d38b345c833f241e1dc353c3691")
	config.FlushQueueSize = 200
	x := amplitude.NewClient(config)

	analytics = x
}

func track(userID string, eventType string, ep map[string]interface{}) {
	analytics.Track({
		UserID: "",
		EventType: "",
		EventProperties: ep
	})
}

func trial() {
	cc := []c{}
	cc = append(cc, c{
		Name: "asdas",
	})
}
