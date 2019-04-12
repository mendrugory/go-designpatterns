package observer

import "fmt"

// Temperature is a float64
type Temperature float64

// Thermometer struct
type Thermometer struct {
	trackers []chan Temperature
	tempChan chan Temperature
}

// TemperatureTracker tracks the temperature
type TemperatureTracker struct {
	tempChan          chan Temperature
	handleTemperature func(t Temperature)
}

// NewThermometer returns a new Thermometer
func NewThermometer() Thermometer {
	return Thermometer{
		tempChan: make(chan Temperature),
	}
}

// Start the Thermometer
func (t Thermometer) Start() {
	go func(c <-chan Temperature) {
		for {
			newTemperature := <-t.tempChan
			t.Notify(newTemperature)
		}
	}(t.tempChan)
}

// Register a tracker
func (t *Thermometer) Register(c chan Temperature) {
	t.trackers = append(t.trackers, c)
}

// Notify the new Temperature to all trackers
func (t Thermometer) Notify(temp Temperature) {
	for _, tracker := range t.trackers {
		tracker <- temp
	}
}

// NewTemperature a new Temperature is provided
func (t Thermometer) NewTemperature(temp Temperature) {
	t.tempChan <- temp
}

// NewTemperatureTracker starts a new tracker and returns it.
func NewTemperatureTracker(handle func(t Temperature)) TemperatureTracker {
	fmt.Println("Creating new Temperature Tracker ...")
	tempChan := make(chan Temperature)
	tracker := TemperatureTracker{
		tempChan:          tempChan,
		handleTemperature: handle,
	}

	go func(t *TemperatureTracker) {
		fmt.Println("Tracking Temperature ...")
		for {
			newTemperature := <-tracker.tempChan
			tracker.handleTemperature(newTemperature)
		}
	}(&tracker)

	return tracker

}
