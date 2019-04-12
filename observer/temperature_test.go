package observer

import (
	"testing"
)

func TestOneTracker(t *testing.T) {
	resultChan := make(chan bool)
	var actualTemperature Temperature = 25.0
	handler := func(temperature Temperature) {
		resultChan <- actualTemperature == temperature
	}
	thermometer := NewThermometer()
	tracker := NewTemperatureTracker(handler)
	thermometer.Register(tracker.tempChan)

	thermometer.Start()
	thermometer.NewTemperature(actualTemperature)

	if result := <-resultChan; !result {
		t.Errorf("Temperatures are not the same.\n")
	}
}
