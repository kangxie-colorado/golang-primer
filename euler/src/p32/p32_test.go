/***

We shall say that an n-digit number is pandigital if it makes use of all the digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can be written as a 1 through 9 pandigital.

HINT: Some products can be obtained in more than one way so be sure to only include it once in your sum.

***/

// three buckets

package main

import (
	"testing"
)

func Test_hasDuplicatedNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"single digits", args{5}, false},
		{"two same digits", args{55}, true},
		{"", args{123}, false},
		{"", args{1123}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasDuplicatedNum(tt.args.num); got != tt.want {
				t.Errorf("hasDuplicatedNum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasZeroDigit(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"single digits", args{5}, false},
		{"two same digits", args{55}, false},
		{"", args{123}, false},
		{"", args{1123}, false},
		{"", args{11230}, true},
		{"", args{11023}, true},

		{"", args{1000}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasZeroDigit(tt.args.num); got != tt.want {
				t.Errorf("hasZeroDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPadigitProd(t *testing.T) {
	type args struct {
		prod int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1234}, false},
		{"", args{7254}, true},
		{"", args{1248}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPadigitProd(tt.args.prod); got != tt.want {
				t.Errorf("isPadigitProd() = %v, want %v", got, tt.want)
			}
		})
	}
}
