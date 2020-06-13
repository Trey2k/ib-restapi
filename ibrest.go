package ibrest

import (
	"fmt"
	"net"
	"time"
)

var isRunning bool = false
var endpoint string

//Start will start the connection to the IB Client Port web API. You must pass in a error channel to test for connection loss. PingDelay is how often in secounds you want to test the connection.
func Start(errChan chan error, pingDelay int) error {
	if isEndpointSet() {
		isRunning = true
		ping, err := PingEndpoint()
		if err != nil {
			return err
		}

		if ping {
			go keepAlive(errChan, pingDelay)
			return nil
		}
		return ErrCantConnect
	}
	return ErrEndpointNotSet
}

//keepAlive is ran in a go routine in the backgroun to maintain the connection.
//It passes all errors to the error channel passed when Start() was ran
func keepAlive(errChan chan error, pingDelay int) {
	for {
		ping, err := PingEndpoint()
		if err != nil {
			errChan <- err
		}

		if !ping { //If ping fails try to reauthenticate
			response, err := Reauthenticate()
			if err != nil {
				errChan <- err
			}

			if !response.Authenticated { //If we cant authenticate send a error through the error channel
				errChan <- ErrCantAuthenticate
			}
		}

		time.Sleep(time.Second * time.Duration(pingDelay))
	}
}

//Endpoint functions

//SetEndpoint is used to set the IPv4 and Port of the ib cpw endpoint
func SetEndpoint(ip string, port int, ssl bool) error {
	if net.ParseIP(ip) != nil {
		if port != 0 {
			if ssl {
				endpoint = "https://" + ip + ":" + fmt.Sprint(port) + "/v1/portal"
			} else {
				endpoint = "http://" + ip + ":" + fmt.Sprint(port) + "/v1/portal"
			}
			return nil
		}
		return ErrInvalidPort
	}
	return ErrInvalidIP
}

//isEndpointSet test if the endpoint has been set
func isEndpointSet() bool {
	return endpoint != ""
}

//PingEndpoint pings the endpoint
func PingEndpoint() (bool, error) {
	response, err := Tickle()
	if err != nil {
		return false, err
	}
	return response.Iserver.Tickle, nil
}
