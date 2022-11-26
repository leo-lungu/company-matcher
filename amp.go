package main

// Import amplitude package
import (
	"github.com/amplitude/analytics-go/amplitude"
)

type c struct {
	Name string
}

// amp func
func amp() {
	analytics := amplitude.NewClient(
		amplitude.NewConfig("ac610d38b345c833f241e1dc353c3691"),
	)

	analytics.Track(amplitude.Event{
		UserID:          "user-id",
		EventType:       "My Event",
		EventProperties: map[string]interface{}{"source": "notification"},
	})

	print(analytics)
}

func trial() {
	cc := []c{}
	cc = append(cc, c{
		Name: "asdas",
	})
}
