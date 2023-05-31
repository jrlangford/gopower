package gopower

import (
	"testing"
	"time"

	th "github.com/jrlangford/gopower/gopowertest"
)

var readablePowerCases = []th.TestCase[Power, ReadablePower]{
	{
		"0 kW",
		0 * Kilowatt,
		ReadablePower{
			Value: "0",
			Units: "W",
		},
	},
	{
		"1 kW",
		1 * Kilowatt,
		ReadablePower{
			Value: "1",
			Units: "kW",
		},
	},
	{
		"2000_000 W",
		2000_000 * Watt,
		ReadablePower{
			Value: "2",
			Units: "MW",
		},
	},
	{
		"245_100_000 W",
		245_100_000 * Watt,
		ReadablePower{
			Value: "245100",
			Units: "kW",
		},
	},
	{
		"-0 W",
		-0 * Watt,
		ReadablePower{
			Value: "0",
			Units: "W",
		},
	},
	{
		"-6000 W",
		-6000 * Watt,
		ReadablePower{
			Value: "-6",
			Units: "kW",
		},
	},
	{
		"-6000 MW",
		-6000 * Megawatt,
		ReadablePower{
			Value: "-6",
			Units: "GW",
		},
	},
	{
		"-17_234_000 W",
		-17_234_000 * Watt,
		ReadablePower{
			Value: "-17234",
			Units: "kW",
		},
	},
}

func TestReadablePower(t *testing.T) {
	for _, tCase := range readablePowerCases {

		rPower := *NewReadablePowerFromPower(tCase.Input)

		if rPower != tCase.ExpectedOutput {
			t.Errorf(
				"\nDescription: %s\nExpected: %v\nGot: %v\n",
				tCase.Description,
				tCase.ExpectedOutput,
				rPower,
			)
		}

	}
}

var readableEnergyCases = []th.TestCase[Energy, ReadableEnergy]{
	{
		"0 kW for 5 hours",
		*NewEnergy(
			0*Kilowatt,
			5*time.Hour,
		),
		ReadableEnergy{
			Value:    "0",
			Units:    "Wh",
			RepError: 0,
		},
	},
	{
		"1 kW for 1 hour",
		*NewEnergy(
			1*Kilowatt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "1",
			Units:    "kWh",
			RepError: 0,
		},
	},
	{
		"2000_000 W for 1 hour",
		*NewEnergy(
			2000_000*Watt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "2",
			Units:    "MWh",
			RepError: 0,
		},
	},
	{
		"245_100_000 W for 1 hour",
		*NewEnergy(
			245_100_000*Watt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "245100",
			Units:    "kWh",
			RepError: 0,
		},
	},
	{
		"-0 W for 1 hour",
		*NewEnergy(
			-0*Watt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "0",
			Units:    "Wh",
			RepError: 0,
		},
	},
	{
		"-6000 W for 1 hour",
		*NewEnergy(
			-6000*Watt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "-6",
			Units:    "kWh",
			RepError: 0,
		},
	},
	{
		"-6000 MW for 1 hour",
		*NewEnergy(
			-6000*Megawatt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "-6",
			Units:    "GWh",
			RepError: 0,
		},
	},
	{
		"-17_234_000 W for 1 hour",
		*NewEnergy(
			-17_234_000*Watt,
			time.Hour,
		),
		ReadableEnergy{
			Value:    "-17234",
			Units:    "kWh",
			RepError: 0,
		},
	},
}

func TestReadableEnergy(t *testing.T) {
	for _, tCase := range readableEnergyCases {

		rEnergy := *NewReadableEnergyFromEnergy(tCase.Input)

		if rEnergy != tCase.ExpectedOutput {
			t.Errorf(
				"\nDescription: %s\nExpected: %v\nGot: %v\n",
				tCase.Description,
				tCase.ExpectedOutput,
				rEnergy,
			)
		}

	}
}
