/***

The prime 41, can be written as the sum of six consecutive primes:

41 = 2 + 3 + 5 + 7 + 11 + 13
This is the longest sum of consecutive primes that adds to a prime below one-hundred.

The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

Which prime, below one-million, can be written as the sum of the most consecutive primes?

*/

/**

analysis:
	this number has following properties
		1. it is a prime
		2. consecutive prime sum, then if sum - max(ele) will still be a prime and a consecutive prime sum

		there is some recursive going on here

	and the searching space is not super, so lets just populate all the primes under 1Million first, actually this can be a libs function

	okay, got that done in libs.PrimesUnderN
	then this becomes a reduce function

	I just need to reduce myself to 0 with 6 consecutive primes,
*/

package main

import (
	"reflect"
	"testing"
)

func Test_findConsecutivePrimes(t *testing.T) {
	type args struct {
		sum        int
		primeSpace []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{2, []int{2, 3, 5, 7}}, []int{}},
		{"", args{3, []int{2, 3, 5, 7}}, []int{}},

		{"", args{5, []int{2, 3, 5, 7}}, []int{2, 3}},
		{"", args{10, []int{2, 3, 5, 7}}, []int{2, 3, 5}},
		{"", args{17, []int{2, 3, 5, 7}}, []int{2, 3, 5, 7}},
		{"", args{28, []int{2, 3, 5, 7, 11}}, []int{2, 3, 5, 7, 11}},
		{"", args{41, []int{2, 3, 5, 7, 11, 13, 17, 19}}, []int{2, 3, 5, 7, 11, 13}},
		{"", args{43, []int{2, 3, 5, 7, 11, 13, 17, 19}}, []int{}},

		{"", args{11, []int{2, 3, 5, 7}}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findConsecutivePrimes(tt.args.sum, tt.args.primeSpace); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findConsecutivePrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
