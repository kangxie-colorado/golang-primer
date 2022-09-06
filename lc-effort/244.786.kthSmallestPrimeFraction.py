"""
https: // leetcode.com/problems/k-th-smallest-prime-fraction/

I think using two heaps
"""


import heapq
from typing import List


class Solution:
    def kthSmallestPrimeFraction(self, arr: List[int], k: int) -> List[int]:
        minH = arr
        maxH = [-i for i in arr]

        heapq.heapify(minH)
        heapq.heapify(maxH)

        minTop = maxTop = 0
        while k > 0:
            minTop = heapq.heappop(minH)
            maxTop = -heapq.heappop(maxH)

            if minH[0]/maxTop < -minTop/maxH[0]:
                heapq.heappush(maxH, -maxTop)
            else:
                heapq.heappush(minH, minTop)
            k -= 1

        return [minTop, maxTop]


""""
naive and wrong...
I might just be over-complicating things
just use the fractional and this matrix
    5   3   2   1
1 1/5
2
3
5

or maybe make it conceptually easier
    1   2   3   5
1               1/5
2
3
5

"""


class Solution:
    def kthSmallestPrimeFraction(self, arr: List[int], k: int) -> List[int]:
        h = []  # (f, i,j)
        for i, n in enumerate(arr):
            heapq.heappush(h, (n/arr[-1], i, len(arr)-1))

        n1, n2 = 0, 0
        while k > 0:
            _, i, j = heapq.heappop(h)
            n1, n2 = arr[i], arr[j]
            if j > 0:
                heapq.heappush(h, (arr[i]/arr[j-1], i, j-1))
            k -= 1
        return [n1, n2]


"""
Runtime: 1543 ms, faster than 67.27% of Python3 online submissions for K-th Smallest Prime Fraction.
Memory Usage: 14.1 MB, less than 89.46% of Python3 online submissions for K-th Smallest Prime Fraction.
"""


if __name__ == '__main__':
    s = Solution()
    print(s.kthSmallestPrimeFraction(arr=[1, 2, 3, 5], k=3))
