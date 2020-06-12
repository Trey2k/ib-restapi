package main

import (
	"log"
	"strings"

	ib "github.com/Trey2k/ib-restapi"
)

func main() {
	var err error
	errChan := make(chan error)

	//Setting Endpoint
	err = ib.SetEndpoint("127.0.0.1", 5000)
	if err != nil {
		log.Panic(err)
	}

	//Starting connection, passing our error channel and setting the ping delay
	err = ib.Start(errChan, 30)
	if err != nil {
		log.Panic(err)
	}

	//Holind call in go routine listening on the error channel
	go func(errChan chan error) {
		err := <-errChan
		if err != nil {
			log.Panic(err)
		}
	}(errChan)

	printVerify("AMZN")
	printVerify("GO")
	printVerify("A")
	printVerify("HH")
	printVerify("IBM")

	//Holding call to keep connection alive for testing purposes. You would not normally use this
	<-make(chan struct{})
}

//Using the contract Search to search for contracts with given ticker aka symbol
func verifyTicker(ticker string) (bool, ib.SearchResponse) {
	//Creating payload for the search function
	payload := ib.SearchPayload{
		Symbol:  ticker,
		Name:    false,
		SecType: "",
	}
	real := false

	//Running the search function and passing it the SearchPayload we created above
	responses, err := ib.Search(payload) //Search returns type of ib.SearchResponses which is just a array of the type ib.SearchResponse this will be true for any type that returns an array
	if err != nil {
		log.Panic(err)
	}

	//Finding which if any have a exact match
	var response ib.SearchResponse
	for i := 0; i < len(responses); i++ {
		if strings.ToUpper(responses[i].Symbol) == strings.ToUpper(ticker) {
			response = responses[i]
			real = true
		}
	}
	return real, response
}

//Running the verifyTicker we made above but priniting it out nicley.
func printVerify(ticker string) {
	real, response := verifyTicker(ticker)
	if real {
		println("[" + ticker + "] is a real ticker it belongs to " + response.CompanyHeader)
	} else {
		println("[" + ticker + "] is not a real ticker")
	}
}
