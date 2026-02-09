package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func main() {
	total := 0
	raw,_:=os.ReadFile("data/stocks_index.json")
	total = len(raw)

	files,_:=filepath.Glob("data/intrinsic/*.json")
	out:=map[string]int{
		"intrinsic_files":len(files),
		"coverage_percent":len(files)*100/10338,
	}
	b,_:=json.MarshalIndent(out,"","  ")
	os.WriteFile("audit/intrinsic_progress.json",b,0644)
}
