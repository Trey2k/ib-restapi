package ibrestapi

//TickleResponse is the response struct for /tickle
type TickleResponse struct {
	SsoExpires int  `json:"ssoExpires"`
	Collission bool `json:"collission"`
	UserID     int  `json:"userId"`
	Iserver    struct {
		Tickle     bool `json:"tickle"`
		AuthStatus struct {
			Authenticated bool   `json:"authenticated"`
			Competing     bool   `json:"competing"`
			Connected     bool   `json:"connected"`
			Message       string `json:"message"`
			MAC           string `json:"MAC"`
		} `json:"authStatus"`
	} `json:"iserver"`
}

//AuthStatusResponse response struct for /iserver/auth/status
type AuthStatusResponse struct {
	Authenticated bool   `json:"authenticated"`
	Competing     bool   `json:"competing"`
	Connected     bool   `json:"connected"`
	Message       string `json:"message"`
	MAC           string `json:"MAC"`
	Fail          string `json:"fail"`
}

//ReauthinticateResponse response struct for /iserver/reauthenticate
type ReauthinticateResponse struct {
	Message string `json:"message"`
}
