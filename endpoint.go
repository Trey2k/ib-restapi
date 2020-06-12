package ibrestapi

import (
	"fmt"
	"net"
)

var endpoint string

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
	if endpoint != "" {
		return true
	}
	return false
}

//Tickle tickles the endpoint to keep the session alive
func Tickle() (TickleResponse, error) {
	var response TickleResponse
	err := get(&response, "/tickle")
	return response, err
}

//GetAuthStatus gets the current authentication status
func GetAuthStatus() (AuthStatusResponse, error) {
	var response AuthStatusResponse
	err := get(&response, "/iserver/auth/status")
	return response, err
}

//Reauthenticate Attempts to reauthenticate the session
func Reauthenticate() (ReauthinticateResponse, error) {
	var response ReauthinticateResponse
	err := get(&response, "/iserver/reauthenticate")
	return response, err
}

//PingEndpoint pings the endpoint
func PingEndpoint() (bool, error) {
	response, err := Tickle()
	if err != nil {
		return false, err
	}
	return response.Iserver.Tickle, nil
}

//IsAuthenticated tests if the session is currently authenticated
func IsAuthenticated() (bool, error) {
	response, err := GetAuthStatus()
	if err != nil {
		return false, err
	}
	return response.Authenticated, nil
}
