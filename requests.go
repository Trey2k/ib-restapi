package ibrest

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func post(payloadStruct interface{}, responseStruct interface{}, path string) error {
	if isRunning {

		payload, err := json.Marshal(&payloadStruct)
		if err != nil {
			return err
		}

		resp, err := http.Post(endpoint+path, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		println(string(bodyBytes))
		err = jsonUnmarshal(bodyBytes, &responseStruct, resp.StatusCode)
		return err
	}
	return ErrNotRunning
}

func get(responseStruct interface{}, path string) error {

	if isRunning {

		resp, err := http.Get(endpoint + path)

		if err != nil {
			return err
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		err = jsonUnmarshal(bodyBytes, &responseStruct, resp.StatusCode)
		return err
	}
	return ErrNotRunning
}

func jsonUnmarshal(bodyBytes []byte, responseStruct interface{}, statusCode int) error {
	switch statusCode {

	case 200:

		err := json.Unmarshal(bodyBytes, &responseStruct)
		return err
	case 400:
		return ErrInitSession
	case 401:
		return ErrNotAuthenticated
	case 500:

		var errResp ErrorResponse

		err := json.Unmarshal(bodyBytes, &errResp)
		if err != nil {
			return err
		}

		return errors.New("Error while processing request: " + errResp.Error)
	default:
		return ErrUnkownResponseCode
	}
}
