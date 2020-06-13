package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"strings"

	"github.com/Trey2k/ibrest"
)

func main() {
	var err error
	errChan := make(chan error)

	//Setting Endpoint IP, Port, SSL
	err = ibrest.SetEndpoint("127.0.0.1", 5000, true)
	if err != nil {
		log.Panic(err)
	}

	//I dont have a cert setup on my cpw but am still running in SSL mod, tis line disables SSL security check globally
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	//Starting connection, passing our error channel and setting the ping delay
	err = ibrest.Start(errChan, 30)
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
func verifyTicker(ticker string) (bool, ibrest.SearchResponse) {
	//Creating payload for the search function
	payload := ibrest.SearchPayload{
		Symbol:  ticker,
		Name:    false,
		SecType: "",
	}
	real := false

	//Running the search function and passing it the SearchPayload we created above
	responses, err := ibrest.Search(payload) //Search returns type of ib.SearchResponses which is just a array of the type ib.SearchResponse this will be true for any type that returns an array
	if err != nil {
		log.Panic(err)
	}

	//Finding which if any have a exact match
	var response ibrest.SearchResponse
	for i := 0; i < len(responses); i++ {
		if strings.ToUpper(responses[i].Symbol) == strings.ToUpper(ticker) {
			response = responses[i]
			println(response.Conid)
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
