package gopower

import (
	"encoding/json"
	"fmt"
	"time"
)

// An Energy represents a charge to be delivered.
type Energy struct {
	power    Power
	duration time.Duration
}

// Creates a new Energy object.
func NewEnergy(power Power, duration time.Duration) *Energy {
	return &Energy{
		power:    power,
		duration: duration,
	}
}

func (e Energy) String() string {
	shortRep := NewReadableEnergyFromEnergy(e)
	return fmt.Sprintf("%s_%s", shortRep.Value, shortRep.Units)
}

func (e Energy) MarshalJSON() ([]byte, error) {
	rEnergy := NewReadableEnergyFromEnergy(e)
	return json.Marshal(rEnergy)
}

// Returns Energy in KiloWattsHour.
func (e *Energy) KilowattsHour() float64 {
	return e.power.Kilowatts() * e.duration.Hours()
}

// Returns the Power that corresponds to the Energy when delivered in a given time.
func (e *Energy) ToPower(d time.Duration) Power {
	ratio := float64(e.duration) / float64(d)
	return Power(float64(e.power) * ratio)
}
