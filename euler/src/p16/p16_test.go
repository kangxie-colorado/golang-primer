package main

import (
	"reflect"
	"testing"
)

func Test_int_as_array_multiply_by_2(t *testing.T) {
	type args struct {
		in0 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0}}, []int{0}},
		{"", args{[]int{1}}, []int{2}},
		{"", args{[]int{2}}, []int{4}},
		{"", args{[]int{3}}, []int{6}},
		{"", args{[]int{4}}, []int{8}},
		{"", args{[]int{5}}, []int{0, 1}},
		{"", args{[]int{6}}, []int{2, 1}},
		{"", args{[]int{7}}, []int{4, 1}},
		{"", args{[]int{8}}, []int{6, 1}},
		{"", args{[]int{9}}, []int{8, 1}},
		{"", args{[]int{0, 1}}, []int{0, 2}},
		{"", args{[]int{1, 1}}, []int{2, 2}},
		{"", args{[]int{2, 1}}, []int{4, 2}},
		{"", args{[]int{3, 1}}, []int{6, 2}},
		{"", args{[]int{4, 1}}, []int{8, 2}},
		{"", args{[]int{4, 8, 3, 6, 1}}, []int{8, 6, 7, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := int_as_array_multiply_by_2(tt.args.in0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("int_as_array_multiply_by_2() = %v, want %v", got, tt.want)
			}
		})
	}
}
