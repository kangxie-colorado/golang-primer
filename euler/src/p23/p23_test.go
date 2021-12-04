/**
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
**/

package main

import (
	"testing"
)

func Test_numCanBeSumOf2Abundant(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"<24", args{20}, false},
		{"<24", args{24}, true},
		{"", args{30}, true},
		{"", args{29}, false},
		{"", args{112}, true},
		{"", args{111}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numCanBeSumOf2Abundant(tt.args.num); got != tt.want {
				t.Errorf("numCanBeSumOf2Abundant() = %v, want %v", got, tt.want)
			}
		})
	}
}
