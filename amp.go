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
func Amp() {
	config := amplitude.NewConfig("ac610d38b345c833f241e1dc353c3691")
	config.FlushQueueSize = 200
	x := amplitude.NewClient(config)

	analytics = x
}

func Track(userID, eventType string, ep map[string]interface{}) {
	analytics.Track(amplitude.Event{
		UserID:          userID,
		EventType:       eventType,
		EventProperties: ep,
	})
}

func Trial() {
	cc := []c{}
	cc = append(cc, c{
		Name: "asdas",
	})
}
