package libs

import (
	"testing"
)

func TestIsNumPandigitToN(t *testing.T) {
	type args struct {
		num int
		n   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1, 1}, true},
		{"", args{1, 2}, false},
		{"", args{2, 1}, false},
		{"", args{1234, 3}, false},
		{"", args{1234, 4}, true},
		{"", args{1234, 5}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumPandigitToN(tt.args.num, tt.args.n); got != tt.want {
				t.Errorf("IsNumPandigitToN() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNumPandigitAtoB(t *testing.T) {
	type args struct {
		num int
		a   int
		b   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{1234, 1, 4}, true},
		{"", args{12304, 1, 4}, false},
		{"", args{12304, 0, 4}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNumPandigitAtoB(tt.args.num, tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsNumPandigitAtoB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameSet(t *testing.T) {
	exists := struct{}{}
	type intSet map[int]struct{}

	type args struct {
		s1 map[int]struct{}
		s2 map[int]struct{}
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"", args{intSet{1: exists, 2: exists, 3: exists}, intSet{1: exists, 2: exists, 3: exists}}, true},
		{"", args{intSet{1: exists, 2: exists, 3: exists}, intSet{2: exists, 3: exists}}, false},

		{"", args{intSet{1: exists, 2: exists, 3: exists}, nil}, false},

		{"", args{intSet{1: exists, 2: exists, 3: exists}, intSet{5: exists, 2: exists, 3: exists}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SameSet(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("SameSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
