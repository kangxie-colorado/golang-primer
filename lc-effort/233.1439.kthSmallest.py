"""
https://leetcode.com/problems/find-the-kth-smallest-sum-of-a-matrix-with-sorted-rows/

first I see a heap solution..

what should the heap save...
1. ?
2,3 the x and y of course

what should 1st part in the heap element..
not the raw value I think it should be the diff
"""


from bisect import bisect
import heapq
from typing import List


class Solution:
    def kthSmallest(self, mat: List[List[int]], k: int) -> int:
        h = []
        S = 0
        m, n = len(mat), len(mat[0])
        for r in mat:
            S += r[0]
            if len(r) > 1:
                heapq.heappush(h, (r[1]-r[0], 1, 0))
        k -= 1
        while k > 0:
            valdiff, x, y = heapq.heappop(h)
            S += valdiff

            if y == n - 1:
                heapq.heappush(h, (5000, -1, -1))
            else:
                heapq.heappush(h, (mat[x][y+1]-mat[x][y], x, y+1))
            k -= 1

        return S


"""
oh.. logical error
1 3 11
2 4 6

[1,3] -> 3

[3,2] -> 5.. but you missed [1 4]
so instead of outting on S, I need to do m sums... ???
then I move one
then I do binary search in the cached results???

"""


class Solution:
    def kthSmallest(self, mat: List[List[int]], k: int) -> int:
        h = []  # (delta, x,y)
        S = 0
        m, n = len(mat), len(mat[0])
        for i in range(m):
            r = mat[i]
            S += r[0]
            if len(r) > 1:
                heapq.heappush(h, (r[1]-r[0], i, 1))

        k -= 1
        while k > 0:
            if k < m:
                # in this range
                while k > 0:
                    delta, _, _ = heapq.heappop(h)
                    k -= 1

                return S + delta

            delta, x, y = heapq.heappop(h)
            S += delta
            if y >= n - 1:
                heapq.heappush(h, (10000, - 1, -1))
            else:
                heapq.heappush(
                    h, (mat[x][y+1] - mat[x][y], x, y+1))
            k -= m

        return S


"""
okay... cannot get it right..
someone says using kSmallestPairs as the merge function
"""


class Solution:
    def kthSmallest(self, mat: List[List[int]], k: int) -> int:
        m, n = len(mat), len(mat[0])
        res = mat[0]
        for i in range(1, m):
            res = self.kSmallestPairs(res, mat[i])
        return res[k-1]

    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int = 200) -> List[List[int]]:
        A, B = nums1, nums2
        m, n = len(A), len(B)

        h = []
        for i in range(m):
            heapq.heappush(h, (A[i]+B[0], i, 0))

        res = []
        while k > 0 and h:
            val, x, y = heapq.heappop(h)
            res.append(val)
            if y < n-1:
                heapq.heappush(h, (A[x]+B[y+1], x, y+1))
            k -= 1

        return res


"""
Runtime: 293 ms, faster than 55.85% of Python3 online submissions for Find the Kth Smallest Sum of a Matrix With Sorted Rows.
Memory Usage: 14.1 MB, less than 85.64% of Python3 online submissions for Find the Kth Smallest Sum of a Matrix With Sorted Rows.

wow...
"""


if __name__ == "__main__":
    s = Solution()
    assert 13 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=7)
    assert 7 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=5)

    assert 3 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=1)
    assert 5 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=2)
    assert 5 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=3)
    assert 7 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=4)
    assert 7 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=5)
    assert 9 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=6)
    # assert 13 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=7)
    assert 15 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=8)
    assert 17 == s.kthSmallest(mat=[[1, 3, 11], [2, 4, 6]], k=9)
