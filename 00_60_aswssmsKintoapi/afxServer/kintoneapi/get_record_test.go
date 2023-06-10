package kintoneapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

const jsonData = `
{
	"record": {
	  "request_status": {
		"type": "DROP_DOWN",
		"value": "受付"
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
	  "$revision": {
		"type": "__REVISION__",
		"value": "1"
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
		"value": "Test"
	  },
	  "request_contents": {
		"type": "MULTI_LINE_TEXT",
		"value": "問い合わせ内容"
	  },
	  "record_no": {
		"type": "RECORD_NUMBER",
		"value": "1"
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
	  },
	  "$id": {
		"type": "__ID__",
		"value": "1"
	  }
	}
  }

`

func TestKintoneClient_GetRecord(t *testing.T) {
	type args struct {
		request *GetRecordRequest
	}

	r := &GetRecordRequest{
		App: AppID,
		ID:  3,
	}

	kclient := &KintoneClient{
		ApiToken:    ApiToken,
		ApiEndpoint: fmt.Sprintf("%s%s", BaseURL, GET_RECODE_URL),
		HttpClient:  &http.Client{},
	}

	var res GetRecordResponse
	err := json.Unmarshal([]byte(jsonDataUpdate), &res)
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name    string
		kapi    *KintoneClient
		args    args
		want    *GetRecordResponse
		wantErr bool
	}{
		{
			name:    "Get1",
			kapi:    kclient,
			args:    args{r},
			want:    &res,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.kapi.getRecord(tt.args.request)
			fmt.Printf("%s\n", got.Record.RequestContents)
			if (err != nil) != tt.wantErr {
				t.Errorf("KintoneClient.GetRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Record.RequestStatus, tt.want.Record.RequestStatus) {
				t.Errorf("KintoneClient.GetRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
