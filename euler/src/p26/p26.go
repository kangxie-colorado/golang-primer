/***
A unit fraction contains 1 in the numerator. The decimal representation of the unit fractions with denominators 2 to 10 are given:

1/2	= 	0.5
1/3	= 	0.(3)
1/4	= 	0.25
1/5	= 	0.2
1/6	= 	0.1(6)
1/7	= 	0.(142857)
1/8	= 	0.125
1/9	= 	0.(1)
1/10	= 	0.1
Where 0.1(6) means 0.166666..., and has a 1-digit recurring cycle. It can be seen that 1/7 has a 6-digit recurring cycle.

Find the value of d < 1000 for which 1/d contains the longest recurring cycle in its decimal fraction part.
***/

package main

import (
	"fmt"
	"math/rand"

	"github.com/kangxie-colorado/golang-primer/euler/libs"
)

/**
Interesting
1/d 10/d 100/d : this push the decimal point to the right
so
1/3 = 0.(3) 10/3 = 3.(3) : can we draw the conclusion that the repeating lenght is 1 (only 3)
1/2 = 0.5 10/2 = 5 : hmm? becaue 10/2 is a int, so it is not a recurring cycle

1/9 = 0.(1) 10/9 = 1.(1) 100/9 = 11.(1) 1000/9 = 111.(1): so can say it is repating 1, but in program how to model this
1/7 = 0.(142857), 10/7 = 1.42857(142857)... 10000000/7=142857.142857: when before/after the decimal point, things look the same, we know

nice thoery but its hard than just do plain calculation

another observation
1/d or n/d, it will end up with the same repeating cycles but only with different start point
so to simplify this, only focus on 1/d, i.e. the unit fraction
**/

/** moved to libs
but keep them here for full context because of those comments
func numAndChange(n int, d int) (int, int) {
	return n / d, n % d
}

func unitFrac(d int, maxLen int) []int {
	// calculate the fractional part until it appears to be repeating or it can divided evenly(e.g. 1/2=0.5)
	// now use the elementary dividing calculation
	fracs := []int{}
	numerator := 1
	for numerator%d != 0 {
		if numerator < d {
			fracs = append(fracs, 0)
			numerator *= 10
		} else {
			num, change := numAndChange(numerator, d)
			fracs = append(fracs, num)
			numerator = change * 10
		}

		// always started with a 0, so the first it > maxLen it is just maxLen fractional digits
		if len(fracs) > maxLen {
			break
		}
	}

	// divided evenly
	if numerator%d == 0 {
		fracs = append(fracs, numerator/d)
	}

	return fracs
}

**/

/***
okay now we have a func to calculate the fraction to any digit
but how to tell the cycle length
by math we know there has to be some circle

the fractions will be
0. arbitrary-digits cycle1 cycle2 cycle3
how to find the cycle

slow pointer, fast pointer? this can prove if I have cycle
then how to find the cycle starting point

actually I only need to know the length of cycle
slow/fast point --
 slow walked by len(arb) + len(some-cycle-part)
 fast walked by len()..

Nah.. this is not a linked list, thinking this way is a dead end!

Hei... because this is unit fraction (1/d), the cycle seems to always start from the first digit?
It seems so
Nah.. 1/6 = 0.1(6)

***/
func testEq(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// this doesn't count the arbitrary number before cycles and is busted
// so hey, why don't we look from the behind, if the fraction is enough long, the end must be already in the cycle
// we know there is a property to 1/d, it at most has a recurring length of d-1
// and it is reasonable to expect in 4000 digits for anything 1/d(d<1000), it has at least 3 repeations right...
// and for 40000 digits, anything d<1000, it has at least 30-40 repeations
func cycleLenForUnitFraction(d int) int {
	// get a very long frac, 10000 digits
	long_fracs := libs.UnitFrac(d, 40000)
	if len(long_fracs) < 40001 {
		// divided evenly already, return 1
		return 1
	}

	// we need to look for the smallest recurring cycle
	candidate := d - 1
	for recurring_len := d - 1; recurring_len > 0; recurring_len-- {
		slice_this_len := long_fracs[len(long_fracs)-recurring_len:]
		next_slice_this_len := long_fracs[len(long_fracs)-recurring_len*2 : len(long_fracs)-recurring_len]
		third_slice_this_len := long_fracs[len(long_fracs)-recurring_len*3 : len(long_fracs)-recurring_len*2]

		if testEq(slice_this_len, next_slice_this_len) && testEq(slice_this_len, third_slice_this_len) {
			// yeah maybe a match
			// but, if it is (0110110112345)011011011 ... it could get 011 as well
			// so let us jump a random times(between 3-30) of this length, in this case, 3 * rand-times, see that 3 digits is still 011
			// jump 3 times and it could be wash the mole out

			rand_jump1 := rand.Intn(30-4) + 4 // +4 to rid of the already compared 3 recurrence; -4 because I want to at most get 30
			rand_jump2 := rand.Intn(30-4) + 4
			rand_jump3 := rand.Intn(30-4) + 4

			slice_rand1 := long_fracs[len(long_fracs)-recurring_len*rand_jump1 : len(long_fracs)-recurring_len*(rand_jump1-1)]
			slice_rand2 := long_fracs[len(long_fracs)-recurring_len*rand_jump2 : len(long_fracs)-recurring_len*(rand_jump2-1)]
			slice_rand3 := long_fracs[len(long_fracs)-recurring_len*rand_jump3 : len(long_fracs)-recurring_len*(rand_jump3-1)]

			if testEq(slice_this_len, slice_rand1) && testEq(slice_this_len, slice_rand2) && testEq(slice_this_len, slice_rand3) {
				if candidate > recurring_len {
					candidate = recurring_len
				}
			}

		}

	}
	return candidate
}

func main() {
	fmt.Println(libs.UnitFrac(4, 10))
	fmt.Println(libs.UnitFrac(11, 100))

	fmt.Println(libs.UnitFrac(19, 100))
	max_cycle_len := 1
	for d := 1; d < 1000; d++ {
		cycle_len := cycleLenForUnitFraction(d)
		if max_cycle_len < cycle_len {
			max_cycle_len = cycle_len
			fmt.Println(max_cycle_len, d)
		}
	}
}
