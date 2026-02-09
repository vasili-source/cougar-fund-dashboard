package main

import (
    "encoding/json"
    "io"
    "net/http"
    "os"
    "time"
)func fetch(url string, headers map[string]string) []byte {
	req, _ := http.NewRequest("GET", url, nil)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte("{}")
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	return body
}

func main() {
	os.MkdirAll("data", 0755)

	// --- SEC: company facts (example: Apple) ---
	secURL := "https://data.sec.gov/api/xbrl/companyfacts/CIK0000320193.json"
	secHeaders := map[string]string{
		"User-Agent": "CSUSM Cougar Fund (vasil016@csusm.edu)",
	}
	secData := fetch(secURL, secHeaders)
	os.WriteFile("data/sec_company_facts.json", secData, 0644)

	// --- FRED: Fed Funds Rate ---
	fredURL := "https://api.stlouisfed.org/fred/series/observations?series_id=FEDFUNDS&api_key=" + os.Getenv("FRED_API_KEY") + "&file_type=json"
	fredData := fetch(fredURL, nil)
	os.WriteFile("data/fed_funds.json", fredData, 0644)

	// --- Metadata ---
	meta := map[string]string{
		"last_update_utc": time.Now().UTC().Format(time.RFC3339),
		"source":          "SEC EDGAR + FRED",
	}
	metaBytes, _ := json.MarshalIndent(meta, "", "  ")
	os.WriteFile("data/update_meta.json", metaBytes, 0644)
}
