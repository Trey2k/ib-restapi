package ibrestapi

import (
	"fmt"
	"strings"
)

//GetContractID get Contract ID by Symbol
func GetContractID(ticker string) (bool, int, error) {
	payload := SearchPayload{
		Symbol:  ticker,
		Name:    false,
		SecType: "",
	}
	found := false

	responses, err := Search(payload)
	if err != nil {
		return false, 0, err
	}

	var response SearchResponse
	for i := 0; i < len(responses); i++ {
		if strings.ToUpper(responses[i].Symbol) == strings.ToUpper(ticker) {
			response = responses[i]
			found = true
		}
	}
	return found, response.Conid, nil
}

//GetContractInfo get contract info by Contract IT
func GetContractInfo(conid int) (GetContractInfoResponse, error) {
	var response GetContractInfoResponse
	path := "/iserver/contract/" + fmt.Sprint(conid) + "/info"
	err := get(&response, path)
	return response, err
}

//Search by symbol or name
func Search(payloadStruct SearchPayload) (SearchResponses, error) {
	var response SearchResponses
	err := post(payloadStruct, &response, "/iserver/secdef/search")
	return response, err
}
