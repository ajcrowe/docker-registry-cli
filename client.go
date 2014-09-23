package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func newRequestGet(path string) []byte {
	url := formatURL(path)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	content := readRequest(resp)
	return content
}

func newRequestPut(path, data string) int {
	url := formatURL(path)

	client := &http.Client{}

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer([]byte(data)))
	setHeaders(req)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	return resp.StatusCode
}

func newRequestDelete(path string) int {
	url := formatURL(path)

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url, nil)
	setHeaders(req)

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

func formatURL(path string) string {
	return fmt.Sprintf("%s/%s/%s", INDEX_URL, INDEX_API_VERSION, path)
}

func setHeaders(req *http.Request) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
}
