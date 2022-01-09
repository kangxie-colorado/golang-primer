package main

import (
	"reflect"
	"testing"
)

func Test_getBinaryString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"12"}, "00010010"},
		{"", args{"A"}, "1010"},

		{"", args{"D2FE28"}, "110100101111111000101000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBinaryString(tt.args.input); got != tt.want {
				t.Errorf("getBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binStrToInt(t *testing.T) {
	type args struct {
		bins string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"00000000011"}, 3},
		{"", args{"000000000011011"}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binStrToInt(tt.args.bins); got != tt.want {
				t.Errorf("binStrToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_processLiteralPayload(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{"", args{"101111111000101000"}, 2021, 15},
		{"", args{"01010"}, 10, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := processLiteralPayload(tt.args.payload)
			if got != tt.want {
				t.Errorf("processLiteralPayload() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("processLiteralPayload() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_processOperatorPayload(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name  string
		args  args
		want  []int
		want1 int
	}{
		// TODO: Add test cases.
		{"", args{"00000000000110111101000101001010010001001000000000"}, []int{6, 2}, 43},
		{"", args{"10000000001101010000001100100000100011000001100000"}, []int{2, 4, 1}, 45},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := processOperatorPayload(tt.args.payload)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processOperatorPayload() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("processOperatorPayload() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
