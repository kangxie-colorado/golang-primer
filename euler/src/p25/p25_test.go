/***
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?
***/

package main

import (
	"reflect"
	"testing"
)

func Test_fibonacciNth(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"3", args{3}, []int{2}},
		{"5", args{5}, []int{5}},
		{"12", args{12}, []int{4, 4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciNth(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacciNth() = %v, want %v", got, tt.want)
			}
		})
	}
}
