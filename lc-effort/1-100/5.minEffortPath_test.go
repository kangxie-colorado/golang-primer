/*
https://leetcode.com/problems/path-with-minimum-effort/

*/

package main

import (
	"reflect"
	"testing"
)

func Test_minEffortOnAllPathsI(t *testing.T) {
	type args struct {
		start   point
		end     point
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"One Dimension Matrix", args{point{0, 0}, point{0, 2}, [][]int{{1, 2, 4}}}, 2},
		{"One Dimension Matrix", args{point{0, 0}, point{0, 2}, [][]int{{1, 100, -1}}}, 101},
		{"One Dimension Matrix", args{point{0, 0}, point{2, 0}, [][]int{{1}, {2}, {4}}}, 2},
		{"One Dimension Matrix", args{point{0, 0}, point{2, 0}, [][]int{{1}, {100}, {-1}}}, 101},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEffortOnAllPathsI(tt.args.start, tt.args.end, tt.args.heights); got != tt.want {
				t.Errorf("minEffortOnAllPathsI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_constructPath(t *testing.T) {
	type args struct {
		path       []pointLink
		currenNode pointLink
	}
	tests := []struct {
		name string
		args args
		want []pointLink
	}{
		// TODO: Add test cases.
		{"", args{[]pointLink{{point{-1, -1}, point{0, 0}}, {point{0, 0}, point{1, 0}}, {point{1, 0}, point{2, 0}},
			{point{2, 0}, point{2, 1}}, {point{2, 1}, point{2, 2}}}, pointLink{point{1, 0}, point{1, 1}}},
			[]pointLink{{point{-1, -1}, point{0, 0}}, {point{0, 0}, point{1, 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := constructPath(tt.args.path, tt.args.currenNode); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("constructPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minimumEffortPath(t *testing.T) {
	type args struct {
		heights [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"two dimension arrar", args{[][]int{{1, 2, 2}, {5, 8, 2}, {5, 3, 5}}}, 3},
		{"two dimension arrar", args{[][]int{{1, 2, 2}, {2, 8, 2}, {5, 3, 5}}}, 3},
		{"two dimension arrar", args{[][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}}, 2},

		{"two dimension arrar", args{[][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}}, 1},
		{"two dimension arrar", args{[][]int{{1, 2, 1, 1, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 2, 1, 2, 1}, {1, 1, 1, 2, 1}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumEffortPath(tt.args.heights); got != tt.want {
				t.Errorf("minimumEffortPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
