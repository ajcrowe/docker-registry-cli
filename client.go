package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func newRequestGet(path string) []byte {
	url := fmt.Sprintf("%s/%s/%s", INDEX_URL, INDEX_API_VERSION, path)
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	content := readRequest(resp)
	return content
}

func newRequestPut(path, data string) int {
	url := fmt.Sprintf("%s/%s/%s", INDEX_URL, INDEX_API_VERSION, path)
	log.Print(url)
	log.Print(data)

	client := &http.Client{}

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode
}

func readRequest(resp *http.Response) []byte {
	contents, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println(err)
	}
	return contents
}
