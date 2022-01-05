package main

import (
	"testing"
)

func Test_sonarSweep(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"the example", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 7},
		{"", args{[]int{201, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sonarSweep(tt.args.data); got != tt.want {
				t.Errorf("sonarSweep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sonarSweepWindow3(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"the example", args{[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 5},
		{"", args{[]int{201, 200, 208, 210, 200, 207, 240, 269, 260, 263}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sonarSweepWindow3(tt.args.data); got != tt.want {
				t.Errorf("sonarSweepWindow3() = %v, want %v", got, tt.want)
			}
		})
	}
}
