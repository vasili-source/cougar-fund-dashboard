package main

import (
	"encoding/json"
	"os"
)

func main() {
	// Academic placeholders; real SEC ratios vary by availability
	out := map[string]any{
		"AAPL": map[string]float64{
			"PE":        28.4,
			"EV_EBITDA": 21.1,
		},
		"MSFT": map[string]float64{
			"PE":        31.2,
			"EV_EBITDA": 24.0,
		},
	}
	b,_:=json.MarshalIndent(out,"","  ")
	os.WriteFile("data/comparables.json",b,0644)
}
