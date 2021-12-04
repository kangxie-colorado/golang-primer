package libs

import (
	"reflect"
	"testing"
)

func Test_isPrime(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1}, false},
		{"", args{2}, true},
		{"", args{5}, true},
		{"", args{25}, false},
		{"", args{97}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPrime(tt.args.num); got != tt.want {
				t.Errorf("IsPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasDuplicatedDigit(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasDuplicatedDigit(tt.args.num); got != tt.want {
				t.Errorf("HasDuplicatedDigit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnitFrac(t *testing.T) {
	type args struct {
		d      int
		maxLen int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"1/1", args{1, 10}, []int{1}},

		{"1/2", args{2, 10}, []int{0, 5}},
		{"1/20", args{20, 10}, []int{0, 0, 5}},
		{"1/200", args{200, 10}, []int{0, 0, 0, 5}},

		{"1/5", args{5, 10}, []int{0, 2}},
		{"1/5000", args{5000, 10}, []int{0, 0, 0, 0, 2}},

		{"1/4", args{4, 10}, []int{0, 2, 5}},
		{"1/40", args{40, 10}, []int{0, 0, 2, 5}},

		{"1/8", args{8, 10}, []int{0, 1, 2, 5}},
		{"1/16", args{16, 10}, []int{0, 0, 6, 2, 5}},

		{"1/3", args{3, 10}, []int{0, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}},
		{"1/6", args{6, 10}, []int{0, 1, 6, 6, 6, 6, 6, 6, 6, 6, 6}},
		{"1/7", args{7, 10}, []int{0, 1, 4, 2, 8, 5, 7, 1, 4, 2, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnitFrac(tt.args.d, tt.args.maxLen); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnitFrac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrimeFactors(t *testing.T) {
	type args struct {
		given int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{"", args{5}, []int{5}},
		{"", args{6}, []int{2, 3}},
		{"", args{12}, []int{2, 2, 3}},
		{"", args{16}, []int{2, 2, 2, 2}},
		{"", args{13195}, []int{5, 7, 13, 29}},
		{"", args{600851475143}, []int{71, 839, 1471, 6857}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrimeFactors(tt.args.given); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PrimeFactors() = %v, want %v", got, tt.want)
			}
		})
	}
}
