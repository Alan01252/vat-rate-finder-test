package vat

import (
	"encoding/json"
	"testing"
)

const testJson = `[
{
   "startDate":"2017-01-01 00:00:00",
   "endDate":"2017-02-01 23:59:59",
   "vatRate":{
      "standard":20,
      "reduced":0.5
   }
},
{
   "startDate":"2017-02-02 00:00:00",
   "endDate":"2017-03-01 23:59:59",
   "vatRate":{
      "standard":30,
      "reduced":0.5
   }
}
]`

type MockJsonFetcher struct {
	url string
}

func (jsonFetcher *MockJsonFetcher) GetJson() (foundJson []interface{}, err error) {
	json.Unmarshal([]byte(testJson), &foundJson)
	return foundJson, nil
}

func TestVatRate_GetVatRate(t *testing.T) {

	jsonFetcher := &MockJsonFetcher{}
	v := VatRateFinder{}

	var dateTests = []struct {
		date     string // input
		expected int    // expected result
	}{
		{"2016-01-01", -1},
		{"2017-01-01", 20},
		{"2017-01-20", 20},
		{"2017-01-31", 20},
		{"2017-02-01", 20},
		{"2017-02-03", 30},
		{"2017-03-02", -1},
	}

	type foundVat float64;

	for _, tt := range dateTests {
		foundVat, _:= v.GetVatRate(jsonFetcher, tt.date)
		if foundVat != float64(tt.expected) {
			t.Error("Incorrect VAT rate for date", tt.date, "found expected", tt.expected, "got", foundVat)
		}
	}

}
