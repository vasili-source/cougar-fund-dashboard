package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type ConvergenceReport struct {
	TimestampUTC       string `json:"timestamp_utc"`
	StocksIndexed      int    `json:"stocks_indexed"`
	IntrinsicComputed  int    `json:"intrinsic_computed"`
	CoveragePercent    int    `json:"coverage_percent"`
	Status             string `json:"status"`
}

type Stock struct {
	Ticker string `json:"ticker"`
}

func main() {
	// Load SEC universe
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

	status := "IN_PROGRESS"
	if percent >= 100 {
		status = "COMPLETE"
	}

	report := ConvergenceReport{
		TimestampUTC:      time.Now().UTC().Format(time.RFC3339),
		StocksIndexed:     indexed,
		IntrinsicComputed: computed,
		CoveragePercent:   percent,
		Status:            status,
	}

	os.MkdirAll("audit", 0755)
	out, _ := json.MarshalIndent(report, "", "  ")
	os.WriteFile("audit/intrinsic_convergence_report.json", out, 0644)
}
