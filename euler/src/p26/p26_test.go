/***
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
***/

package main

import (
	"testing"
)

func Test_cycleLenForUnitFraction(t *testing.T) {
	type args struct {
		d int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{2}, 1},
		{"", args{2000}, 1},
		{"", args{3}, 1},
		{"", args{7}, 6},
		{"", args{13}, 6},
		{"", args{17}, 16},
		{"", args{19}, 18},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cycleLenForUnitFraction(tt.args.d); got != tt.want {
				t.Errorf("cycleLenForUnitFraction() = %v, want %v", got, tt.want)
			}
		})
	}
}
