package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewRequestGet(rURL string) []byte {
	url := fmt.Sprintf("%s/%s/%s", INDEX_URL, INDEX_API_VERSION, rURL)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	content := readRequest(resp)
	return content
}

func readRequest(resp *http.Response) []byte {
	contents, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	return contents
}
