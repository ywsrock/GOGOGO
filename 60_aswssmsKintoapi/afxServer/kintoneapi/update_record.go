package kintoneapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type UpdateRecordRequest struct {
	App      int       `json:"app"`
	ID       int       `json:"id"`
	Record   RecordUdp `json:"record"`
	Revision *int      `json:"revision,omitempty"`
}

type UpdateRecordResponse struct {
	Revision string `json:"revision"`
}

func (kapi *KintoneClient) updateRecord(request *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	// url := "https://" + domain + "/k/v1/record.json"
	url := kapi.ApiEndpoint

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Cybozu-API-Token", kapi.ApiToken)

	client := kapi.HttpClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response UpdateRecordResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
