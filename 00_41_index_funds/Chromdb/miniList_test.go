package Chromdb

import (
	"testing"
)

func TestGetMini(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
	}{
		{"test", args{url: "https://finance.yahoo.co.jp/indices/list"}},
		// {"test", args{url: "https://fu.minkabu.jp/index-price"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetMini(tt.args.url)
		})
	}
}
