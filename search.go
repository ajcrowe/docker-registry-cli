package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
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

func doSearch(c *cli.Context) {
	searchTerm := c.Args().First()

	var results search
	json.Unmarshal(newRequestGet(fmt.Sprintf("search?q=%s", searchTerm)), &results)

	w := getTabWriter()
	defer w.Flush()

	writeLine(w, fmt.Sprintf("%d results for: %s", results.ResultCount, results.Query))
	writeHeader(w, "Repository", "Description")

	for _, result := range results.Results {
		writeLine(w, result.Name, result.Description)
	}
}

func doSearchAll(c *cli.Context) {
	var results search
	json.Unmarshal(newRequestGet("search"), &results)

	w := getTabWriter()
	defer w.Flush()

	writeLine(w, fmt.Sprintf("%d results", results.ResultCount))
	writeHeader(w, "Repository", "Description")

	for _, result := range results.Results {
		writeLine(w, result.Name, result.Description)
	}
}
