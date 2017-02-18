package main

import (
	"fmt"
	"os"
	"vat"

	log "github.com/Sirupsen/logrus"
	"logger"
)

var requestLogger *log.Entry

func init() {
	requestLogger = logger.GetLogger("main")
}

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please enter a date string in the format YYYY-MM-DD")
		os.Exit(-1)
	}

	requestLogger.WithFields(log.Fields{
		"date": args[0],
	}).Debug("User requested vat rate for date")

	jsonFetcher := vat.NewUrlJsonFetcher("http://192.168.37.211/index.html")
	vatRateFinder := vat.NewVatRateFinder()

	foundVatRate, err := vatRateFinder.GetVatRate(jsonFetcher, args[0])
	if err != nil {
		fmt.Println(err)
	} else {

		requestLogger.WithFields(log.Fields{
			"date":      args[0],
			"foundrate": foundVatRate,
		}).Debug("User requested vat rate for date")
		fmt.Println("Found VAT rate of", foundVatRate)
	}
}
