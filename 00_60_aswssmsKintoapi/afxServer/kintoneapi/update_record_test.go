package kintoneapi

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"testing"
)

func TestKintoneClient_UpdateRecord(t *testing.T) {
	type args struct {
		request *UpdateRecordRequest
	}

	record := RecordUdp{
		RequestStatus: CommonType{
			// Type:  "DROP_DOWN",
			Value: "対応中",
		},
		// UpdateDate: CommonType{
		// 	// Type:  "UPDATED_TIME",
		// 	Value: "2023-03-29T05:49:00Z",
		// },
		ReceptionContents: CommonType{
			// Type:  "MULTI_LINE_TEXT",
			Value: "営業所から処理中",
		},
		RequestContents: CommonType{
			Value: "更新処理処理",
		},
		// UpdateUser: CommonType{
		// 	// Type: "MODIFIER",
		// 	Value: User{
		// 		Code: "en.b@di-square.co.jp",
		// 		Name: "YAN",
		// 	},
		// },
	}

	r := &UpdateRecordRequest{
		App:    AppID,
		ID:     76,
		Record: record,
	}

	kclient := &KintoneClient{
		ApiToken:    ApiToken,
		ApiEndpoint: fmt.Sprintf("%s%s", BaseURL, UPDATE_RECODE_URL),
		HttpClient:  &http.Client{},
	}

	tests := []struct {
		name    string
		kapi    *KintoneClient
		args    args
		want    string
		wantErr bool
	}{
		{name: "Update1", kapi: kclient, args: args{request: r}, want: "3", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.kapi.updateRecord(tt.args.request)
			if (err != nil) != tt.wantErr {
				log.Println(err)
				t.Errorf("KintoneClient.UpdateRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.Revision, tt.want) {
				t.Errorf("KintoneClient.UpdateRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
