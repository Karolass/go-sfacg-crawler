package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func ReadResponseBody(res *http.Response) []byte {
	buf := bytes.NewBuffer(make([]byte, 0, res.ContentLength))
	_, err := buf.ReadFrom(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return buf.Bytes()
}

func URLQueryFormatter(u string, path string, params map[string]string) string {
	URL, err := url.Parse(u)
	if err != nil {
		panic("boom")
	}

	URL.Path += path
	param := url.Values{}
	for key, value := range params {
		param.Add(key, value)
	}
	URL.RawQuery = param.Encode()

	return URL.String()
}

func RandomTime() (t time.Duration) {
	rand.Seed(time.Now().UnixNano())
	for {
		x := rand.Intn(2500)
		if x > 500 {
			t = time.Duration(x) * time.Millisecond
			return
		}
	}
}
