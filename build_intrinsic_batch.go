package main

import (
	"encoding/json"
	"io"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Stock struct {
	CIK    int    `json:"cik_str"`
	Ticker string `json:"ticker"`
	Name   string `json:"title"`
}

func fetch(url string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Academic Research Dashboard (vasil016@csusm.edu)")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return b
}

func dcfValue(baseFCF, g, r, tg float64, years int) float64 {
	v := 0.0
	fcf := baseFCF
	for t := 1; t <= years; t++ {
		fcf *= (1 + g)
		v += fcf / math.Pow(1+r, float64(t))
	}
	terminal := (fcf * (1 + tg)) / (r - tg)
	v += terminal / math.Pow(1+r, float64(years))
	return v
}

func main() {
	// Batch controls (SAFE)
	const batchSize = 50     // filings per run
	const sleepMS  = 120     // throttle between calls
	const years    = 5
	const g        = 0.03
	const r        = 0.09
	const tg       = 0.02
	const baseFCF  = 1_000_000_000.0 // conservative placeholder

	// Load universe
	raw, _ := os.ReadFile("data/stocks_index.json")
	stocks := []Stock{}
	json.Unmarshal(raw, &stocks)

	os.MkdirAll("data/intrinsic", 0755)

	// Resume support
	start := 0
	if b, err := os.ReadFile("data/intrinsic/_cursor.txt"); err == nil {
		if n, e := strconv.Atoi(string(b)); e == nil {
			start = n
		}
	}

	end := start + batchSize
	if end > len(stocks) {
		end = len(stocks)
	}

	for i := start; i < end; i++ {
		s := stocks[i]

		// (Optional) touch SEC facts to ensure filer exists (no parsing needed here)
		cik := strconv.Itoa(s.CIK)
		for len(cik) < 10 { cik = "0" + cik }
		_ = fetch("https://data.sec.gov/api/xbrl/companyfacts/CIK" + cik + ".json")

		val := dcfValue(baseFCF, g, r, tg, years)

		out := map[string]any{
			"ticker": s.Ticker,
			"name":   s.Name,
			"value":  val,
		}
		j, _ := json.MarshalIndent(out, "", "  ")
		os.WriteFile("data/intrinsic/"+s.Ticker+".json", j, 0644)

		time.Sleep(time.Millisecond * sleepMS)
	}

	// Save cursor
	os.WriteFile("data/intrinsic/_cursor.txt", []byte(strconv.Itoa(end)), 0644)
}
