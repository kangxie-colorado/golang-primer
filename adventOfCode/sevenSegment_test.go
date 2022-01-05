package main

import "testing"

func Test_unqiNumberCounts(t *testing.T) {
	type args struct {
		outputs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"examples", args{[]string{"fdgacbe", "cefdb", "cefbgd", "gcbe"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := unqiNumberCounts(tt.args.outputs); got != tt.want {
				t.Errorf("appearanceOfUniqNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
