// https://leetcode.com/problems/bulls-and-cows/

package main

import "strconv"

/*
this seems quickly like a hashmap problem
*/

func getBullsCows(secPos []int, guessPos []int) (int, int) {
	m := make(map[int]struct{})
	for _, n := range secPos {
		m[n] = struct{}{}
	}

	bulls := 0
	for _, g := range guessPos {
		if _, found := m[g]; found {
			bulls++
		}
	}

	cows := min(len(secPos)-bulls, len(guessPos)-bulls)
	return bulls, cows
}

func getHint(secret string, guess string) string {
	secMap := make(map[rune][]int)
	for i, s := range secret {
		secMap[s] = append(secMap[s], i)
	}

	guessMap := make(map[rune][]int)
	for i, g := range guess {
		guessMap[g] = append(guessMap[g], i)
	}

	bulls := 0
	cows := 0
	for g, gPos := range guessMap {
		if sPos, found := secMap[g]; found {
			b, c := getBullsCows(sPos, gPos)
			bulls += b
			cows += c
		}
	}

	return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}

/*
Runtime: 3 ms, faster than 62.77% of Go online submissions for Bulls and Cows.
Memory Usage: 4.9 MB, less than 5.32% of Go online submissions for Bulls and Cows.

good enough for me
take a look at others' since this is a popular one

Runtime: 0 ms, faster than 100.00% of Go online submissions for Bulls and Cows.
Memory Usage: 4.9 MB, less than 5.32% of Go online submissions for Bulls and Cows.

and yes, there is always a more concise versison

public String getHint(String secret, String guess) {
    int bulls = 0;
    int cows = 0;
    int[] numbers = new int[10];
    for (int i = 0; i<secret.length(); i++) {
        int s = Character.getNumericValue(secret.charAt(i));
        int g = Character.getNumericValue(guess.charAt(i));
        if (s == g) bulls++;
        else {
            if (numbers[s] < 0) cows++;
            if (numbers[g] > 0) cows++;
            numbers[s] ++;
            numbers[g] --;
        }
    }
    return bulls + "A" + cows + "B";
}

yes, it is not hard to see this
just go thru both strings in tandom

if a char is exact the same.. then bull
otherwise, just keep an array of counters or a hashmap...

but it is still tricky to notice the order
when it prev appears in secret, then guess, then it means when the char appears by guess, the count must be >0
when it prev appears in guess, then secrect, then it means when the char appears by secrect, then count must be <0

each time, just offset it by 1 towards 0..
the exact case is taken care by first branch
*/
