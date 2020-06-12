package ibrestapi

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"time"
)

var isRunning bool = false
var endpoint string

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

//Endpoint functions

//SetEndpoint is used to set the IPv4 and Port of the ib cpw endpoint
func SetEndpoint(ip string, port int64) error {
	if net.ParseIP(ip) != nil {
		if port != 0 {
			endpoint = "https://" + ip + ":" + fmt.Sprint(port) + "/v1/portal"
			return nil
		}
		return ErrInvalidPort
	}
	return ErrInvalidIP
}

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
