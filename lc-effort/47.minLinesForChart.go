// https://leetcode.com/problems/minimum-lines-to-represent-a-line-chart/

/*
analysis

I first thought this can be any randome connections, including jumping over the dots to form disjoint lines
but that should not be true?

then just need to track the slope change
I heard people talking about precision - then I think I can use fraction number to represent slop
*/

package main

import (
	"fmt"
	"sort"
)

type Slope struct {
	x int
	y int
}

func (s *Slope) equal(other *Slope) bool {
	return s.x*other.y == s.y*other.x
}

func minimumLines(stockPrices [][]int) int {
	lines := 0
	slope := Slope{0, 1} // unreal value to ensure there is always a change

	sort.Slice(stockPrices, func(i, j int) bool { return stockPrices[i][0] < stockPrices[j][0] })

	for p1, p2 := 0, 1; p2 < len(stockPrices); p1, p2 = p1+1, p2+1 {
		s := Slope{
			x: stockPrices[p2][0] - stockPrices[p1][0],
			y: stockPrices[p2][1] - stockPrices[p1][1],
		}

		if !s.equal(&slope) {
			lines++
			slope = s
		}
	}

	return lines
}

/*
Runtime: 253 ms, faster than 97.80% of Go online submissions for Minimum Lines to Represent a Line Chart.
Memory Usage: 18 MB, less than 70.33% of Go online submissions for Minimum Lines to Represent a Line Chart.

*/

func testMimimumLines() {
	//a := [][]int{{72, 98}, {62, 27}, {32, 7}, {71, 4}, {25, 19}, {91, 30}, {52, 73}, {10, 9}, {99, 71}, {47, 22}, {19, 30}, {80, 63}, {18, 15}, {48, 17}, {77, 16}, {46, 27}, {66, 87}, {55, 84}, {65, 38}, {30, 9}, {50, 42}, {100, 60}, {75, 73}, {98, 53}, {22, 80}, {41, 61}, {37, 47}, {95, 8}, {51, 81}, {78, 79}, {57, 95}}
	//fmt.Println(minimumLines(a))
	b := [][]int{{6, 5}, {57, 57}, {20, 13}, {30, 20}, {1, 27}, {50, 42}, {28, 57}, {58, 51}, {12, 55}, {68, 43}, {53, 24}, {19, 26}, {83, 39}, {96, 100}, {78, 5}, {32, 94}, {66, 59}, {47, 59}, {90, 76}, {69, 15}, {16, 78}, {3, 25}, {54, 74}, {48, 1}}
	fmt.Println(minimumLines(b))
}
