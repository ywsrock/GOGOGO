package kintoneapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetRecordsRequest struct {
	App int `json:"app,omitempty"`
	// ID  int `json:"id,omitempty"`
	// Fields     []string `json:"fields,omitempty"`
	// Query      string   `json:"query,omitempty"`
	TotalCount bool `json:"totalCount,omitempty"`
}

type GetRecordsResponse struct {
	Record     []RecordQuery `json:"records"`
	TotalCount string        `json:"totalCount"`
}

func (kapi *KintoneClient) getRecords(request *GetRecordsRequest) (*GetRecordsResponse, error) {
	url := kapi.ApiEndpoint

	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(body))
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

	var response GetRecordsResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
