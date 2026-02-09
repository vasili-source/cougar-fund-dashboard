package main

import (
	"encoding/json"
	"math"
	"os"
)

type Metrics struct {
	AnnualReturn      float64 `json:"annual_return"`
	AnnualVolatility  float64 `json:"annual_volatility"`
	SharpeRatio       float64 `json:"sharpe_ratio"`
	MaxDrawdown       float64 `json:"max_drawdown"`
	RealRate          float64 `json:"real_rate"`
}

func main() {
	// Example placeholder inputs (will be wired to real data next)
	dailyReturns := []float64{0.001, -0.002, 0.0005, 0.0012}
	riskFreeRate := 0.03   // annual
	inflation := 0.025    // annual

	// Mean return
	var sum float64
	for _, r := range dailyReturns {
		sum += r
	}
	meanDaily := sum / float64(len(dailyReturns))
	annualReturn := meanDaily * 252

	// Volatility
	var variance float64
	for _, r := range dailyReturns {
		variance += math.Pow(r-meanDaily, 2)
	}
	variance /= float64(len(dailyReturns))
	annualVol := math.Sqrt(variance) * math.Sqrt(252)

	// Sharpe
	sharpe := (annualReturn - riskFreeRate) / annualVol

	// Max drawdown
	equity := 1.0
	peak := 1.0
	maxDD := 0.0
	for _, r := range dailyReturns {
		equity *= (1 + r)
		if equity > peak {
			peak = equity
		}
		dd := (peak - equity) / peak
		if dd > maxDD {
			maxDD = dd
		}
	}

	metrics := Metrics{
		AnnualReturn:     annualReturn,
		AnnualVolatility: annualVol,
		SharpeRatio:      sharpe,
		MaxDrawdown:      maxDD,
		RealRate:         riskFreeRate - inflation,
	}

	out, _ := json.MarshalIndent(metrics, "", "  ")
	os.WriteFile("data/clean_metrics.json", out, 0644)
}
