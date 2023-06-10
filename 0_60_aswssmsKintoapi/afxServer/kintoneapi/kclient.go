package kintoneapi

import (
	"fmt"
	"net/http"
)

type KintoneClient struct {
	ApiToken    string
	ApiEndpoint string
	// HttpRequest func(method, url string, body []byte) (*http.Response, error)
	HttpClient *http.Client
}

func NewKintoneClient(uri string) *KintoneClient {
	kclient := &KintoneClient{
		ApiToken:    ApiToken,
		ApiEndpoint: fmt.Sprintf("%s%s", BaseURL, uri),
		HttpClient:  &http.Client{},
	}
	return kclient
}

// 新規
func (k *KintoneClient) Create(request *CreateRecordRequest) (string, error) {
	return k.createRecord(request)
}

// 更新
func (k *KintoneClient) Update(request *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return k.updateRecord(request)
}

// 検索1件
func (k *KintoneClient) Query(request *GetRecordRequest) (*GetRecordResponse, error) {
	return k.getRecord(request)
}

// 検索複数件
func (k *KintoneClient) QueryAll(request *GetRecordsRequest) (*GetRecordsResponse, error) {
	return k.getRecords(request)
}
