package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var (
	parseURL = "http://localhost:8080/parse/classes/"
)

func ParseCreate(url string, body io.Reader) (res *http.Response) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("X-Parse-Application-Id", "myAppId")
	req.Header.Set("Content-Type", "application/json")

	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (catalog *Catalog) create() {
	url := parseURL + "catalog"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(catalog)

	ParseCreate(url, body)
}
