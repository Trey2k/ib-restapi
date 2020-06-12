package ibrestapi

//Tickle tickles the endpoint to keep the session alive
func Tickle() (TickleResponse, error) {
	var response TickleResponse
	err := get(&response, "/tickle")
	return response, err
}

//Logout loggs you out of cpw
func Logout() (LogoutResponse, error) {
	var response LogoutResponse
	err := get(&response, "/logout")
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

//IsAuthenticated tests if the session is currently authenticated
func IsAuthenticated() (bool, error) {
	response, err := GetAuthStatus()
	if err != nil {
		return false, err
	}
	return response.Authenticated, nil
}
