package main

import (
	"reflect"
	"testing"
)

func Test_reduce_bigger_routes(t *testing.T) {
	type args struct {
		upper_row []int
		lower_row []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{nil, []int{2, 3}}, nil},
		{"", args{[]int{2, 3}, nil}, []int{2, 3}},
		{"", args{[]int{1}, []int{2, 3}}, []int{4}},
		{"", args{[]int{8, 5, 9, 3}, nil}, []int{8, 5, 9, 3}},
		{"", args{[]int{2, 4, 6}, []int{8, 5, 9, 3}}, []int{10, 13, 15}},
		{"", args{[]int{7, 4}, []int{10, 13, 15}}, []int{20, 19}},
		{"", args{[]int{3}, []int{20, 19}}, []int{23}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reduce_bigger_routes(tt.args.upper_row, tt.args.lower_row); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reduce_bigger_routes() = %v, want %v", got, tt.want)
			}
		})
	}
}
