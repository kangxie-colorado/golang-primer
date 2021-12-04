/****
By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.

By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the first member of this family, is the smallest prime with this property.

Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an eight prime value family.

***/

/**
	analysis:
		at first glance, very hard to get started
		then notice this number should have following properties
		1. it is a prime
		2. it should have at least two digits the same? no necessarily

		then what?
		I have the prime space, not sure if it is enough

		then notice
		there are at most ten replacements

		and you can replace any digits (any<=len(num)): at least 1 replaced, otherwise nonsense
		at most, all replaced but it won't ever be qualify for the most prime family, so just a benigh outlier; and we can disregard it totally


		n digits, replace[1:n], kind of can use the bitmask method and mapping it to the 10base world

		2 digist, [1:2): 0b1
		3 digits: [1:3): [0b1, 0b10]
		5 digits: [1:5): [0b1, 0b10, 0b11, 0b100]

		nope.. wrong direction, 5 digits we can replace any 1 digit; or any 2 digits
		for any 1 digit,

		this is the combination function.. but I am implementing it using bit mask... cool


		for any 1 digit in 5, 00000, left shit 1 5 times.
		for any 2 digit in 5, 00000, left shit 1 5 times, then for each 1, shift 2nd 1 a total of 4 times(5 times, but when it colide, skip)
		then this can build up...


		then we need a mapping...


		===
		ha, when I think about the combinations of any digits
		then 5 digits space
	 	00000 - 11111
		if there zero 1s, it is family 0
		if there is one 1, it is family 1
		...

		after all, the full combination is 2^(n+1)?
		2^6 = 64 =? 1(all 0) + 1(all 1) + 20(two 1) + 20(three 1, two 0) + 5(four 1) + 5(four 0, one 1) = 52?


		ugh... now pure combination because 11000 and 00011 are different but in combination term this is the same
		not permutation either, because 5 1s is only one scenario.. I forgot the exact math term but any
		the space is this many

		okay I am confused
		but to turn the whole space 00000->1111 into families, is not a problem

		ah, right, c(5,4) = 5*4/2 = 10, so above becomes
		2^5 = 32 = 1(all 0) + 1(all 1) + 10(two 1) + 10(three 1, two 0) + 5(four 1) + 5(four 0, one 1) = 32

		now the theory is clean, lets go

**/

package main

import (
	"reflect"
	"testing"
)

func Test_getCombFamilies(t *testing.T) {
	type args struct {
		digits int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{"", args{2}, [][]int{{0}, {1, 2}, {3}}},
		{"", args{3}, [][]int{{0}, {1, 2, 4}, {3, 5, 6}, {7}}},
		{"", args{5}, [][]int{{0}, {1, 2, 4, 8, 16}, {3, 5, 6, 9, 10, 12, 17, 18, 20, 24}, {7, 11, 13, 14, 19, 21, 22, 25, 26, 28}, {15, 23, 27, 29, 30}, {31}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCombFamilies(tt.args.digits); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCombFamilies() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_intToSlice(t *testing.T) {
	type args struct {
		num  int
		base int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{6, 2}, []int{0, 1, 1}},
		{"", args{8, 2}, []int{0, 0, 0, 1}},
		{"", args{8, 10}, []int{8}},
		{"", args{5678, 10}, []int{8, 7, 6, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToSlice(tt.args.num, tt.args.base); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyComb(t *testing.T) {
	type args struct {
		num  int
		comb int
	}

	tests := []struct {
		name string
		args args
		want Series
	}{
		// TODO: Add test cases.
		{"", args{13, 2}, Series{13, 10}},

		{"", args{4567, 6}, Series{4007, 110}},
		{"", args{4567, 8}, Series{4567, 1000}},
		{"", args{4567, 16}, Series{4567, 0}},
		{"", args{4567, 0}, Series{4567, 0}},
		{"", args{45678, 5}, Series{45070, 101}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := applyComb(tt.args.num, tt.args.comb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyComb() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_seriesToPrimesSlice(t *testing.T) {
	type args struct {
		s Series
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"13", args{Series{13, 10}}, []int{13, 23, 43, 53, 73, 83}},
		{"13", args{Series{13, 0}}, []int{13}},

		{"56003", args{Series{56003, 110}}, []int{56003, 56113, 56333, 56443, 56663, 56773, 56993}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seriesToPrimesSlice(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("seriesToPrimesSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_longestPrimeFamily(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{13}, []int{13, 23, 43, 53, 73, 83}},
		{"", args{56453}, []int{56003, 56113, 56333, 56443, 56663, 56773, 56993}},
		{"", args{121313}, []int{121313, 222323, 323333, 424343, 525353, 626363, 828383, 929393}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPrimeFamily(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestPrimeFamily() = %v, want %v", got, tt.want)
			}
		})
	}
}
