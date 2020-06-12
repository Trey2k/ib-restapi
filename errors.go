package ibrestapi

import "errors"

//ErrNotRunning ib-resapi is not runnning
var ErrNotRunning error = errors.New("ib-restapi is not running make sure to call Start() first")

//ErrInvalidIP Invlid IPv4
var ErrInvalidIP error = errors.New("IP is not a valid IPv4")

//ErrInvalidPort Invlid Port Number
var ErrInvalidPort error = errors.New("Port is not a valid port number")

//ErrEndpointNotSet Endpoint has not been set
var ErrEndpointNotSet error = errors.New("Endpoint not set make sure to SetEndpoint() first")

//ErrCantAuthenticate Lost connection to cpw
var ErrCantAuthenticate error = errors.New("Unable to authenticate session. PLease try restarting your cpw")

//ErrUnknowError Unkown error from cpw
var ErrUnknowError error = errors.New("cpw sent an unkown error")
