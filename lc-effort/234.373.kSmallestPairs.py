"""
https://leetcode.com/problems/find-k-pairs-with-smallest-sums/


turns out I did this before..
but anyway I had forgotten let me do it since I see it
"""


import heapq
from typing import List


class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        m, n = len(nums1), len(nums2)
        A, B = nums1, nums2

        h = heapq.heapify([(A[0]+B[0], 0, 0)])

        res = []
        while k > 0:
            _, x, y = heapq.heappop(h)
            res.append(nums1[x], nums2[y])

            if y < n-1:
                heapq.heappush(A[x]+B[y+1], x, y+1)
            if x < m-1:
                heapq.heappush(A[x+1]+B[y], x+1, y)

            k += 1

        return res


"""
I didn't start from a whole colum.. and only extend to right...
I start with one point..
then I have to deal with duplicate add

so I bring the M back to mark
"""


class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        m, n = len(nums1), len(nums2)
        A, B = nums1, nums2
        M = [[]] * m
        for i in range(m):
            M[i] = [0]*n

        h = []
        heapq.heappush(h, (A[0]+B[0], 0, 0))

        res = []
        while k > 0 and h:
            _, x, y = heapq.heappop(h)
            if M[x][y] == 1:
                continue
            M[x][y] = 1
            res.append([nums1[x], nums2[y]])

            if y < n-1:
                heapq.heappush(h, (A[x]+B[y+1], x, y+1))
            if x < m-1:
                heapq.heappush(h, (A[x+1]+B[y], x+1, y))

            k -= 1

        return res


"""
huh.. memory limit exceeded...

>>> len(a)
21435
if b is similar length
I would need 400M... indeed... for M alone...

change to set for a try
"""


class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        m, n = len(nums1), len(nums2)
        A, B = nums1, nums2
        seen = set((0, 0))

        h = []
        heapq.heappush(h, (A[0]+B[0], 0, 0))

        res = []
        while k > 0 and h:
            _, x, y = heapq.heappop(h)
            res.append([A[x], B[y]])

            if y < n-1 and (x, y+1) not in seen:
                seen.add((x, y+1))
                heapq.heappush(h, (A[x]+B[y+1], x, y+1))
            if x < m-1 and (x+1, y) not in seen:
                seen.add((x+1, y))
                heapq.heappush(h, (A[x+1]+B[y], x+1, y))

            k -= 1

        return res


"""
Runtime: 2012 ms, faster than 40.57% of Python3 online submissions for Find K Pairs with Smallest Sums.
Memory Usage: 33.4 MB, less than 64.08% of Python3 online submissions for Find K Pairs with Smallest Sums.

Runtime: 1187 ms, faster than 96.28% of Python3 online submissions for Find K Pairs with Smallest Sums.
Memory Usage: 33.3 MB, less than 69.05% of Python3 online submissions for Find K Pairs with Smallest Sums.

let me just recover the memory for doing one column
"""


class Solution:
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        A, B = nums1, nums2
        m, n = len(A), len(B)

        h = []
        for i in range(m):
            heapq.heappush(h, (A[i]+B[0], i, 0))

        res = []
        while k > 0 and h:
            _, x, y = heapq.heappop(h)
            res.append([A[x], B[y]])
            if y < n-1:
                heapq.heappush(h, (A[x]+B[y+1], x, y+1))
            k -= 1

        return res


"""
Runtime: 1322 ms, faster than 87.15% of Python3 online submissions for Find K Pairs with Smallest Sums.
Memory Usage: 40 MB, less than 11.26% of Python3 online submissions for Find K Pairs with Smallest Sums.

now go to 1439 use this function as a merge function for any two rows...
"""


if __name__ == '__main__':
    s = Solution()
    print(s.kSmallestPairs(nums1=[1, 2], nums2=[3], k=3))
