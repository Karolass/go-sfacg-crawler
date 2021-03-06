package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	parseURL = "http://localhost:8080/parse/classes/"
)

/* Parse */
func ParseFind(url string) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	// defer req.Body.Close()

	req.Header.Set("X-Parse-Application-Id", "myAppId")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	return ReadResponseBody(res)
}

func ParseCreate(url string, body io.Reader) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	req.Header.Set("X-Parse-Application-Id", "myAppId")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	return ReadResponseBody(res)
}

func ParseUpdate(url string, body io.Reader) []byte {
	client := &http.Client{}

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		log.Fatalln(err)
	}
	defer req.Body.Close()

	req.Header.Set("X-Parse-Application-Id", "myAppId")
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	return ReadResponseBody(res)
}

/* Catalog */
func (catalogs *Catalogs) find(skip, limit int) {
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
		"skip":  fmt.Sprintf("%d", skip),
	}
	URL := URLQueryFormatter(parseURL, "catalog", params)

	bytes := ParseFind(URL)

	type results struct {
		Results Catalogs
		Count   int
	}
	var c = new(results)
	json.Unmarshal(bytes, &c)
	*catalogs = c.Results
}

func (catalogs *Catalogs) count() int {
	params := map[string]string{
		"limit": fmt.Sprintf("%d", 0),
		"count": fmt.Sprintf("%d", 1),
	}
	URL := URLQueryFormatter(parseURL, "catalog", params)

	bytes := ParseFind(URL)

	type results struct {
		Results Catalogs
		Count   int
	}
	var c = new(results)
	json.Unmarshal(bytes, &c)
	return c.Count
}

func (catalog *Catalog) create() string {
	url := parseURL + "catalog"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(catalog)

	bytes := ParseCreate(url, body)

	type results struct {
		ObjectId  string
		CreatedAt time.Time
	}
	r := new(results)
	json.Unmarshal(bytes, &r)

	return r.ObjectId
}

/* Chapter */
func (chapters *Chapters) find(skip, limit int, pageNull bool) {
	params := map[string]string{
		"limit": fmt.Sprintf("%d", limit),
		"skip":  fmt.Sprintf("%d", skip),
	}
	if pageNull {
		params["where"] = `{"pages":null}`
	}
	URL := URLQueryFormatter(parseURL, "chapter", params)

	bytes := ParseFind(URL)

	type results struct {
		Results Chapters
		Count   int
	}
	var c = new(results)
	json.Unmarshal(bytes, &c)
	*chapters = c.Results
}

func (chapters *Chapters) count(pageNull bool) int {
	params := map[string]string{
		"limit": fmt.Sprintf("%d", 0),
		"count": fmt.Sprintf("%d", 1),
	}
	if pageNull {
		params["where"] = `{"pages":null}`
	}
	URL := URLQueryFormatter(parseURL, "chapter", params)

	bytes := ParseFind(URL)

	type results struct {
		Results Chapters
		Count   int
	}
	var c = new(results)
	json.Unmarshal(bytes, &c)
	return c.Count
}

func (chapter *Chapter) create() string {
	url := parseURL + "chapter"
	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(chapter)
	bytes := ParseCreate(url, body)

	type results struct {
		ObjectId  string
		CreatedAt time.Time
	}
	r := new(results)
	json.Unmarshal(bytes, &r)

	return r.ObjectId
}

func (chapter *Chapter) update(chapterPage ChapterPage) time.Time {
	url := parseURL + "chapter/" + chapter.ObjectId

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(chapterPage)

	bytes := ParseUpdate(url, body)

	type results struct {
		UpdatedAt time.Time
	}
	r := new(results)
	json.Unmarshal(bytes, &r)

	return r.UpdatedAt
}

func (chapter *Chapter) addRelation(objId string, cObjId string) time.Time {
	url := parseURL + "chapter/" + objId
	jsonStr := fmt.Sprintf(`{
                    "catalog": {
                        "__op": "AddRelation",
                        "objects": [
                            {
                                "__type": "Pointer",
                                "className": "catalog",
                                "objectId": "%s"
                            }
                        ]
                    }
                }`, cObjId)
	body := bytes.NewBufferString(jsonStr)

	bytes := ParseUpdate(url, body)

	type results struct {
		UpdatedAt time.Time
	}
	r := new(results)
	json.Unmarshal(bytes, &r)

	return r.UpdatedAt
}
