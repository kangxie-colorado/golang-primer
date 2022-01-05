package main

import (
	"testing"
)

func Test_getPowerRates(t *testing.T) {
	type args struct {
		bitLen int
		nums   []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
		{"the example", args{5, []int{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}}, 22, 9},
		{"", args{5, []int{4, 30, 22}}, 22, 9},
		{"", args{5, []int{4, 30, 22, 23}}, 22, 9},
		{"", args{5, []int{4, 30, 22, 23, 21}}, 22, 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getPowerRates(tt.args.bitLen, tt.args.nums)
			if got != tt.want {
				t.Errorf("getRates() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getRates() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getO2Rate(t *testing.T) {
	type args struct {
		bitLen int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"the example", args{5, []int{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}}, 23},
		{"the example", args{5, []int{4, 30, 22, 23}}, 23},
		{"the example", args{5, []int{4, 30, 22}}, 30},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getO2Rate(tt.args.bitLen, tt.args.nums); got != tt.want {
				t.Errorf("getO2Rate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCO2Rate(t *testing.T) {
	type args struct {
		bitLen int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"the example", args{5, []int{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}}, 10},
		{"the example", args{5, []int{4, 30, 22, 23}}, 4},
		{"the example", args{5, []int{4, 30, 22}}, 4},
		{"the example", args{5, []int{30, 22, 23}}, 0},
		{"the example", args{5, []int{30, 22, 23, 10}}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCO2Rate(tt.args.bitLen, tt.args.nums); got != tt.want {
				t.Errorf("getCO2Rate() = %v, want %v", got, tt.want)
			}
		})
	}
}
