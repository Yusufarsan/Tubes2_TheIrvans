package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func makeRequest(url string) *goquery.Document {
	res, err := http.Get(url)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	if err != nil {
		return nil
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		return nil
	}

	return doc
}
