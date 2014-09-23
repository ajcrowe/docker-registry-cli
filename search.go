package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"text/tabwriter"
)

type search struct {
	Query       string         `json:"query"`
	ResultCount int            `json:"num_results"`
	Results     []searchResult `json:"results"`
}

type searchResult struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (s *search) printResults(w *tabwriter.Writer) {
	writeHeader(w, "Repository", "Description")
	for _, r := range s.Results {
		writeLine(w, r.Name, r.Description)
	}
}

func doSearch(c *cli.Context) {
	searchTerm := c.Args().First()

	var results search
	json.Unmarshal(newRequestGet(fmt.Sprintf("search?q=%s", searchTerm)), &results)

	w := getTabWriter()
	defer w.Flush()

	writeLine(w, fmt.Sprintf("%d results for: %s", results.ResultCount, results.Query))
	results.printResults(w)
}

func doSearchAll(c *cli.Context) {
	var results search
	json.Unmarshal(newRequestGet("search"), &results)

	w := getTabWriter()
	defer w.Flush()

	writeLine(w, fmt.Sprintf("%d results", results.ResultCount))
	results.printResults(w)
}
