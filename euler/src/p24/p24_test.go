/**
A permutation is an ordered arrangement of objects. For example, 3124 is one possible permutation of the digits 1, 2, 3 and 4. If all of the permutations are listed numerically or alphabetically, we call it lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

012   021   102   120   201   210

What is the millionth lexicographic permutation of the digits 0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
**/

package main

import (
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	type args struct {
		perm []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 2}}, []int{0, 2, 1}},
		{"", args{[]int{0, 2, 1}}, []int{1, 0, 2}},
		{"", args{[]int{1, 0, 2}}, []int{1, 2, 0}},
		{"", args{[]int{1, 2, 0}}, []int{2, 0, 1}},
		{"", args{[]int{2, 0, 1}}, []int{2, 1, 0}},
		{"", args{[]int{2, 1, 0}}, nil},

		{"", args{[]int{0, 1, 2, 3}}, []int{0, 1, 3, 2}},
		{"", args{[]int{3, 2, 0, 1}}, []int{3, 2, 1, 0}},
		{"", args{[]int{0, 2, 3, 1}}, []int{0, 3, 1, 2}},
		{"", args{[]int{0, 3, 4, 2, 1}}, []int{0, 4, 1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextPermutation(tt.args.perm); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_nextNPermutaion(t *testing.T) {
	type args struct {
		perm []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 2}, 1}, []int{0, 2, 1}},
		{"", args{[]int{0, 1, 2}, 0}, []int{0, 1, 2}},
		{"", args{[]int{0, 1, 2}, 3}, []int{1, 2, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nextNPermutaion(tt.args.perm, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextNPermutaion() = %v, want %v", got, tt.want)
			}
		})
	}
}
