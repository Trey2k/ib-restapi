package ibrestapi

import (
	"crypto/tls"
	"net/http"
	"time"
)

var isRunning bool = false

//Start will start the connection to the IB Client Port web API. you must pass in a error channel to test for connection loss. PingDelay is how often in secounds you want to test the connection
func Start(errChan chan error, pingDelay int) error {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //Disabling SSL security check globally
	if isEndpointSet() {
		isRunning = true
		go keepAlive(errChan, pingDelay)
		return nil
	}
	return ErrEndpointNotSet
}

func keepAlive(errChan chan error, pingDelay int) {
	for {
		ping, err := PingEndpoint()
		if err != nil {
			errChan <- err
		}

		if !ping {
			auth, err := IsAuthenticated()
			if err != nil {
				errChan <- err
			}

			if !auth {
				response, err := Reauthenticate()
				if err != nil {
					errChan <- err
				}

				if response.Message != "" {
					auth, err = IsAuthenticated()
					if err != nil {
						errChan <- err
					}

					if !auth {
						errChan <- ErrCantAuthenticate
					}
				}
			}
		}
		time.Sleep(time.Second * time.Duration(pingDelay))
	}
}
