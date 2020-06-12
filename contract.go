package ibrestapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

//Search search by symbol or name
func Search(payloadStruct SearchPayload) (SearchResponses, error) {
	var response SearchResponses
	if isRunning {
		var err error
		payload, err := json.Marshal(&payloadStruct)
		if err != nil {
			return response, err
		}

		resp, err := http.Post(endpoint+"/iserver/secdef/search", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return response, err
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return response, err
		}

		if string(bodyBytes) != "" && strings.Contains(string(bodyBytes), `"error":`) == false {
			err = json.Unmarshal(bodyBytes, &response)
			if err != nil {
				return response, err
			}
		}

		return response, nil
	}
	return response, ErrNotRunning
}
