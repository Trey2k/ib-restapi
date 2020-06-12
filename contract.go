package ibrestapi

//Search search by symbol or name
func Search(payloadStruct SearchPayload) (SearchResponses, error) {
	var response SearchResponses
	err := post(payloadStruct, &response, "/iserver/secdef/search")
	return response, err
}
