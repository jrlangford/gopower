package gopower

import (
	"encoding/json"
	"fmt"
)

// A Power represents electric power in microwatts.
type Power int64

const (
	minPower Power = -1 << 63  //-9,223,372,036,854,775,808
	maxPower       = 1<<63 - 1 //9,223,372,036,854,775,807
)

// Common amounts of Power.
const (
	ZeroPower Power = 0
	Microwatt       = 1
	Milliwatt       = 1000 * Microwatt
	Watt            = 1000 * Milliwatt
	Kilowatt        = 1000 * Watt
	Megawatt        = 1000 * Kilowatt
	Gigawatt        = 1000 * Megawatt
	Terawatt        = 1000 * Gigawatt
)

var PowerUnits = map[Power]string{
	ZeroPower: "W",
	Microwatt: "µW",
	Milliwatt: "mW",
	Watt:      "W",
	Kilowatt:  "kW",
	Megawatt:  "MW",
	Gigawatt:  "GW",
	Terawatt:  "TW",
}

func NewPower(nanowatts int64) Power {
	return Power(nanowatts)
}

func (p Power) String() string {
	rPower := NewReadablePowerFromPower(p)
	return fmt.Sprintf("%s%s", rPower.Value, rPower.Units)
}

func (p Power) MarshalJSON() ([]byte, error) {
	rPower := NewReadablePowerFromPower(p)
	return json.Marshal(rPower)
}

// Returns power in Watts.
func (p Power) Watts() float64 {
	return float64(p) / float64(Watt)
}

// Returns power in KiloWatts.
func (p Power) Kilowatts() float64 {
	return float64(p) / float64(Kilowatt)
}

// Returns power in MegaWatts.
func (p Power) Megawatts() float64 {
	return float64(p) / float64(Megawatt)
}

// Returns power in GigaWatts.
func (p Power) Gigawatts() float64 {
	return float64(p) / float64(Gigawatt)
}

// Returns power in Terawatts.
func (p Power) Terwatts() float64 {
	return float64(p) / float64(Terawatt)
}
