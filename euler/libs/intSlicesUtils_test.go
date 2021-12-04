package libs

import (
	"reflect"
	"testing"
)

func TestAddTwoIntSlices(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1+1", args{[]int{1}, []int{1}}, []int{2}},
		{"1+1", args{[]int{1}, []int{9}}, []int{0, 1}},

		{"", args{[]int{1, 2, 3}, []int{1}}, []int{2, 2, 3}},
		{"", args{[]int{1, 2, 3}, []int{}}, []int{1, 2, 3}},
		{"", args{[]int{1, 2, 3}, []int{9, 9, 9}}, []int{0, 2, 3, 1}},
		{"", args{[]int{9, 0, 9, 0, 1}, []int{9, 9, 9}}, []int{8, 0, 9, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddTwoIntSlices(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTwoIntSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutliplySliceByInt(t *testing.T) {
	type args struct {
		slice []int
		m     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3, 4, 5}, 1}, []int{1, 2, 3, 4, 5}},
		{"", args{[]int{1, 2, 3, 4, 5}, 2}, []int{2, 4, 6, 8, 0, 1}},
		{"", args{[]int{1, 5, 0, 9, 9}, 9}, []int{9, 5, 4, 1, 9, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutliplySliceByInt(tt.args.slice, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MutliplySliceByInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiplySliceBySlice(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{[]int{1, 2, 3}, []int{2}}, []int{2, 4, 6}},
		{"", args{[]int{1, 2, 3}, []int{0, 1}}, []int{0, 1, 2, 3}},
		{"", args{[]int{1, 2, 3}, []int{3, 2, 1}}, []int{3, 8, 4, 9, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MultiplySliceBySlice(tt.args.slice1, tt.args.slice2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MultiplySliceBySlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareTwoNumSlices(t *testing.T) {
	type args struct {
		slice1 []int
		slice2 []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 2, 3}, []int{0, 1, 1, 1}}, 1},
		{"", args{[]int{0, 1, 2, 3}, []int{0, 2, 1, 1}}, -1},
		{"", args{[]int{0, 1, 2, 3}, []int{0, 1, 2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTwoNumSlices(tt.args.slice1, tt.args.slice2); got != tt.want {
				t.Errorf("CompareTwoNumSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isNumInSlice(t *testing.T) {
	type args struct {
		num   int
		slice []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{2, []int{1, 2, 3}}, true},
		{"", args{2, []int{1, 1, 3}}, false},

		{"", args{9, []int{1, 1, 3, 5, 6, 7, 8}}, false},
		{"", args{7, []int{1, 1, 3, 5, 6, 7, 8}}, true},
		{"", args{0, []int{1, 1, 3, 5, 6, 7, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumInSlice(tt.args.num, tt.args.slice); got != tt.want {
				t.Errorf("IsNumInSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

type slice_int []int

func TestXPowY(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1^2", args{1, 2}, slice_int{1}},
		{"", args{2, 2}, slice_int{4}},
		{"", args{2, 3}, slice_int{8}},
		{"", args{2, 10}, slice_int{4, 2, 0, 1}},
		{"", args{3, 5}, slice_int{3, 4, 2}},
		{"", args{10, 5}, slice_int{0, 0, 0, 0, 0, 1}},
		{"", args{177, 3}, slice_int{3, 3, 2, 5, 4, 5, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := XPowY(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("XPowY() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToIntBase(t *testing.T) {
	type args struct {
		nums []int
		base int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{0, 1, 1}, 2}, 6},
		{"", args{[]int{0, 1, 1}, 10}, 110},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SliceToIntBase(tt.args.nums, tt.args.base); got != tt.want {
				t.Errorf("SliceToIntBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
