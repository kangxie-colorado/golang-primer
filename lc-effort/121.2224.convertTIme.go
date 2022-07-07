// https://leetcode.com/problems/minimum-number-of-operations-to-convert-time/

package main

import (
	"strconv"
	"strings"
)

/*
I thought the coin change... then I think this might just be a greedy
notice 1/5/15/60 are dividable from low to high... so it is simpler than coin change

let me try
*/

func convertTime(current string, correct string) int {
	hh_mm1 := strings.Split(current, ":")
	hh_mm2 := strings.Split(correct, ":")

	hh2, _ := strconv.Atoi(hh_mm2[0])
	hh1, _ := strconv.Atoi(hh_mm1[0])
	mm2, _ := strconv.Atoi(hh_mm2[1])
	mm1, _ := strconv.Atoi(hh_mm1[1])

	diff := (hh2-hh1)*60 + mm2 - mm1

	n := diff / 60
	diff %= 60
	n += diff / 15
	diff %= 15
	n += diff / 5
	diff %= 5
	n += diff

	return n
}
