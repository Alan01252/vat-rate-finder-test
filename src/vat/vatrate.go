package vat

import (
	"errors"
	"time"
)

const shortForm = "2006-01-02 15:04:05"

type VatRateFinder struct {
	jsonFetcher JsonFetcher
}

func NewVatRateFinder() *VatRateFinder {
	vatRateFinder := VatRateFinder{}
	return &vatRateFinder
}

func findVatRateInJson(json []interface{}, parsedDate time.Time) (foundVatRate float64) {

	foundVatRate = -1

	for _, v := range json {

		vatRate := v.(map[string]interface{})
		// Assuming that the time in VAT json is always correct
		startDate, _ := time.Parse(shortForm, vatRate["startDate"].(string))
		endDate, _ := time.Parse(shortForm, vatRate["endDate"].(string))

		if parsedDate.Unix() >= startDate.Unix() && parsedDate.Unix() <= endDate.Unix() {
			vatDetails := vatRate["vatRate"].(map[string]interface{})
			return vatDetails["standard"].(float64)
		}
	}

	return foundVatRate
}

func (vatRateFinder *VatRateFinder) GetVatRate(jsonFetcher JsonFetcher, requestedDate string) (foundVatRate float64, err error) {

	foundVatRate = -1
	parsedDate, err := time.Parse(shortForm, requestedDate+" 00:00:00")

	if err != nil {
		return foundVatRate, err
	}

	foundJson, err := jsonFetcher.GetJson()
	if err != nil {
		return foundVatRate, err
	}

	foundVatRate = findVatRateInJson(foundJson, parsedDate)
	if foundVatRate == -1 {
		return foundVatRate, errors.New("Could not find vat rate")
	}

	return foundVatRate, nil
}
