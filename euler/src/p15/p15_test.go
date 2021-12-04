package main

import (
	"testing"
)

func Test_routes_tl_to_br(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"nan", args{-1, -1}, 0},
		{"dot", args{0, 0}, 1},
		{"h-line", args{0, 2}, 1},
		{"v-line", args{2, 0}, 1},
		{"1x1", args{1, 1}, 2},
		{"2x2", args{2, 2}, 6},
		{"20x20", args{20, 20}, 137846528820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := routes_tl_to_br(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("routes_tl_to_br() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_routes_tl_to_br_loop(t *testing.T) {
	type args struct {
		row int
		col int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"nan", args{-1, -1}, 0},
		{"dot", args{0, 0}, 1},
		{"h-line", args{0, 2}, 1},
		{"v-line", args{2, 0}, 1},
		{"1x1", args{1, 1}, 2},
		{"2x2", args{2, 2}, 6},
		{"20x20", args{20, 20}, 137846528820},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := routes_tl_to_br_loop(tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("routes_tl_to_br_loop() = %v, want %v", got, tt.want)
			}
		})
	}
}
