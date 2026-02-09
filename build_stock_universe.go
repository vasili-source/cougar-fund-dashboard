package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type TickerEntry struct {
	CIK    int    `json:"cik_str"`
	Ticker string `json:"ticker"`
	Name   string `json:"title"`
}

func main() {
	req, _ := http.NewRequest(
		"GET",
		"https://www.sec.gov/files/company_tickers.json",
		nil,
	)
	req.Header.Set(
		"User-Agent",
		"Academic Research Dashboard (vasil016@csusm.edu)",
	)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	raw := map[string]TickerEntry{}
	json.Unmarshal(body, &raw)

	list := []TickerEntry{}
	for _, v := range raw {
		list = append(list, v)
	}

	os.MkdirAll("data", 0755)
	out, _ := json.MarshalIndent(list, "", "  ")
	os.WriteFile("data/stocks_index.json", out, 0644)
}
