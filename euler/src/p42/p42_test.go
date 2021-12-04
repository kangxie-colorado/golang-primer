/***

The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?

***/

package main

import "testing"

func Test_isTriangleNum(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1}, true},
		{"", args{3}, true},
		{"", args{6}, true},
		{"", args{28}, true},
		{"", args{55}, true},
		{"", args{100}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTriangleNum(tt.args.num); got != tt.want {
				t.Errorf("isTriangleNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
