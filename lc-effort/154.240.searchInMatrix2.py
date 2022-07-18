"""
https://leetcode.com/problems/search-a-2d-matrix-ii/

so I came here from https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

I know binary search can be used
but there is some subtle differences... 

it is search for an existence so when found, it can short cut to return true
then if not existent, I will be search for a smallest number that is bigger than it 

hmm... I jumped too quick
it is not that same.. 

why don't I jump to compare to last element of this row
if it is smaller, then it will in the row range..
if it is bigger, it cannot be on this row

found row, then need to find the column...
huh... in the example
Input: matrix = [[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]], target = 5

5 is smaller than every row end, and every column end.. so this insight is useless
huh... deny also too fast

you can compare with left, 
if smaller than left/top, then it cannot be on that row/column
if bigger than right/bottom, then it cannot be on that row/column

so let me do this?
if in first row range... search for column... 
if in first col range... search for row...
"""


from re import T


class Solution(object):
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """

        """
        so I entered the wrong idea of search by range, becausing coming off the kth element problem
        then I think this can search by index...

        if m>target, then the index range to search betcome, [l.x,l.y] -> [m.x, m.y]
        otherwise, it becomes [0,m.y+1] -> [r.x, r.y] and [m.x+1, 0]->[r.x, r.y]

        so it forms a recursive structure 
        """
        M = len(matrix)
        N = len(matrix[0])

        def helper(startX, startY, endX, endY):
            if startX < 0 or endX >= M or startX > endX:
                return False

            if startY < 0 or endY >= N or startY > endY:
                return False

            midX = startX + (endX-startX)//2
            midY = startY + (endY-startY)//2

            if matrix[midX][midY] == target:
                return True

            if matrix[midX][midY] > target:
                return helper(startX, startY, midX-1, endY) or \
                    helper(midX, startY, endX, midY-1)
            else:
                return helper(startX, midY+1, midX, endY) or \
                    helper(midX+1, startY, endX, endY)

        return helper(0, 0, M-1, N-1)


"""
failed 
[[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]]
5

hmm... 
I see.. the hypothesis of reduce the searching range to  [l.x,l.y] -> [m.x, m.y] is not correct
need to do like the bigger case..

search two parts.. 
part1: [l.x,l.y] -> [m.x-1, r.y]
part2: [m.x, l.y] -> [r.x, m.y-1]

Runtime: 245 ms, faster than 23.87% of Python online submissions for Search a 2D Matrix II.
Memory Usage: 20 MB, less than 5.62% of Python online submissions for Search a 2D Matrix II.

Runtime: 139 ms, faster than 90.83% of Python online submissions for Search a 2D Matrix II.
Memory Usage: 20.1 MB, less than 5.62% of Python online submissions for Search a 2D Matrix II.

this solution of course is to eliminate 1/4 search space every time
there is another.. which eliminates one row/column at a time
https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/66140/My-concise-O(m%2Bn)-Java-solution
https://leetcode.com/problems/search-a-2d-matrix-ii/discuss/66139/C%2B%2B-search-from-top-right

I actually thought about it a bit but my sight is limited to think
only eliminate rows; or do columns.. 
I never thought I could eliminate row or col interleavely 

actually I almost touched it... but didn't pan out
so here the insights for search is it is usally about shrinking the space

shrinking by 1/4 is good; shrink by one row or one column is cool too

Runtime: 214 ms, faster than 43.98% of Python online submissions for Search a 2D Matrix II.
Memory Usage: 19.7 MB, less than 20.12% of Python online submissions for Search a 2D Matrix II.
"""


class Solution(object):
    def searchMatrix(self, matrix, target):

        M, N = len(matrix), len(matrix[0])
        row, col = 0, N-1
        while row < M and col >= 0:
            if matrix[row][col] == target:
                return True
            elif matrix[row][col] < target:
                # cannot be on this row because row end is too smaller
                row += 1
            else:
                # cannot be on this col because col start is too big
                col -= 1
        return False


if __name__ == '__main__':
    s = Solution()
    matrix = [[1, 2, 3, 4, 5], [6, 7, 8, 9, 10],
              [11, 12, 13, 14, 15], [16, 17, 18, 19, 20], [21, 22, 23, 24, 25]]

    target = 5
    print(s.searchMatrix(matrix, target))
