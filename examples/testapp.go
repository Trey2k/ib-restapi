package main

import (
	"log"
	"strings"

	ib "github.com/Trey2k/ib-restapi"
)

func main() {
	var err error
	errChan := make(chan error)
	err = ib.SetEndpoint("127.0.0.1", 5000)
	if err != nil {
		log.Panic(err)
	}

	err = ib.Start(errChan, 30)
	if err != nil {
		log.Panic(err)
	}

	go func(errChan chan error) {
		err := <-errChan
		if err != nil {
			log.Panic(err)
		}
	}(errChan)
	printVerify("AMZ")
	printVerify("GO")
	printVerify("A")
	printVerify("HH")
	printVerify("IBM")
	resp, err := ib.PingEndpoint()
	if err != nil {
		log.Panic(err)
	}
	println(resp)

	<-make(chan struct{})
}

func verifyTicker(ticker string) (bool, ib.SearchResponse) {
	payload := ib.SearchPayload{
		Symbol:  ticker,
		Name:    false,
		SecType: "",
	}
	real := false

	responses, err := ib.Search(payload)
	if err != nil {
		log.Panic(err)
	}

	var response ib.SearchResponse
	for i := 0; i < len(responses); i++ {
		if strings.ToUpper(responses[i].Symbol) == strings.ToUpper(ticker) {
			response = responses[i]
			real = true
		}
	}
	return real, response
}

func printVerify(ticker string) {
	real, response := verifyTicker(ticker)
	if real {
		println("[" + ticker + "] is a real ticker it belongs to " + response.CompanyHeader)
	} else {
		println("[" + ticker + "] is not a real ticker")
	}
}
