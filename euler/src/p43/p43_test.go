/****

The number, 1406357289, is a 0 to 9 pandigital number because it is made up of each of the digits 0 to 9 in some order, but it also has a rather interesting sub-string divisibility property.

Let d1 be the 1st digit, d2 be the 2nd digit, and so on. In this way, we note the following:

d2d3d4=406 is divisible by 2
d3d4d5=063 is divisible by 3
d4d5d6=635 is divisible by 5
d5d6d7=357 is divisible by 7
d6d7d8=572 is divisible by 11
d7d8d9=728 is divisible by 13
d8d9d10=289 is divisible by 17
Find the sum of all 0 to 9 pandigital numbers with this property.

***/

package main

import "testing"

func Test_sliceNumber(t *testing.T) {
	type args struct {
		num int
		s   int
		e   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{123456, 1, 3}, 123},
		{"", args{123456, 2, 5}, 2345},
		{"", args{123456, 2, 5}, 2345},

		{"", args{123456, 6, 7}, -1},
		{"", args{123456, 7, 7}, -1},
		{"", args{123456, 1, 7}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceNumber(tt.args.num, tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("sliceNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
