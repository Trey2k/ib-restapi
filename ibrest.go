package ibrestapi

import (
	"crypto/tls"
	"net/http"
	"time"
)

var isRunning bool = false

//Start will start the connection to the IB Client Port web API. you must pass in a error channel to test for connection loss
func Start(errChan chan error) error {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //Disabling SSL security check globally
	if isEndpointSet() {
		isRunning = true
		go keepAlive(errChan)
		return nil
	}
	return ErrEndpointNotSet
}

func keepAlive(errChan chan error) {
	for {
		ping, err := PingEndpoint()
		if err != nil {
			errChan <- err
		}

		if !ping {
			authStatus, err := IsAuthenticated()
			if err != nil {
				errChan <- err
			}
			if !authStatus {
				response, err := Reauthenticate()
				if err != nil {
					errChan <- err
				}
				if response.Message != "" {
					authStatus, err = IsAuthenticated()
					if err != nil {
						errChan <- err
					}
					if !authStatus {
						errChan <- ErrCantAuthenticate
					}
				}
			}
		}
		time.Sleep(time.Second * 30)
	}
}
