// https://leetcode.com/problems/validate-ip-address/

package main

import (
	"strconv"
	"strings"
)

/*
analysis
	why is this hard?  Accepted 126,893    Submissions 483,235  --- so low accepted rate

	so the rules for IPv4
	1. 4 parts delimited by .
	2. each part can be converted into a number [0,255]
		- so can use a exception based
	3. no leading zeros...


*/
func ip4Rule4Parts(ip string) bool {
	return strings.Count(ip, ".") == 3
}

func ip4RuleValidSegment(ip string) bool {
	parts := strings.Split(ip, ".")

	for _, p := range parts {
		if len(p) == 0 || (len(p) > 1 && p[0] == '0') || p[0] < '0' || p[0] > '9' {
			// leading char is not 1-9, cannot be right; this include the leading zeros case
			return false
		}

		if intVar, err := strconv.Atoi(p); err != nil || intVar > 255 {
			return false
		}
	}

	return true

}

func ip6Rule8Parts(ip string) bool {
	return strings.Count(ip, ":") == 7
}

func ip6RuleValidSegment(ip string) bool {
	parts := strings.Split(ip, ":")

	for _, p := range parts {
		if len(p) < 1 || len(p) > 4 {
			return false
		}

		for _, c := range strings.ToUpper(p) {
			if c < '0' || c > 'F' {
				return false
			}
		}
	}

	return true
}

func validIPAddress(queryIP string) string {

	if strings.Contains(queryIP, ".") && !strings.Contains(queryIP, ":") {
		// possible IPv4
		if ip4Rule4Parts(queryIP) && ip4RuleValidSegment(queryIP) {
			return "IPv4"
		}

	} else if strings.Contains(queryIP, ":") && !strings.Contains(queryIP, ".") {
		// possible IPv6
		if ip6Rule8Parts(queryIP) && ip6RuleValidSegment(queryIP) {
			return "IPv6"
		}
	}

	return "Neither"
}

/*

8 / 73 test cases passed.
Status: Runtime Error
Submitted: 0 minutes ago
Runtime Error Message:
panic: runtime error: index out of range [0] with length 0

input: "1.0.1."

okay.. all kinds of strangities
*/

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Validate IP Address.
Memory Usage: 2 MB, less than 59.57% of Go online submissions for Validate IP Address.
Next challenges:

well where is the challendhe
*/
