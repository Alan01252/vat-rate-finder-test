package vat

import (
	"errors"
	"time"
)

const shortForm = "2006-01-02 15:04:05"

type VatRateStruct []struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	VatRate   struct {
		Standard int     `json:"standard"`
		Reduced  float64 `json:"reduced"`
	} `json:"vatRate"`
}

type VatRateFinder struct {
	jsonFetcher JsonFetcher
}

func NewVatRateFinder() *VatRateFinder {
	vatRateFinder := VatRateFinder{}
	return &vatRateFinder
}

func findVatRateInJson(foundVatList VatRateStruct, parsedDate time.Time) (foundVatRate int) {

	foundVatRate = -1

	for _, v := range foundVatList {

		// Assuming that the time in VAT json is always correct
		startDate, _ := time.Parse(shortForm, v.StartDate)
		endDate, _ := time.Parse(shortForm, v.EndDate)

		if parsedDate.Unix() >= startDate.Unix() && parsedDate.Unix() <= endDate.Unix() {
			vatDetails := v.VatRate
			return vatDetails.Standard
		}
	}

	return foundVatRate
}

func (vatRateFinder *VatRateFinder) GetVatRate(jsonFetcher JsonFetcher, requestedDate string) (foundVatRate int, err error) {

	foundVatRate = -1
	parsedDate, err := time.Parse(shortForm, requestedDate+" 00:00:00")

	if err != nil {
		return foundVatRate, err
	}

	foundVatList, err := jsonFetcher.GetJson()
	if err != nil {
		return foundVatRate, err
	}

	foundVatRate = findVatRateInJson(foundVatList, parsedDate)
	if foundVatRate == -1 {
		return foundVatRate, errors.New("Could not find vat rate")
	}

	return foundVatRate, nil
}
