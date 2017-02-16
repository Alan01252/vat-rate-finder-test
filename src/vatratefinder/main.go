package main

import (
	"fmt"
	"os"
	"vat"
)

func main() {

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Please enter a date string in the format YYYY-MM-DD")
		os.Exit(-1)
	}

	jsonFetcher := vat.NewUrlJsonFetcher("http://192.168.37.211/index.html")
	vatRateFinder := vat.NewVatRateFinder()

	foundVatRate, err := vatRateFinder.GetVatRate(jsonFetcher, args[0])
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Found VAT rate of", foundVatRate)
	}
}
