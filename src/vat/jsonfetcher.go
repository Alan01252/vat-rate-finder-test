package vat

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type JsonFetcher interface {
	GetJson() (foundJson VatRateStruct, err error)
}

type UrlJsonFetcher struct {
	url string
}

func (jsonFetcher *UrlJsonFetcher) GetJson() (foundJson VatRateStruct, err error) {
	response, _ := http.Get(jsonFetcher.url)
	defer response.Body.Close()

	htmlData, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(htmlData, &foundJson); err != nil {
		log.Print("Error parsing json:", err)
		return nil, err
	}

	return foundJson, nil
}

func NewUrlJsonFetcher(url string) *UrlJsonFetcher {
	jsonFetcher := UrlJsonFetcher{}
	jsonFetcher.url = url

	return &jsonFetcher
}
