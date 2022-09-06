package db

import (
	"pageredder/model"
	"reflect"
	"testing"
)

func init()  {
	DBInit()
}

//func setUP()  {
//	DBInit()
//}



func TestSaveWord(t *testing.T) {

	DBInit()

	type args struct {
		wordinfo *model.WordInfo
	}

	a := args{wordinfo: &model.WordInfo{Word: "what",Info: "test word",VoiceLink: "www",Memo: "test case 01"}}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name:"save01",args: a,wantErr: true,},
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
	//setUP()
	tests := []struct {
		name string
		want []model.WordInfo
	}{
		{name: "123",want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindAll(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() = %v, want %v", got, tt.want)
			}
		})
	}
}