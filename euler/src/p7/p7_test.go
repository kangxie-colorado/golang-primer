package main

import (
	"testing"
)

func Test_getNthPrime(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"1st", args{1}, 2},
		{"6th", args{6}, 13},
		{"7th", args{7}, 17},
		{"8th", args{8}, 19},
		{"9th", args{9}, 23},
		{"10th", args{10}, 29},
		{"11th", args{11}, 31},
		{"12th", args{12}, 37},

		{"100th", args{100}, 541},
		{"10001th", args{10001}, 104743},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNthPrime(tt.args.n); got != tt.want {
				t.Errorf("getNthPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	type args struct {
		num int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"4", args{4}, false},
		{"8", args{8}, false},
		{"5", args{5}, true},
		{"17", args{17}, true},
		{"1001", args{1001}, false},
		{"10001", args{10001}, false},
		{"1000001", args{1000001}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPrime(tt.args.num); got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
