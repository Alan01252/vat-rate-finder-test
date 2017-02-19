package vat

import (
	"encoding/json"
	"testing"

	"errors"
	"fmt"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

func TestVat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "VAT Test Suite")
}

type MockJsonFetcher struct {
	url string
}

func (jsonFetcher *MockJsonFetcher) GetJson() (foundVatList VatRateStruct, err error) {
	json.Unmarshal([]byte(testJson), &foundVatList)
	return foundVatList, nil
}

type MockInvalidJsonFetcher struct {
	url string
}

func (jsonFetcher *MockInvalidJsonFetcher) GetJson() (foundVatList VatRateStruct, err error) {
	return nil, errors.New("Invalid Json Found")
}

var _ = Describe("VAT Rate Fetcher", func() {

	It("to return a new vatRateFinder", func() {

		Expect(reflect.TypeOf(NewVatRateFinder()).String()).To(Equal("*vat.VatRateFinder"))
	})

	It("to return the VAT rate for a specific date", func() {

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

		for _, tt := range dateTests {
			foundVat, _ := v.GetVatRate(jsonFetcher, tt.date)
			if foundVat != tt.expected {
				Expect(foundVat).To(Equal(tt.expected))
			}
		}

	})

	It("to return an error when there's invalid date", func() {
		jsonFetcher := &MockJsonFetcher{}
		v := VatRateFinder{}

		foundVat, err := v.GetVatRate(jsonFetcher, "somedate")

		Expect(err).NotTo(Equal(nil))
		Expect(foundVat).To(Equal(-1))
	})

	It("to return an error when there's invalid json found", func() {
		jsonFetcher := &MockInvalidJsonFetcher{}
		v := VatRateFinder{}

		_, err := v.GetVatRate(jsonFetcher, "2017-01-01")

		fmt.Println(err)
		Expect(err).ToNot(BeNil())
	})

})
