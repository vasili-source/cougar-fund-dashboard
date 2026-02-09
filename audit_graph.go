package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type GraphAudit struct {
	Pages          []string            `json:"pages"`
	Adjacency      map[string][]string `json:"adjacency"`
	Unreachable    map[string][]string `json:"unreachable_from"`
	Conclusions    []string            `json:"conclusions"`
}

func main() {
	// Load previous audit
	raw, _ := os.ReadFile("audit/site_audit.json")
	prev := struct {
		Pages []string            `json:"pages"`
		Links map[string][]string `json:"links"`
	}{}
	json.Unmarshal(raw, &prev)

	graph := GraphAudit{
		Pages:       prev.Pages,
		Adjacency:   prev.Links,
		Unreachable: map[string][]string{},
	}

	// Reachability check (BFS from each page)
	for _, src := range prev.Pages {
		seen := map[string]bool{src: true}
		queue := []string{src}

		for len(queue) > 0 {
			n := queue[0]
			queue = queue[1:]
			for _, dst := range prev.Links[n] {
				if filepath.Ext(dst) == ".html" && !seen[dst] {
					seen[dst] = true
					queue = append(queue, dst)
				}
			}
		}

		for _, p := range prev.Pages {
			if !seen[p] {
				graph.Unreachable[src] = append(graph.Unreachable[src], p)
			}
		}
	}

	if len(graph.Unreachable) == 0 {
		graph.Conclusions = append(graph.Conclusions, "All pages mutually reachable via navigation")
	} else {
		graph.Conclusions = append(graph.Conclusions, "Some pages are not reachable from others")
	}

	os.MkdirAll("audit", 0755)
	out, _ := json.MarshalIndent(graph, "", "  ")
	os.WriteFile("audit/graph_audit.json", out, 0644)
}
