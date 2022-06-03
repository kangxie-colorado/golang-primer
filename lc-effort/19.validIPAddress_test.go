// https://leetcode.com/problems/validate-ip-address/

package main

import "testing"

func Test_validIPAddress(t *testing.T) {
	type args struct {
		queryIP string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"", args{"172.16.254.1"}, "IPv4"},
		{"", args{"192.168.1.1"}, "IPv4"},
		{"", args{"192.168.0.1"}, "IPv4"},
		{"", args{"0.0.0.0"}, "IPv4"},

		{"", args{"2001:0db8:85a3:0000:0000:8a2e:0370:7334"}, "IPv6"},
		{"", args{"2001:0db8:85a3:0:0:8A2E:0370:7334"}, "IPv6"},

		{"", args{"192.168@1.1"}, "Neither"},
		{"", args{"192.168.01.1"}, "Neither"},
		{"", args{"256.256.256.256"}, "Neither"},

		{"", args{"2001:0db8:85a3::8A2E:037j:7334"}, "Neither"},
		{"", args{"02001:0db8:85a3:0000:0000:8a2e:0370:7334"}, "Neither"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validIPAddress(tt.args.queryIP); got != tt.want {
				t.Errorf("validIPAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
