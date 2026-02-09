package main

import (
	"encoding/json"
	"math"
	"os"
)

type DCFInput struct {
	Ticker        string  `json:"ticker"`
	FCF           float64 `json:"fcf"`
	GrowthRate    float64 `json:"growth_rate"`
	DiscountRate  float64 `json:"discount_rate"`
	TerminalRate  float64 `json:"terminal_rate"`
	Years         int     `json:"years"`
}

type DCFOutput struct {
	Ticker          string  `json:"ticker"`
	IntrinsicValue  float64 `json:"intrinsic_value"`
}

func main() {
	input := DCFInput{
		Ticker:       "AAPL",
		FCF:          100000000000, // example FCF
		GrowthRate:   0.05,
		DiscountRate: 0.09,
		TerminalRate: 0.025,
		Years:        5,
	}

	value := 0.0
	fcf := input.FCF

	for t := 1; t <= input.Years; t++ {
		fcf *= (1 + input.GrowthRate)
		value += fcf / math.Pow(1+input.DiscountRate, float64(t))
	}

	terminal := (fcf * (1 + input.TerminalRate)) /
		(input.DiscountRate - input.TerminalRate)

	value += terminal / math.Pow(1+input.DiscountRate, float64(input.Years))

	output := DCFOutput{
		Ticker:         input.Ticker,
		IntrinsicValue: value,
	}

	out, _ := json.MarshalIndent(output, "", "  ")
	os.WriteFile("data/intrinsic_value.json", out, 0644)
}
