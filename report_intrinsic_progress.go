package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Stock struct {
	Ticker string `json:"ticker"`
}

type Progress struct {
	StocksIndexed     int `json:"stocks_indexed"`
	IntrinsicComputed int `json:"intrinsic_computed"`
	CoveragePercent   int `json:"coverage_percent"`
}

func main() {
	// Load stock universe properly
	raw, _ := os.ReadFile("data/stocks_index.json")
	stocks := []Stock{}
	json.Unmarshal(raw, &stocks)
	indexed := len(stocks)

	// Count intrinsic files
	files, _ := filepath.Glob("data/intrinsic/*.json")
	computed := len(files)

	percent := 0
	if indexed > 0 {
		percent = computed * 100 / indexed
	}

	out := Progress{
		StocksIndexed:     indexed,
		IntrinsicComputed: computed,
		CoveragePercent:   percent,
	}

	b, _ := json.MarshalIndent(out, "", "  ")
	os.WriteFile("audit/intrinsic_progress.json", b, 0644)
} 