// https://leetcode.com/problems/valid-palindrome/

package main

import "fmt"

/*
this is a pretty easy problem
but my submission failed...

something I don't see and understand
let me debug
*/

func isPalindrome(s string) bool {
	toLower := func(c byte) byte {
		if c >= 'A' && c <= 'Z' {
			return c - 'A' + 'a'
		}

		return c
	}

	isChar := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
	}
	isNumber := func(c byte) bool {
		return c >= '0' && c <= '9'
	}

	for l, r := 0, len(s)-1; l < r; {
		for l < len(s) && !isChar(s[l]) && !isNumber(s[l]) {
			l++
		}

		for r >= 0 && !isChar(s[r]) && !isNumber(s[r]) {
			r--
		}

		if l < r && toLower(s[l]) != toLower(s[r]) {
			return false
		}

		l++
		r--
	}

	return true
}
func testIsPalindrome() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
}

/*
	for !isChar(s[l]) {
		l++
	}

	I used if... this is a mistake I make pretty often... yikes..
	if !isChar(s[l]) {
		l++
	}

okay. edge case
".,"

yeah I also now often forget checking edge cases
edge cases are score loser
        for l<len(s) && !isChar(s[l]) {
            l++
        }

        for r>=0 && !isChar(s[r]) {
            r--
        }

        if l<r&&toLower(s[l]) != toLower(s[r]) {
            return false
        }

then I failed at
"0P"

I assumed all lower english letters...
but it ain't

so this is an easy problem!
yes, the algorithm is not any conceptually hard but the edge cases kill

has to be relentless...
*/
