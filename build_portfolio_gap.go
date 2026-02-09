package main

import (
	"encoding/json"
	"os"
)

func main() {
	// Example portfolio weights (can be replaced later)
	portfolio := map[string]float64{
		"AAPL": 0.4,
		"MSFT": 0.6,
	}

	results := []map[string]any{}

	for ticker, weight := range portfolio {
		raw, err := os.ReadFile("data/intrinsic/" + ticker + ".json")
		if err != nil {
			continue
		}
		v := map[string]any{}
		json.Unmarshal(raw, &v)

		results = append(results, map[string]any{
			"ticker":    ticker,
			"weight":    weight,
			"intrinsic": v["value"],
		})
	}

	out, _ := json.MarshalIndent(results, "", "  ")
	os.WriteFile("data/portfolio_gap.json", out, 0644)
}
