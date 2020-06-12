package ibrestapi

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
