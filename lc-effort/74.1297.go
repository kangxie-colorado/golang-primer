// https://leetcode.com/problems/maximum-number-of-occurrences-of-a-substring/

package main

import "fmt"

/*
not much idea except just sliding thru

so open up at minSize, sliding up to maxSize
then sliding down to minSize.. then sliding up

and maintain the number of diff chars
and use hashmap to keep the counts of each eligible substrs
*/

func _messy_maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	substrMap := make(map[string]int)
	charMap := make(map[byte]int)
	maxFreq := 0

	i, j := 0, 0
	for j < minSize {
		charMap[s[j]]++
		j++
	}
	j -= 1

	for j < len(s) {

		for {
			if len(charMap) <= maxLetters {
				substrMap[s[i:j+1]]++
				maxFreq = max(maxFreq, substrMap[s[i:j+1]])
			}
			if j-i+1 < maxSize {
				// can go to ==, and let that break
				j++
				charMap[s[j]]++

			} else {
				break
			}
		}

		for {
			if j-i+1 > minSize {
				charMap[s[i]]--
				if charMap[s[i]] == 0 {
					delete(charMap, s[i])
				}
				i++

				if len(charMap) <= maxLetters {
					substrMap[s[i:j+1]]++
					maxFreq = max(maxFreq, substrMap[s[i:j+1]])
				}
			} else {
				break
			}
		}

		j++
		charMap[s[j]]++
	}

	return maxFreq
}

/*
"aabcabcab"
2
2
3

okay... this expanding, and shrinking is so hard to get right in one loop
let me just do the minsize,minsize+1...maxsize iterations

actually the complexity is the same..

*/

func _2_maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	substrMap := make(map[string]int)
	charMap := make(map[byte]int)
	maxFreq := 0

	windowSize := minSize

	for windowSize <= maxSize {
		i, j := 0, 0
		for j < windowSize {
			charMap[s[j]]++
			j++
		}

		for j <= len(s) {
			if len(charMap) <= maxLetters {
				substrMap[s[i:j]]++
				maxFreq = max(maxFreq, substrMap[s[i:j]])
			}
			charMap[s[i]]--
			if charMap[s[i]] == 0 {
				delete(charMap, s[i])
			}
			i++
			if j < len(s) {
				charMap[s[j]]++
			}
			j++
		}

		windowSize++

	}

	return maxFreq
}

/*
Runtime: 480 ms, faster than 54.55% of Go online submissions for Maximum Number of Occurrences of a Substring.
Memory Usage: 18.1 MB, less than 72.73% of Go online submissions for Maximum Number of Occurrences of a Substring.

let me try improving the speed a bit?

I cannot think of any better algorithm
let me replace the map with A[26]
*/

func howManyChars(charMap []int) int {
	cnt := 0
	for _, n := range charMap {
		if n != 0 {
			cnt++

		}
	}

	return cnt
}

func tooManyChars(charMap []int, maxChars int) bool {
	cnt := 0
	for _, n := range charMap {
		if n != 0 {
			cnt++
			if cnt > maxChars {
				return false
			}
		}
	}

	return true
}

func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	substrMap := make(map[string]int)
	charMap := make([]int, 26)
	maxFreq := 0

	windowSize := minSize

	// haha.. minSize is the biggest possible..
	// I was blinded
	//for windowSize <= maxSize {
	i, j := 0, 0
	for j < windowSize {
		charMap[s[j]-'a']++
		j++
	}

	for j <= len(s) {
		//if howManyChars(charMap) <= maxLetters {
		if tooManyChars(charMap, maxLetters) {
			substrMap[s[i:j]]++
			maxFreq = max(maxFreq, substrMap[s[i:j]])
		}
		charMap[s[i]-'a']--
		i++
		if j < len(s) {
			charMap[s[j]-'a']++
		}
		j++
	}

	//	windowSize++

	//}

	return maxFreq
}

/*
Runtime: 209 ms, faster than 72.73% of Go online submissions for Maximum Number of Occurrences of a Substring.
Memory Usage: 18 MB, less than 72.73% of Go online submissions for Maximum Number of Occurrences of a Substring.

yep, golang really scans much faster on sequential memory esp when it is small
than dealing with map


Runtime: 189 ms, faster than 72.73% of Go online submissions for Maximum Number of Occurrences of a Substring.
Memory Usage: 19.2 MB, less than 54.55% of Go online submissions for Maximum Number of Occurrences of a Substring.

changing howManyChars() to tooManyChars(), saved a bit more time..


so what is the best solution
aha...
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {

    if maxLetters >26 {
        return 0
    }
    chars := []rune(s)
    checkLetter := make([]int, 26)
    st := 0
    exist := map[string]int{}
    cur := []rune{}

    for j:= 0 ; j<len(chars); j++{
        checkLetter[chars[j]-'a']++
        cur = append(cur, chars[j])
        for len(cur) > minSize{
            checkLetter[chars[st]-'a']--
            cur = cur[1:]
            st++
        }
        if len(cur) == minSize && valid(checkLetter, maxLetters){
            exist[s[st:j+1]]++
        }
    }

    max := 0

    for _ , val := range exist{
        if val > max {
            max = val
        }
    }

    return max
}

func valid(cL []int, maxLetters int) bool{

    for _, val := range cL {
        if val != 0{
            maxLetters--
        }
        if maxLetters < 0 {
            return false
        }
    }

    return true
}

maxSize is actually a distraction?
what the fuck..

yeah.. when you can match 4 letters in a row
you must be able to match 3..

so I removed the min->max size logic and run it on minSize only
Runtime: 31 ms, faster than 90.91% of Go online submissions for Maximum Number of Occurrences of a Substring.
Memory Usage: 6.5 MB, less than 90.91% of Go online submissions for Maximum Number of Occurrences of a Substring.

hahaha.. such a trick
*/

func testMaxFreq() {

	fmt.Println(maxFreq("aabcabcab", 2, 2, 3))
}
