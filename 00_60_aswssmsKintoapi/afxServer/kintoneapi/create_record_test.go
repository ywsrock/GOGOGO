package kintoneapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const jsonDataUpdate = `
{
	"record": {
	  "request_status": {
		"type": "DROP_DOWN",
		"value": "対応中"
	  },
	  "reception_user": {
		"type": "USER_SELECT",
		"value": [
		  {
			"code": "en.b@di-square.co.jp",
			"name": "YAN"
		  }
		]
	  },
	  "update_date": {
		"type": "UPDATED_TIME",
		"value": "2023-03-29T05:49:00Z"
	  },
	  "reception_contents": {
		"type": "MULTI_LINE_TEXT",
		"value": ""
	  },
	  "update_user": {
		"type": "MODIFIER",
		"value": {
		  "code": "en.b@di-square.co.jp",
		  "name": "YAN"
		}
	  },
	  "request_title": {
		"type": "SINGLE_LINE_TEXT",
		"value": "APIから作成"
	  },
	  "request_contents": {
		"type": "MULTI_LINE_TEXT",
		"value": "問い合わせ内容問い合わせ内容問い合わせ内容問い合わせ内容"
	  },
	  "create_user": {
		"type": "CREATOR",
		"value": {
		  "code": "en.b@di-square.co.jp",
		  "name": "YAN"
		}
	  },
	  "create_date": {
		"type": "CREATED_TIME",
		"value": "2023-03-29T05:49:00Z"
	  },
	  "group": {
		"type": "GROUP_SELECT",
		"value": [
		  {
			"code": "Administrators",
			"name": "Administrators"
		  }
		]
	  }
	}
  }
`

func TestKintoneClient_CreateRecord(t *testing.T) {
	type args struct {
		apiToken string
		domain   string
		request  *CreateRecordRequest
	}

	var recordWrapper struct {
		Record RecordCreate `json:"record"`
	}

	err := json.Unmarshal([]byte(jsonDataUpdate), &recordWrapper)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v\n", recordWrapper)

	kclient := &KintoneClient{
		ApiToken:    ApiToken,
		ApiEndpoint: fmt.Sprintf("%s%s", BaseURL, ADD_RECODE_URL),
		HttpClient:  &http.Client{},
	}
	r := &CreateRecordRequest{
		App:    int(AppID),
		Record: recordWrapper.Record,
	}

	w := `
	{"id": "2",	"revision": "1" }
	  `
	tests := []struct {
		name    string
		kapi    *KintoneClient
		args    args
		want    string
		wantErr bool
	}{
		{name: "ADD1", kapi: kclient,
			args: args{apiToken: ApiToken, domain: "w1ylayksvqqk", request: r}, want: w, wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.kapi.createRecord(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("KintoneClient.CreateRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KintoneClient.CreateRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
