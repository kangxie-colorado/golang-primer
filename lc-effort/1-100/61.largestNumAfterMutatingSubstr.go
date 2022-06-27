// https://leetcode.com/problems/largest-number-after-mutating-substring/

/*
analysis
	honestly, I don't know what trick this problem is playing
	as it appears, I just need to replace the consecutive first found numbers that can make it bigger...

	what is the trick... let me find out that
*/

package main

func maximumNumber(num string, change []int) string {
	bytes := []byte(num)

	replacing := false
	for i := range bytes {
		d := int(bytes[i] - '0')
		if change[d] > d {
			bytes[i] = '0' + byte(change[d])
			replacing = true
		}

		if replacing && change[d] < int(bytes[d]-'0') {
			break
		}
	}

	return string(bytes)
}
