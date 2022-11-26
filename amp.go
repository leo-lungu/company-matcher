package main

// Import amplitude package
import (
	"github.com/amplitude/analytics-go/amplitude"
)

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
