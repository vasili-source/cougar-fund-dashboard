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

type CompanyFacts struct {
	Facts map[string]map[string]map[string][]struct {
		Val float64 `json:"val"`
	} `json:"facts"`
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

// Try to extract recent FCF from SEC facts:
// FCF ≈ NetCashProvidedByOperatingActivities − CapitalExpenditures
func extractFCF(raw []byte) (float64, bool) {
	var cf CompanyFacts
	if err := json.Unmarshal(raw, &cf); err != nil {
		return 0, false
	}
	us := cf.Facts["us-gaap"]
	if us == nil {
		return 0, false
	}

	getLast := func(key string) (float64, bool) {
		m := us[key]
		if m == nil {
			return 0, false
		}
		u := m["USD"]
		if len(u) == 0 {
			return 0, false
		}
		return u[len(u)-1].Val, true
	}

	op, ok1 := getLast("NetCashProvidedByUsedInOperatingActivities")
	capex, ok2 := getLast("PaymentsToAcquirePropertyPlantAndEquipment")
	if ok1 && ok2 {
		return op - capex, true
	}
	return 0, false
}

func dcf(baseFCF, g, r, tg float64, years int) float64 {
	v := 0.0
	f := baseFCF
	for t := 1; t <= years; t++ {
		f *= (1 + g)
		v += f / math.Pow(1+r, float64(t))
	}
	term := (f * (1 + tg)) / (r - tg)
	v += term / math.Pow(1+r, float64(years))
	return v
}

func main() {
	const (
		batchSize = 40
		sleepMS  = 150
		years    = 5
		g        = 0.03
		r        = 0.09
		tg       = 0.02
	)

	raw, _ := os.ReadFile("data/stocks_index.json")
	stocks := []Stock{}
	json.Unmarshal(raw, &stocks)

	os.MkdirAll("data/intrinsic", 0755)

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
		cik := strconv.Itoa(s.CIK)
		for len(cik) < 10 {
			cik = "0" + cik
		}
		url := "https://data.sec.gov/api/xbrl/companyfacts/CIK" + cik + ".json"
		rawCF := fetch(url)
		if rawCF == nil {
			continue
		}
		fcf, ok := extractFCF(rawCF)
		if !ok || fcf <= 0 {
			continue
		}
		val := dcf(fcf, g, r, tg, years)
		out := map[string]any{
			"ticker": s.Ticker,
			"name":   s.Name,
			"value":  val,
		}
		j, _ := json.MarshalIndent(out, "", "  ")
		os.WriteFile("data/intrinsic/"+s.Ticker+".json", j, 0644)
		time.Sleep(time.Millisecond * sleepMS)
	}

	os.WriteFile("data/intrinsic/_cursor.txt", []byte(strconv.Itoa(end)), 0644)
}
