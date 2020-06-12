package ibrestapi

import (
	"bytes"
	"encoding/json"
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

		if string(bodyBytes) != "" {
			err = json.Unmarshal(bodyBytes, &responseStruct)
			if err != nil {
				return err
			}

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

		if string(bodyBytes) != "" {
			err = json.Unmarshal(bodyBytes, &responseStruct)
			if err != nil {
				return err
			}

		}
		return nil

	}
	return ErrNotRunning
}
