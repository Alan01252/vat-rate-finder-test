package vat

import (
	"testing"

	"net/http"
	"net/http/httptest"
	"reflect"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/ginkgo/reporters"
)

const urlJson = `[
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

func TestJsonFetcher(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Vat Suite", []Reporter{junitReporter})
}

var _ = Describe("JSON URL Fetcher", func() {

	It("to return a VatRateStruct when returing valid json", func() {

		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(urlJson))
		}))
		defer ts.Close()

		jsonFetcher := NewUrlJsonFetcher(ts.URL)
		aStruct, _ := jsonFetcher.GetJson()

		Expect(reflect.TypeOf(aStruct).String()).To(Equal("vat.VatRateStruct"))

	})

	It("to return an error when given invalid json", func() {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("{invalidjson}"))
		}))
		defer ts.Close()

		jsonFetcher := NewUrlJsonFetcher(ts.URL)
		_, err := jsonFetcher.GetJson()

		Expect(err).ToNot(BeNil())

	})
})
