/*
https://leetcode.com/problems/max-points-on-a-line/description/

I didn't pass with python
but I think 27M calculation should not TLE
so let me try go

python code here
    def maxPoints(self, points: List[List[int]]) -> int:
        res = 1

        for i in range(len(points)):
            for j in range(i+1, len(points)):
                p1, p2 = points[i], points[j]
                # p1,p2 is a line
                lineRes = 2
                x1, y1 = p1
                x2, y2 = p2
                for x, y in points:
                    if [x, y] != [x1, y1] and [x, y] != [x2, y2] and (y-y1)*(x-x2) == (y-y2) * (x-x1):
                        lineRes += 1

                res = max(res, lineRes)
        return res
*/

package main

func maxPoints(points [][]int) int {
	res := 1

	for i := 0; i < len(points); i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			p1, p2 := points[i], points[j]
			lineRes := 2
			x1, y1 := p1[0], p1[1]
			x2, y2 := p2[0], p2[1]

			for k := 0; k < len(points); k += 1 {
				if k == i || k == j {
					continue
				}

				x, y := points[k][0], points[k][1]
				if (y-y1)*(x-x2) == (y-y2)*(x-x1) {
					lineRes += 1
				}
			}
			res = max(res, lineRes)

		}
	}

	return res
}

/*
Runtime 39 ms
Beats 29.49%

okay.. at least this brute force is a pass
*/
