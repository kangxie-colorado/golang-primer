//The four adjacent digits in the 1000-digit number that have the greatest product are 9 × 9 × 8 × 9 = 5832.

/*
73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450
*/

//Find the thirteen adjacent digits in the 1000-digit number that have the greatest product. What is the value of this product?

package main

import "testing"

var numStream string = "73167176531330624919225119674426574742355349194934" +
	"96983520312774506326239578318016984801869478851843" +
	"85861560789112949495459501737958331952853208805511" +
	"12540698747158523863050715693290963295227443043557" +
	"66896648950445244523161731856403098711121722383113" +
	"62229893423380308135336276614282806444486645238749" +
	"30358907296290491560440772390713810515859307960866" +
	"70172427121883998797908792274921901699720888093776" +
	"65727333001053367881220235421809751254540594752243" +
	"52584907711670556013604839586446706324415722155397" +
	"53697817977846174064955149290862569321978468622482" +
	"83972241375657056057490261407972968652414535100474" +
	"82166370484403199890008895243450658541227588666881" +
	"16427171479924442928230863465674813919123162824586" +
	"17866458359124566529476545682848912883142607690042" +
	"24219022671055626321111109370544217506941658960408" +
	"07198403850962455444362981230987879927244284909188" +
	"84580156166097919133875499200524063689912560717606" +
	"05886116467109405077541002256983155200055935729725" +
	"71636269561882670428252483600823257530420752963450"

func Test_greatestProdOfNAdjacentDigits(t *testing.T) {
	type args struct {
		n         int
		numStream string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1 digit in stream", args{1, "2397086"}, 9},
		{"0s", args{1, "000"}, 0},
		{"2 digit in stream of 2", args{2, "23"}, 6},
		{"2 digit in stream of 3", args{2, "231"}, 6},
		{"3 digit in stream of 3", args{3, "231"}, 6},

		{"3 digit in stream of 3 with 0", args{3, "201"}, 0},

		{"3 digit in stream of 4", args{3, "2316"}, 18},

		{"2 digits expect 0", args{2, "002020200"}, 0},
		{"2 digit in stream of 2", args{2, "23"}, 6},
		{"2 digit in stream of 2", args{2, "232"}, 6},

		{"2 digit in stream", args{2, "2397086"}, 63},
		{"4 digit in stream", args{4, "1234"}, 24},
		{"4 digit in stream", args{4, "12345"}, 120},
		{"4 digit in stream", args{4, "123045"}, 0},
		{"4 digit in stream", args{4, "40319989000889"}, 5832},

		{"4 digit in stream", args{4, "82166370484403199890008895243450658541227588666881"}, 23514624000},

		{"4 in this stream", args{4, numStream}, 5832},
		{"13 in this stream", args{4, numStream}, 5832},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := greatestProdOfNAdjacentDigits(tt.args.n, tt.args.numStream); got != tt.want {
				t.Errorf("greatestProdOfNAdjacentDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}