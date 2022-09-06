package main

import (
	"testing"
)

func TestOpenTemplateExcel(t *testing.T) {
	type args struct {
		excel string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case1",
			args: args{
				excel: "template.xlsx",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OpenTemplateExcel(tt.args.excel)
		})
	}
}

func TestTrimF(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{r: ' '}, want: true},
		{name: "2", args: args{r: '"'}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimF(tt.args.r); got != tt.want {
				t.Errorf("TrimF() = %v, want %v", got, tt.want)
			}
		})
	}
}
