"""
https://leetcode.com/problems/search-a-2d-matrix/

this obviously should a modified binary search
if you flatten the matrix to 1-D, then this is to trivial

let me do that just for fun
"""


from ast import List
from math import fabs


class Solution1:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        A = []
        for r in matrix:
            A.extend(r)

        def bSearch():
            l, r = 0, len(A)-1
            while l < r:
                m = l+(r-l)//2
                if A[m] > target:
                    r = m
                else:
                    l = m+1

            return A[l] == target

        return bSearch()


"""
[[1]]
1

edge case...
so instead of doing three branches
    if A[m] == target:
        return True
    elif A[m] > target:
        r = m
    else:
        l = m+1
just do two but then you will at the merce of other bugs

pay attention if it should be contracting the right or contracting the left
if my conditions are like this 
if A[m] > target:
    r = m-1
else:
    l = m

then it means when A[m] == target, I will be moving l... so I should never skip m
thus l=m; hence r should m-1
and because of that I should calculate m this way: m = r-(r-l)//2

otherwise it could be dead loop

Runtime: 70 ms, faster than 46.84% of Python3 online submissions for Search a 2D Matrix.
Memory Usage: 14.4 MB, less than 42.57% of Python3 online submissions for Search a 2D Matrix.

ironically this isn't super bad
but I get it should not have extra memory
"""


class Solution2:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        M, N = len(matrix), len(matrix[0])
        l, r = 0, M*N-1

        while l < r:
            mid = l+(r-l)//2
            if matrix[mid//N][mid % N] < target:
                l = mid+1
            else:
                r = mid

        return matrix[l//N][l % N] == target


"""
[[1,1]]
2
hmm.. 

the coordinate calculation is wrong should l%N for column 

also.. when l=1, 1/M will be 1 and there is no idx-1 row..
ah shit. l should also l/N, not M... 


Runtime: 55 ms, faster than 74.98% of Python3 online submissions for Search a 2D Matrix.
Memory Usage: 14.4 MB, less than 42.57% of Python3 online submissions for Search a 2D Matrix.

oh shit.. the binary search tree is beyond smart..
"""


class Solution:
    def searchMatrix(self, matrix: List[List[int]], target: int) -> bool:
        M, N = len(matrix), len(matrix[0])
        x, y = 0, N-1

        while x < M and y >= 0:
            if matrix[x][y] > target:
                y -= 1
            elif matrix[x][y] < target:
                x += 1
            else:
                return True
        return False


'''
Runtime: 48 ms, faster than 89.69% of Python3 online submissions for Search a 2D Matrix.
Memory Usage: 14.5 MB, less than 42.57% of Python3 online submissions for Search a 2D Matrix.
'''
