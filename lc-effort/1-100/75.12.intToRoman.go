// https://leetcode.com/problems/integer-to-roman/

package main

import "strings"

/*
so what I now think is do a hard map and just cheat it out
*/

func _1_intToRoman(num int) string {
	singles := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundres := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands := []string{"", "M", "MM", "MMM"}

	replaces := [][]string{singles, tens, hundres, thousands}
	replaceIdx := 0
	res := ""

	for num > 0 {
		d := num % 10
		res = replaces[replaceIdx][d] + res
		replaceIdx++
		num /= 10
	}

	return res
}

/*
Runtime: 14 ms, faster than 53.43% of Go online submissions for Integer to Roman.
Memory Usage: 3.2 MB, less than 92.70% of Go online submissions for Integer to Roman.

change to string builder
*/

func intToRoman(num int) string {
	singles := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundres := []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands := []string{"", "M", "MM", "MMM"}

	replaces := [][]string{singles, tens, hundres, thousands}
	replaceIdx := 0
	res := []string{}
	var sb strings.Builder

	for num > 0 {
		d := num % 10
		res = append(res, replaces[replaceIdx][d])
		replaceIdx++
		num /= 10
	}

	for i := len(res) - 1; i >= 0; i-- {
		sb.WriteString(res[i])
	}

	return sb.String()
}

/*
Runtime: 18 ms, faster than 32.85% of Go online submissions for Integer to Roman.
Memory Usage: 3.7 MB, less than 20.15% of Go online submissions for Integer to Roman.

not better

but of course, the universal solution is like this

var nums = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

var table = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func intToRoman(num int) string {
	roman := ""
	for num > 0 {
		for _, a := range nums {
			if num >= a {
				roman += table[a]
				num -= a
				break
			}
		}
	}
	return roman
}

it is super smart..
*/
