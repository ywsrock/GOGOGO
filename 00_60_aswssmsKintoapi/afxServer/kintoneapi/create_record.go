package kintoneapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type CreateRecordRequest struct {
	App    int          `json:"app"`
	Record RecordCreate `json:"record"`
}

func (kapi *KintoneClient) createRecord(request *CreateRecordRequest) (string, error) {
	url := kapi.ApiEndpoint

	body, err := json.Marshal(request)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Cybozu-API-Token", kapi.ApiToken)

	// client := &http.Client{}

	resp, err := kapi.HttpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
