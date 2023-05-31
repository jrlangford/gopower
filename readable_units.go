package gopower

import (
	"github.com/jrlangford/gopower/numeric"
)

type ReadablePower struct {
	Value string
	Units string
}

func NewReadablePower(shortVal string, units string) *ReadablePower {
	return &ReadablePower{
		Value: shortVal,
		Units: units,
	}
}

func NewReadablePowerFromPower(p Power) *ReadablePower {
	if p == ZeroPower {
		return NewReadablePower(
			"0",
			PowerUnits[ZeroPower],
		)
	}
	short, t := numeric.GetShortestRepresentationByThousands(int64(p))
	scale := numeric.IntPow(1000, t)
	return NewReadablePower(
		short,
		PowerUnits[Power(scale)],
	)
}

type ReadableEnergy struct {
	Value    string
	Units    string
	RepError float64
}

func NewReadableEnergy(shortVal string, units string, repError float64) *ReadableEnergy {
	return &ReadableEnergy{
		Value:    shortVal,
		Units:    units,
		RepError: repError,
	}
}

func NewReadableEnergyFromEnergy(e Energy) *ReadableEnergy {
	if e.power == ZeroPower {
		return NewReadableEnergy(
			"0",
			PowerUnits[ZeroPower]+"h",
			0,
		)
	}
	energyHour := float64(e.power) * e.duration.Hours()
	intEnergyHour := int64(energyHour)
	energyError := energyHour - float64(intEnergyHour)
	shortEnergy, t := numeric.GetShortestRepresentationByThousands(intEnergyHour)
	scale := numeric.IntPow(1000, t)
	return NewReadableEnergy(
		shortEnergy,
		PowerUnits[Power(scale)]+"h",
		energyError,
	)
}
