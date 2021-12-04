/***
The decimal number, 585 = 10010010012 (binary), is palindromic in both bases.

Find the sum of all numbers, less than one million, which are palindromic in base 10 and base 2.

(Please note that the palindromic number, in either base, may not include leading zeros.)
***/

package main

import (
	"testing"
)

func Test_getReversedNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{1}, 1},
		{"", args{21}, 12},
		{"", args{210}, 12},
		{"", args{2103}, 3012},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getReversedNum(tt.args.num); got != tt.want {
				t.Errorf("getReversedNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRvereredBin(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{1}, 1},
		{"", args{2}, 1},
		{"", args{3}, 3},
		{"", args{585}, 585},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getRvereredBin(tt.args.num); got != tt.want {
				t.Errorf("getRvereredBin() = %v, want %v", got, tt.want)
			}
		})
	}
}
