// https://leetcode.com/problems/robot-bounded-in-circle/

package main

/*
feels like just follow the instruction
but will it return to 0,0...

ah, so the difficult is how to tell the circl
is it returning to 0 considered a cricle, but if it never return to 0?? how do I know if it walked enough

apparently some quick termination
1. no L or R
2. L followed by R

also if follow condition is met, then there is definitely a cycle
1. sinlge L or R

Double "LL" is


aha... cycle...
fast/slow pointers...

if fast catch up slow then yes
if slow finished, fast didn't catch slow, no...

nah... "GL" will take 4 times...

okay... maybe G doesn't matter because when it turns, it will always offset back

so important is the L R and its count and distribution

L
LL: this can eliminate equal number of GG from both ends "GGLLGG" will be ""
LLL: this is bascially a R
LLLL: this is basically no ops...

LGGL -> LGGLLGGL -> LL
any leading L(R) can be removed, because it is just turn a direction and won't bother the final result

LG -> G (apparently not right)
but we can pad and remove  LG->LGLG->GLG (only messing more up)

LR is nothing
actually R can be replaced by LLL or vice versal

LGR: ...

hmm... okay I admin I have no idea how to make this
check hints, before I crush myself

- Calculate the final vector of how the robot travels after executing all instructions once - it consists of a change in position plus a change in direction.
- The robot stays in the circle if and only if (looking at the final vector) it changes direction (ie. doesn't stay pointing north), or it moves 0.


okay... I crush myself overthinking
so as long as it changes direction it will return
or it doesn't change direction but it return homes...

what I was doing was purely overthinking.... aha....

*/

func isRobotBounded(instructions string) bool {
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	pos := []int{0, 0}
	dir := 0

	for _, c := range instructions {
		switch c {
		case 'G':
			pos[0] = pos[0] + dirs[dir][0]
			pos[1] = pos[1] + dirs[dir][1]
		case 'L':
			dir = (dir - 1 + 4) % 4
		case 'R':
			dir = (dir + 1) % 4
		}
	}

	return dir != 0 || (pos[0] == 0 && pos[1] == 0)

}

/*
Runtime: 3 ms, faster than 28.07% of Go online submissions for Robot Bounded In Circle.
Memory Usage: 2 MB, less than 47.37% of Go online submissions for Robot Bounded In Circle.


yeah.. overthinking kills... don't!!!!
*/
