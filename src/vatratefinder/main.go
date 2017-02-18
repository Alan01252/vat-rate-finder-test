package main

import (
	"fmt"
	"os"
	"vat"
	log "github.com/Sirupsen/logrus"
	"reflect"
)

var requestLogger *log.Entry


func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
	requestLogger = log.WithFields(log.Fields{"package": "main"})
	fmt.Println(reflect.TypeOf(requestLogger))

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
