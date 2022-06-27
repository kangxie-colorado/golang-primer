// https://leetcode.com/problems/smallest-integer-divisible-by-k/

/*
analysis
	apparently you cannot hold the number in any numeric form
	10003 -> 1423 or something, so not possible hold it by numeric

	so array

	[1] -> 1
	[1,1] -> 11
	[1,1,1] -> 111
	...

	which one is the high bit? make it easier...
	maybe not a difference.

	calculate div by arrray, I have done this before


*/

package main

import (
	"reflect"
	"testing"
)

func Test_divSliceByInt(t *testing.T) {
	type args struct {
		num []int
		d   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 1, 1}, 1}, []int{1, 1, 1}},
		{"", args{[]int{1, 1, 1}, 3}, []int{3, 7}},
		{"", args{[]int{1, 1, 1}, 37}, []int{3}},
		{"", args{[]int{1, 1, 1}, 2}, nil},

		{"", args{[]int{1, 1, 1, 1}, 11}, []int{1, 0, 1}},
		{"", args{[]int{1, 1, 1, 1}, 101}, []int{1, 1}},

		{"", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 9}, []int{1, 2, 3, 4, 5, 6, 7, 9}},
		{"", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 12345679}, []int{9}},
		{"", args{[]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 8888}, nil},

		{"", args{[]int{1, 1, 1}, 111}, []int{1}},
		{"", args{[]int{1, 1, 1}, 112}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divSliceByInt(tt.args.num, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("divSliceByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
