// https://leetcode.com/problems/sequential-digits/

package main

import "testing"

func Test_getStep(t *testing.T) {
	type args struct {
		start int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{123}, 111},
		{"", args{1234}, 1111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStep(tt.args.start); got != tt.want {
				t.Errorf("getStep() = %v, want %v", got, tt.want)
			}
		})
	}
}
