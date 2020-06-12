package ibrestapi

//Geneirc Structs

//ErrorResponse response struct
type ErrorResponse struct {
	Error string
}

//Start Session Structs

//TickleResponse is the response struct for /tickle
type TickleResponse struct {
	SsoExpires int
	Collission bool
	UserID     int
	Iserver    struct {
		Tickle     bool
		AuthStatus AuthStatusResponse
	}
}

//LogoutResponse is the return struct for /logout
type LogoutResponse struct {
	Confirmed bool
}

//AuthStatusResponse response struct for /iserver/auth/status
type AuthStatusResponse struct {
	Authenticated bool
	Competing     bool
	Connected     bool
	Message       string
	MAC           string
	Fail          string
}

//ReauthinticateResponse response struct for /iserver/reauthenticate
type ReauthinticateResponse struct {
	Authenticated bool
	Connected     bool
	Competing     bool
	Fail          string
	Message       string
	Prompts       []string
}

//End Session Structs
//Start Contract Structs

//GetContractInfoResponse response struct for /iserver/contract/{conid}/info
type GetContractInfoResponse struct {
	RTH            bool
	ConID          int
	CompanyName    string
	Exchange       string
	LocalSymbol    string
	InstrumentType string
	Currency       string
	Category       string
	Industry       string
	Rules          struct {
		OrderTypes        []string
		OrderTypesOutside []string
		DefaultSize       int
		SizeIncrement     int
		TifTypes          []string
		LimitPrice        int
		Stopprice         int
		Preview           bool
		DisplaySize       string
		Increment         float64
	}
}

//SearchResponses array of SearchResponse
type SearchResponses []SearchResponse

//SearchResponse response struct for /iserver/secdef/search
type SearchResponse struct {
	Conid         int
	CompanyHeader string
	CompanyName   string
	Symbol        string
	Description   string
	Opt           string
	War           string
	Error         string
	Sections      []struct {
	}
}
