"""
https://leetcode.com/problems/largest-plus-sign/

I have thought a few directions but not super clear
I see a way that is to scan inwards

that requires I to build a matrix and actually manipulate that
not sure if that is even necessary 

but that is the way I see now
let me check
"""


from typing import List


class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        if len(mines) >= n*n:
            return 0

        A = [[1]*n for i in range(n)]
        for x, y in mines:
            A[x][y] = 0

        s, d = 1, n-2
        res = 1
        while s <= d:
            for x in range(s, d+1):
                for y in (s, d):
                    if A[x][y] == 1:
                        A[x][y] = min(A[x][y+1], A[x][y-1],
                                      A[x-1][y], A[x+1][y]) + 1
                        res = max(res, A[x][y])

            for y in range(s, d+1):
                for x in (s, d):
                    if A[x][y] == 1:
                        A[x][y] = min(A[x][y+1], A[x][y-1],
                                      A[x-1][y], A[x+1][y]) + 1
                        res = max(res, A[x][y])

            s, d = s+1, d-1

        return res


"""
3
[[0,0]]

failed here... 

okay.. logically wrong!!
when n>=8, it will not be right at all...

the loop 2 can at most be 2
but loop 3 can at most be 3... need to look 2 positons over??
optimized brute force???
maybe it can use the partial sum idea here...???

that is abs( A[i] - A[i+2]) == 2, then it can extend by 2, k can be 3...
yes... let me try this...
"""


class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        if len(mines) >= n*n:
            return 0

        A = [[1]*n for i in range(n)]
        sumH = [[1]*n for i in range(n)]
        sumV = [[1]*n for i in range(n)]
        for x, y in mines:
            A[x][y] = 0
            sumH[x][y] = 0
            sumV[x][y] = 0

        for i in range(n):
            prefixH = 0
            for j in range(n):
                if sumH[i][j]:
                    sumH[i][j] += prefixH
                    prefixH = sumH[i][j]
                else:
                    prefixH = 0

        for j in range(n):
            prefixV = 0
            for i in range(n):
                if sumH[i][j]:
                    sumV[i][j] += prefixV
                    prefixV = sumV[i][j]
                else:
                    prefixV = 0

        s, d = 1, n-2
        res = 1
        while s <= d:
            for x in range(s, d+1):
                for y in (s, d) if s < d else (s,):
                    if A[x][y] == 1:
                        trials = s
                        while trials > 0:
                            if sumV[x][y] - sumV[x-trials][y] == trials and \
                                    sumV[x+trials][y] - sumV[x][y] == trials and \
                                    sumH[x][y] - sumH[x][y-trials] == trials and \
                                    sumH[x][y+trials] - sumH[x][y] == trials:
                                A[x][y] += trials
                                break
                            trials -= 1
                        res = max(res, A[x][y])

            for y in range(s, d+1):
                for y in (s, d) if s < d else (s,):
                    if A[x][y] == 1:
                        trials = s
                        while trials > 0:
                            if sumV[x][y] - sumV[x-trials][y] == trials and \
                                    sumV[x+trials][y] - sumV[x][y] == trials and \
                                    sumH[x][y] - sumH[x][y-trials] == trials and \
                                    sumH[x][y+trials] - sumH[x][y] == trials:
                                A[x][y] += trials
                                break
                            trials -= 1
                        res = max(res, A[x][y])

            s, d = s+1, d-1

        return res


"""
if this is correct, I would be super surprised... 
but think I am doing this

I don't need to peel loop by loop. 
just scan all over..

nah.. still better to loop by loops..
but maybe I can start from inner loop

okay.. nasty/naive logical errors here
at the center, it was dealt multiple times... 

okay... I don't think this can do any better...
but one last try..
"""


class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        if len(mines) >= n*n:
            return 0

        A = [[1]*n for i in range(n)]
        sumH = [[1]*n for i in range(n)]
        sumV = [[1]*n for i in range(n)]
        for x, y in mines:
            A[x][y] = 0
            sumH[x][y] = -1
            sumV[x][y] = -1

        for i in range(n):
            prefixH = 0
            for j in range(n):
                if sumH[i][j] > 0:
                    sumH[i][j] += prefixH
                    prefixH = sumH[i][j]
                else:
                    prefixH = 0

        for j in range(n):
            prefixV = 0
            for i in range(n):
                if sumH[i][j] > 0:
                    sumV[i][j] += prefixV
                    prefixV = sumV[i][j]
                else:
                    prefixV = 0

        s = n//2
        d = n-1-s
        res = 1
        while s >= 0:
            pts = set()
            for x in range(s, d+1):
                for y in (s, d):
                    pts.add((x, y))
            for y in range(s, d+1):
                for x in (s, d):
                    pts.add((x, y))

            for x, y in pts:
                if A[x][y] == 1:
                    trials = s
                    while trials > 0:
                        if sumV[x][y] - sumV[x-trials][y] == trials and \
                                sumV[x+trials][y] - sumV[x][y] == trials and \
                                sumH[x][y] - sumH[x][y-trials] == trials and \
                                sumH[x][y+trials] - sumH[x][y] == trials:
                            A[x][y] += trials
                            break
                        trials -= 1
                    res = max(res, A[x][y])

            s -= 1

        return res


"""
https://leetcode.com/problems/largest-plus-sign/discuss/113314/JavaC%2B%2BPython-O(N2)-solution-using-only-one-grid-matrix

https://leetcode.com/problems/largest-plus-sign/discuss/1453636/Intuitiveor-Explained-with-image-or-Short-and-Clean-or-C%2B%2B

I spent time to crack a good algorithm but I think that effort is mis-placed
it basically has to count 4 directions 
but how to reduce the O(N^3) to O(N^2)

the first link has something to do that

1. Create an N-by-N matrix grid, with all elements initialized with value N.
2. Reset those elements to 0 whose positions are in the mines list.
3. For each position (i, j), find the maximum length of 1's in each of the four directions and set grid[i][j] to the minimum of these four lengths. Note that there is a simple recurrence relation relating the maximum length of 1's at current position with previous position for each of the four directions (labeled as l, r, u, d).
4. Loop through the grid matrix and choose the maximum element which will be the largest axis-aligned plus sign of 1's contained in the grid.

that initialized to N and decrease the value is first key piece
then this is too
            l = l + 1 if grid[i][j] != 0 else 0
            if l < grid[i][j]:
                grid[i][j] = l
            
            # looking to left, if I am not zero, then left-axis can be l+1
            # e,g, 1 1*, it can be 2
            # 2 1*, it can be 3... that 2 means there are 1 1 to the left already 
            # 1 0*, then it can only be 0
            # 0 1*, then it can be only be 1
    now walk thru the logical transition I admire the author more...
    there is some smart dealing here

woow...
    for i in range(N):
        l, r, u, d = 0, 0, 0, 0
            
        for j, k in zip(range(N), reversed(range(N))):
            l = l + 1 if grid[i][j] != 0 else 0
            if l < grid[i][j]:
                grid[i][j] = l
            
            r = r + 1 if grid[i][k] != 0 else 0
            if r < grid[i][k]:
                grid[i][k] = r

            u = u + 1 if grid[j][i] != 0 else 0
            if u < grid[j][i]:
                grid[j][i] = u
                
            d = d + 1 if grid[k][i] != 0 else 0
            if d < grid[k][i]:
                grid[k][i] = d
    notice it is actually using the idea of looping in..
    see the comment here https://leetcode.com/problems/largest-plus-sign/discuss/113314/JavaC++Python-O(N2)-solution-using-only-one-grid-matrix/114381
    yes.. this is same idea but way better organization than mine

    it didn't stop at the middle and cross over..
    it will then cover all the ground..

    wow!wow! I hope this help me to see some light
    I was fixated on meeting and stopping for such issues... here, it cross and one variable always looks for on direction
    when it cross over and reach each other's originating point, then it cover all grounds... 

    pure genius... 

and I want to code it
"""


class Solution:
    def orderOfLargestPlusSign(self, n: int, mines: List[List[int]]) -> int:
        A = [[n]*n for i in range(n)]
        for i, j in mines:
            A[i][j] = 0

        for i in range(n):
            l, r, u, d = [0]*4
            # this is actually a loop-by-loop procedure and it traverse inwards
            # l,r,u,d need to be kept and maintained
            for j, k in zip(range(n), reversed(range(n))):
                l = l+1 if A[i][j] != 0 else 0
                A[i][j] = min(A[i][j], l)

                r = r+1 if A[i][k] != 0 else 0
                A[i][k] = min(A[i][k], r)

                # because this is a n*n, so swap row/col we can walk the vertical columns
                u = u+1 if A[j][i] != 0 else 0
                A[j][i] = min(A[j][i], u)

                d = d+1 if A[k][i] != 0 else 0
                A[k][i] = min(A[k][i], d)

        return max([max(R) for R in A])


"""
Runtime: 7043 ms, faster than 33.53% of Python3 online submissions for Largest Plus Sign.
Memory Usage: 18.3 MB, less than 91.76% of Python3 online submissions for Largest Plus Sign.
"""
if __name__ == '__main__':
    s = Solution()
    print(s.orderOfLargestPlusSign(3, [[0, 0]]))
    print(s.orderOfLargestPlusSign(3, [[0, 2], [1, 0], [2, 0]]))
