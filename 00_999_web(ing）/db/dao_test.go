package db

import (
	"pageredder/model"
	"reflect"
	"testing"
)

func init() {
	DBInit()
}

// func setUP()  {
//	DBInit()
// }

func TestSaveWord(t *testing.T) {

	DBInit()

	type args struct {
		wordinfo *model.WordInfo
	}

	a := args{wordinfo: &model.WordInfo{Word: "what", Info: "test word", VoiceLink: "www", Memo: "test case 01"}}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "save01", args: a, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveWord(tt.args.wordinfo); (err != nil) != tt.wantErr {
				t.Errorf("SaveWord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

}

func TestFindAll(t *testing.T) {
	// setUP()
	tests := []struct {
		name string
		want []model.WordInfo
	}{
		{name: "123", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindDayCount(t *testing.T) {
	tests := []struct {
		name string
		want []map[string]interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindDayCount(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDayCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindMonthCiybt(t *testing.T) {
	m := make([]map[string]interface{}, 1)
	tests := []struct {
		name string
		want []map[string]interface{}
	}{
		{name: "aa", want: m},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMonthCiybt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMonthCiybt() = %v, want %v", got, tt.want)
			}
		})
	}
}
