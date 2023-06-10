package kintoneapi

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type GetRecordRequest struct {
	App int `json:"app"`
	ID  int `json:"id"`
}

type GetRecordResponse struct {
	Record RecordQuery `json:"record"`
}

func (kapi *KintoneClient) getRecord(request *GetRecordRequest) (*GetRecordResponse, error) {
	// url := "https://" + domain + "/k/v1/record.json?" + "app=" + strconv.Itoa(request.App) + "&id=" + strconv.Itoa(request.ID)

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

	var response GetRecordResponse

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
