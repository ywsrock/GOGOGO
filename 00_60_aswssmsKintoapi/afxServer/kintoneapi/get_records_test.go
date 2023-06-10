package kintoneapi

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"

	"github.com/goccy/go-json"
)

const JsonData_all = `
{
    "records": [
        {
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
                "value": "APIから作成"
            },
            "request_contents": {
                "type": "MULTI_LINE_TEXT",
                "value": "問い合わせ内容問い合わせ内容問い合わせ内容問い合わせ内容"
            },
            "record_no": {
                "type": "RECORD_NUMBER",
                "value": "3"
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
                "value": "3"
            }
        },
        {
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
                "value": "2"
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
                "value": "2"
            }
        },
        {
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
            "$revision": {
                "type": "__REVISION__",
                "value": "3"
            },
            "update_date": {
                "type": "UPDATED_TIME",
                "value": "2023-03-30T01:38:00Z"
            },
            "reception_contents": {
                "type": "MULTI_LINE_TEXT",
                "value": "営業所から処理中"
            },
            "update_user": {
                "type": "MODIFIER",
                "value": {
                    "code": "Administrator",
                    "name": "Administrator"
                }
            },
            "request_title": {
                "type": "SINGLE_LINE_TEXT",
                "value": "Test"
            },
            "request_contents": {
                "type": "MULTI_LINE_TEXT",
                "value": "更新処理処理"
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
    ],
    "totalCount": "3"
}

`

func TestKintoneClient_GetRecords(t *testing.T) {
	type args struct {
		request *GetRecordsRequest
	}

	r := &GetRecordsRequest{
		App:        AppID,
		TotalCount: true,
	}

	var w GetRecordsResponse
	err := json.Unmarshal([]byte(JsonData_all), &w)
	if err != nil {
		log.Fatalln(err)
	}

	kclient := &KintoneClient{
		ApiToken:    ApiToken,
		ApiEndpoint: fmt.Sprintf("%s%s", BaseURL, GET_RECODES_URL),
		HttpClient:  &http.Client{},
	}

	tests := []struct {
		name    string
		kapi    *KintoneClient
		args    args
		want    *GetRecordsResponse
		wantErr bool
	}{
		{
			name: "GETALL1",
			kapi: kclient,
			args: args{
				request: r,
			},
			want:    &w,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.kapi.getRecords(tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("KintoneClient.GetRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.TotalCount, tt.want.TotalCount) {
				t.Errorf("KintoneClient.GetRecords() = %v, want %v", got, tt.want)
			}
		})
	}
}
