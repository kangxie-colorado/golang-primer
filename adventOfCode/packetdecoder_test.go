package main

import (
	"reflect"
	"testing"
)

func Test_getBinaryString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"12"}, "00010010"},
		{"", args{"A"}, "1010"},

		{"", args{"D2FE28"}, "110100101111111000101000"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBinaryString(tt.args.input); got != tt.want {
				t.Errorf("getBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAllVersions(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{"8A004A801A8002F478"}, 16},
		{"", args{"620080001611562C8802118E34"}, 12},
		{"", args{"C0015000016115A2E0802F182340"}, 23},
		{"", args{"A0016C880162017C3686B18A3D4780"}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAllVersions(tt.args.input); got != tt.want {
				t.Errorf("getAllVersions() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculatePacket(t *testing.T) {
	type args struct {
		packet string
	}
	tests := []struct {
		name string
		args args
		want []int64
	}{
		// TODO: Add test cases.
		{"", args{"C200B40A82"}, []int64{3}},
		{"", args{"04005AC33890"}, []int64{54}},
		{"", args{"880086C3E88112"}, []int64{7}},
		{"", args{"CE00C43D881120"}, []int64{9}},
		{"", args{"D8005AC2A8F0"}, []int64{1}},
		{"", args{"F600BC2D8F"}, []int64{0}},
		{"", args{"9C005AC2F8F0"}, []int64{0}},
		{"", args{"9C0141080250320F1802104A08"}, []int64{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculatePacket(tt.args.packet); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculatePacket() = %v, want %v", got, tt.want)
			}
		})
	}
}
