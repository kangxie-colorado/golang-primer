package main

// https://leetcode.com/problems/adding-spaces-to-a-string/submissions/

func _1_addSpaces(s string, spaces []int) string {

	str := ""
	lastSpace := 0
	for _, space := range spaces {
		if space != 0 {
			str += s[:space-lastSpace] + " "
			s = s[space-lastSpace:]
			lastSpace = space
		} else {
			str += " "
		}

	}

	return str + s

}

/*
this O(n) failed here
https://leetcode.com/submissions/detail/710064148/testcase/

62 / 66 test cases passed.

is it O(n)?
how can there be better than O(n)

maybe my str is not pre-allocated.. so the append is expensive
let me pre-allocated it and see

*/

func _2_addSpaces(s string, spaces []int) string {

	str := make([]rune, len(s)+len(spaces))
	si := 0
	lastSpace := 0
	for _, space := range spaces {
		if space != 0 {
			// str += s[:space-lastSpace] + " "
			for _, b := range s[:space-lastSpace] {
				str[si] = b
				si++
			}

			str[si] = ' '
			si++

			s = s[space-lastSpace:]
			lastSpace = space
		} else {
			str[si] = ' '
			si++
		}
	}

	// append rest runes in s leftover
	for _, s := range s {
		str[si] = s
		si++
	}

	return string(str)

}

/*
Success
Details
Runtime: 210 ms, faster than 7.69% of Go online submissions for Adding Spaces to a String.
Memory Usage: 17.9 MB, less than 15.38% of Go online submissions for Adding Spaces to a String.
Next challenges:


change a bit
			// str += s[:space-lastSpace] + " "
			for _, b := range s[:space-lastSpace]  {
				str[si] = b
				si++
			}

            str[si] = ' '
			si++
reduce memory allocation
Runtime: 120 ms, faster than 76.92% of Go online submissions for Adding Spaces to a String.
Memory Usage: 16.6 MB, less than 23.08% of Go online submissions for Adding Spaces to a String.
Next challenges:

hmm... let me make it no need to move the slice pointer

*/

func addSpaces(s string, spaces []int) string {

	str := make([]byte, len(s)+len(spaces))
	si := 0    // writer pointer
	readi := 0 // read pointer
	lastSpace := 0
	for _, space := range spaces {
		if space != 0 {
			// str += s[:space-lastSpace] + " "

			for read := 0; read < space-lastSpace; read++ {
				str[si] = s[readi+read]
				si++
			}
			readi += space - lastSpace

			str[si] = ' '
			si++

			lastSpace = space
		} else {
			str[si] = ' '
			si++
		}
	}

	// append rest runes in s leftover
	for readi < len(s) {
		str[si] = s[readi]
		si++
		readi++
	}

	return string(str)

}

/*
Runtime: 144 ms, faster than 38.46% of Go online submissions for Adding Spaces to a String.
Memory Usage: 20.4 MB, less than 7.69% of Go online submissions for Adding Spaces to a String.
Next challenges:

hmm.. use byte instead of rune
Success
Details
Runtime: 104 ms, faster than 84.62% of Go online submissions for Adding Spaces to a String.
Memory Usage: 15 MB, less than 38.46% of Go online submissions for Adding Spaces to a String.
Next challenges:
*/
