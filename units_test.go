package gopower

import (
	"testing"
	"time"

	th "github.com/jrlangford/gopower/gopowertest"
)

var wattConversionCases = []th.TestCase[Power, float64]{
	{
		"watt conversion",
		Power(1),
		0.000_001,
	},
}

var kilowattConversionCases = []th.TestCase[Power, float64]{
	{
		"kilowatt conversion",
		Power(1),
		0.000_000_001,
	},
}

func TestPowerConversions(t *testing.T) {

	for _, tCase := range wattConversionCases {
		kw := tCase.Input.Watts()
		if kw != tCase.ExpectedOutput {
			t.Errorf(
				"\nDescription: %s\nExpected %v\nGot %v\n",
				tCase.Description,
				tCase.ExpectedOutput,
				kw,
			)
		}

	}

	for _, convCase := range kilowattConversionCases {
		kw := convCase.Input.Kilowatts()
		if kw != convCase.ExpectedOutput {
			t.Errorf(
				"Expected %v, got %v",
				convCase.ExpectedOutput,
				kw,
			)
		}

	}

}

var kilowattHourConversionCases = []th.TestCase[*Energy, float64]{
	{
		"kilowattHour conversion",
		NewEnergy(
			Power(1)*Kilowatt,
			time.Hour,
		),
		1.0,
	},
	{
		"kilowattHour conversion",
		NewEnergy(
			Power(1)*Kilowatt,
			time.Hour*2,
		),
		2.0,
	},
	{
		"kilowattHour conversion",
		NewEnergy(
			Power(2)*Kilowatt,
			time.Hour*2,
		),
		4.0,
	},
}

func TestChargeConversions(t *testing.T) {
	for _, convCase := range kilowattHourConversionCases {
		kwh := convCase.Input.KilowattsHour()
		if kwh != convCase.ExpectedOutput {
			t.Errorf(
				"Expected %v, got %v",
				convCase.ExpectedOutput,
				kwh,
			)
		}

	}
}

type testInput struct {
	e *Energy
	d time.Duration
}

var energyToPowerCases = []th.TestCase[testInput, Power]{
	{
		"1 KWh deliverd during 1 hour",
		testInput{
			NewEnergy(
				1*Kilowatt,
				time.Hour,
			),
			1 * time.Hour,
		},
		1 * Kilowatt,
	},
	{
		"1 KWh delivered during 30 mins",
		testInput{
			NewEnergy(
				1*Kilowatt,
				time.Hour,
			),
			30 * time.Minute,
		},
		2 * Kilowatt,
	},
	{
		"2 KWh delivered during 2 Hours",
		testInput{
			NewEnergy(
				2*Kilowatt,
				time.Hour,
			),
			2 * time.Hour,
		},
		1 * Kilowatt,
	},
	{
		"2 KWh delivered during 2 Hours, alternate definition of charge",
		testInput{
			NewEnergy(
				1*Kilowatt,
				2*time.Hour,
			),
			2 * time.Hour,
		},
		1 * Kilowatt,
	},
	{
		"53 KWh delivered during 15 minutes",
		testInput{
			NewEnergy(
				53*Kilowatt,
				time.Hour,
			),
			15 * time.Minute,
		},
		212 * Kilowatt,
	},
	{
		"-456 KWh delivered during 15 minutes",
		testInput{
			NewEnergy(
				-456*Kilowatt,
				time.Hour,
			),
			15 * time.Minute,
		},
		-1824 * Kilowatt,
	},
}

func TestEnergyToPowerConversions(t *testing.T) {
	for _, tCase := range energyToPowerCases {

		energy := tCase.Input.e
		duration := tCase.Input.d

		p := energy.ToPower(duration)

		if p != tCase.ExpectedOutput {
			t.Errorf(
				"\nDescription: %s\nExpected: %v\nGot: %v\n",
				tCase.Description,
				tCase.ExpectedOutput,
				p,
			)
		}

	}
}
