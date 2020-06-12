package ibrestapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func post(payloadStruct interface{}, responseStruct interface{}, path string) error {
	if isRunning {
		var err error

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

		if resp.StatusCode == 200 {
			err = json.Unmarshal(bodyBytes, &responseStruct)
			if err != nil {
				return err
			}
		} else if resp.StatusCode == 500 {
			var errResp ErrorResponse

			err = json.Unmarshal(bodyBytes, &errResp)
			if err != nil {
				return err
			}
			return errors.New("Error while processing request: " + errResp.Error)
		} else {
			return ErrUnkownResponseCode
		}
		return nil

	}
	return ErrNotRunning
}

func get(responseStruct interface{}, path string) error {

	if isRunning {
		var err error
		resp, err := http.Get(endpoint + path)

		if err != nil {
			return err
		}

		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if resp.StatusCode == 200 {
			err = json.Unmarshal(bodyBytes, &responseStruct)
			if err != nil {
				return err
			}
		} else if resp.StatusCode == 500 {
			var errResp ErrorResponse

			err = json.Unmarshal(bodyBytes, &errResp)
			if err != nil {
				return err
			}
			return errors.New("Error while processing request: " + errResp.Error)
		} else {
			return ErrUnkownResponseCode
		}
		return nil

	}
	return ErrNotRunning
}
