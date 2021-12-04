package libs

import (
	"reflect"
	"testing"
)

func TestPrimesUnderN(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{2}, []int{}},
		{"", args{3}, []int{2}},
		{"", args{5}, []int{2, 3}},
		{"", args{42}, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimesUnderN(tt.args.N); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimesUnderN() = %v, want %v", got, tt.want)
			}
		})
	}
}
