/***
The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.

Find the sum of the only eleven primes that are both truncatable from left to right and right to left.

NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.
***/
package main

import "testing"

func Test_leftTruncate(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{123}, 23},
		{"", args{12}, 2},
		{"", args{12321321}, 2321321},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := leftTruncate(tt.args.num); got != tt.want {
				t.Errorf("leftTruncate() = %v, want %v", got, tt.want)
			}
		})
	}
}
