package operations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// CouchbaseDoc parses .doc property from sync gateway documents
type CouchbaseDoc struct {
	Doc map[string]string `json:"doc"`
}

// Results deconstruct... results is a property of body, and is an array of obects
type Results struct {
	Results []byte `json:"results"`
}

func createURLs(channels []string) map[string]string {
	urlMap := make(map[string]string)

	domain := "swap" + strings.TrimSpace(os.Getenv("env"))
	bucket := strings.TrimSpace(os.Getenv("bucket"))
	for _, doctype := range channels {
		urlMap[doctype] = fmt.Sprintf("https://%s.endpoint.com/%s/_changes?filter=sync_gateway/bychannel&channels=%s&include_docs=true", domain, bucket, doctype)
	}

	return urlMap
}

func getChangesFeed(url string, ch chan map[string]string) {
	resp, _ := http.Get(url)

	body, _ := ioutil.ReadAll(resp.Body)

	go parseBody(body, ch)
}

func parseBody(body []byte, ch chan map[string]string) {
	var results Results
	var doc CouchbaseDoc
	json.Unmarshal(body, &results)
	json.Unmarshal(results.Results, &doc)
	// write to responses
	ch <- doc.Doc
}

func fetchDocs(channels []string) {
	urls := createURLs(channels)

	// Response channel is where all go routines will do the dirty
	responses := make(chan map[string]string)
	for _, url := range urls {
		go getChangesFeed(url, responses)
	}

	// Read from responses channel
	docs := <-responses
	for doc := range docs {
		fmt.Println(doc) // This should print results ??
	}
}
