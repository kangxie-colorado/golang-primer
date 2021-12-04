// 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package main

import (
	"testing"
)

func create_array(start int, end_inclusive int) []int {
	var ret []int
	for i := start; i <= end_inclusive; i++ {
		ret = append(ret, i)
	}

	return ret
}

func Test_getSmallestLCMOfArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[1]", args{nums: []int{1}}, 1},
		{"[1,2]", args{nums: []int{1, 2}}, 2},
		{"[1,2,3]", args{nums: []int{1, 2, 3}}, 6},
		{"[1,2,3,4]", args{nums: []int{1, 2, 3, 4}}, 12},
		{"[1,2,3,4,5]", args{nums: []int{1, 2, 3, 4, 5}}, 60},
		{"[1:10]", args{nums: create_array(1, 10)}, 2520},
		{"[1:20]", args{nums: create_array(1, 20)}, 232792560},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLCMOfArray(tt.args.nums); got != tt.want {
				t.Errorf("getSmallestLCMOfArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLCM(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[2,3]", args{2, 3}, 6},

		{"[2,4]", args{2, 4}, 4},
		{"[2,7]", args{2, 7}, 14},
		{"[4,6]", args{4, 6}, 12},
		{"[48,64]", args{48, 64}, 192},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLCM(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("getLCM() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGCD(t *testing.T) {
	type args struct {
		num1 int
		num2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"[1,1]", args{1, 1}, 1},
		{"[1,2]", args{1, 2}, 1},

		{"[2,3]", args{2, 3}, 1},
		{"[2,4]", args{2, 4}, 2},
		{"[3,4]", args{3, 4}, 1},
		{"[4,6]", args{4, 6}, 2},
		{"[4,64]", args{4, 64}, 4},
		{"[48,64]", args{48, 64}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGCD(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("getGCD() = %v, want %v", got, tt.want)
			}
		})
	}
}
