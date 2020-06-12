package main

import (
	"log"

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

	response, err := ib.GetAuthStatus()
	if err != nil {
		log.Panic(err)
	}
	println(response.Authenticated)

	<-make(chan struct{})
}
