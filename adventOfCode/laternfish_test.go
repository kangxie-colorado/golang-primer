package main

import (
	"testing"
)

func Test_fishNumAfterDays(t *testing.T) {
	type args struct {
		fishes []int
		days   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"the example", args{[]int{3, 4, 3, 1, 2}, 2}, 6},
		{"the example", args{[]int{3, 4, 3, 1, 2}, 18}, 26},
		{"the example", args{[]int{3, 4, 3, 1, 2}, 80}, 5934},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fishNumAfterDays(tt.args.fishes, tt.args.days); got != tt.want {
				t.Errorf("fishNumAfterDays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fishNumAfterDaysParallel(t *testing.T) {
	type args struct {
		fishes []int
		days   int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"the example", args{[]int{3, 4, 3, 1, 2}, 2}, 6},
		{"the example", args{[]int{3, 4, 3, 1, 2}, 18}, 26},
		{"the example", args{[]int{3, 4, 3, 1, 2}, 80}, 5934},
		//{"the example", args{[]int{3, 4, 3, 1, 2}, 256}, 26984457539},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fishNumAfterDaysParallel(tt.args.fishes, tt.args.days); got != tt.want {
				t.Errorf("fishNumAfterDaysParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}
