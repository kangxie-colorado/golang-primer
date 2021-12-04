// The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
// Find the sum of all the primes below two million.
package main

import "testing"

func Test_sumOfPrimsUnder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"below 3", args{3}, 2},
		{"below 4", args{4}, 5},
		{"below 10", args{10}, 17},
		{"below 18", args{18}, 17 + 11 + 13 + 17},
		{"below 30", args{30}, 17 + 11 + 13 + 17 + 19 + 23 + 29},
		{"below 2000000", args{2000000}, 142913828922},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfPrimsUnder(tt.args.n); got != tt.want {
				t.Errorf("sumOfPrimsUnder() = %v, want %v", got, tt.want)
			}
		})
	}
}
