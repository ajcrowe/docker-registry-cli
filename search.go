package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"log"
	"strconv"
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

	var JSONResults []byte
	JSONResults = NewRequestGet(fmt.Sprintf("search?q=%s", searchTerm))

	var results search
	json.Unmarshal(JSONResults, &results)
	log.Println("found: " + strconv.Itoa(results.ResultCount) + " when searching for " + results.Query)

	for _, result := range results.Results {
		log.Println("Name: " + result.Name + " Description: " + result.Description)
	}
}

func doSearchAll(c *cli.Context) {
	log.Println("listing all images in " + INDEX_URL)
	var JSONResults []byte
	JSONResults = NewRequestGet("search")

	var results search
	json.Unmarshal(JSONResults, &results)
	log.Println("found: " + strconv.Itoa(results.ResultCount))

	for _, result := range results.Results {
		log.Println("Name: " + result.Name + " Description: " + result.Description)
	}
}
