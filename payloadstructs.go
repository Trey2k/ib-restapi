package ibrestapi

//SearchPayload the payload struct for Search
type SearchPayload struct {
	Symbol  string `json:"symbol"`
	Name    bool   `json:"name"`
	SecType string `json:"secType"`
}
