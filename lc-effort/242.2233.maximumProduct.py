"""
https://leetcode.com/problems/maximum-product-after-k-increments/

this is quite like the row-based heap
calculate the the product if increasing this row number by 1..

but will the one by one max, render the overall maxsum?

let me see       next prod      next prod
            6       126,7           189,7
            3       144,4           216,4
            3       144,4           216,4
            2       162,3           216,4
init prod 108       162

maybe it can
but looks like I need to keep a tag on the value's changing history 2->3->4
so maybe in the heap I save index then use index to refer back to values

try a bit

also when calculate init sum, if a 0, increment it to 1 and reduce k by 1
otherwise, nothing can move
"""


from functools import reduce
from sortedcontainers import SortedList
import enum
import heapq
from multiprocessing import heap
from typing import List


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        baseProd = -1
        for i, n in enumerate(nums):
            if n == 0 and k > 0:
                nums[i] = 1
                k -= 1
            baseProd *= nums[i]
        if k == 0:
            return -baseProd
        h = []  # (newProd, index)
        for i, n in enumerate(nums):
            heapq.heappush(h, (baseProd//n*(n+1), i))

        while h and k:
            baseProd, idx = heapq.heappop(h)
            nums[idx] += 1
            k -= 1
            h = []
            for i, n in enumerate(nums):
                heapq.heappush(h, (baseProd//n*(n+1), i))

        return -baseProd % (10**9+7)


"""
not too far and TLE
27 / 73 test cases passed.

maybe I should just use float number to record the multiples..
"""


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        baseProd = 1
        for i, n in enumerate(nums):
            if n == 0 and k > 0:
                nums[i] = 1
                k -= 1
            baseProd *= nums[i]
        if k == 0:
            return baseProd
        h = []
        for i, n in enumerate(nums):
            heapq.heappush(h, (-(n+1)/n, i))

        while k:
            magifier, idx = heapq.heappop(h)
            n = nums[idx] = nums[idx]+1

            k -= 1
            baseProd *= magifier
            heapq.heappush(h, (-(n+1)/n, idx))

        return round(abs(baseProd)) % (10**9+7)


"""
[21,100]
58

as expected, 7899 vs 7900
the float will lose precisions..

change to round failed here
[92,36,15,84,57,60,72,86,70,43,16]
62

800212627
800222867

this is beyond my ability to debug...
so okay.. when I am doing float it cames to me
the smaller number it is, the magnifier will be bigger... so just simply use the smallest number
"""


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        prod = 1
        for i, n in enumerate(nums):
            if n == 0 and k > 0:
                nums[i] = 1
                k -= 1
            prod *= nums[i]
        if k == 0:
            return prod

        h = []
        for n in nums:
            heapq.heappush(h, n)

        while k:
            n = heapq.heappop(h)
            prod = prod//n*(n+1)
            heapq.heappush(h, n+1)
            k -= 1
        return prod % (10**9+7)


"""
54 / 73 test cases passed and TLE..

hmm.. 
let me change to SortedList
"""


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        prod = 1
        for i, n in enumerate(nums):
            if n == 0 and k > 0:
                nums[i] = 1
                k -= 1
            prod *= nums[i]
        if k == 0:
            return prod

        sl = SortedList(nums)
        while k:
            n = sl[0]
            prod = prod//n*(n+1)
            k -= 1
            sl.remove(n)
            sl.add(n+1)

        return prod


"""
51 / 73 test cases passed.
and TLE...

hmm...
class Solution:
    def maximumProduct(self, A, k):
        heapify(A)
        for i in range(k):
            heappush(A, heappop(A) + 1)
        return reduce(lambda p, a: p * a % (10**9 + 7), A, 1)

same idea.. but it passes
Runtime: 2189 ms, faster than 64.77% of Python3 online submissions for Maximum Product After K Increments.
Memory Usage: 24.7 MB, less than 41.48% of Python3 online submissions for Maximum Product After K Increments.

so maybe python is slow at calculating divide???
let me rewrite mine
"""


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        heapq.heapify(nums)
        while k:
            heapq.heappush(nums, heapq.heappop(nums)+1)
            k -= 1

        prod = 1
        for n in nums:
            prod *= n
            prod %= 10**9+7
        return prod


"""
Runtime: 2321 ms, faster than 57.38% of Python3 online submissions for Maximum Product After K Increments.
Memory Usage: 24.6 MB, less than 70.45% of Python3 online submissions for Maximum Product After K Increments.

yeah.. it passes
"""


class Solution:
    def maximumProduct(self, nums: List[int], k: int) -> int:
        prod = 1
        for i, n in enumerate(nums):
            if n == 0 and k > 0:
                nums[i] = 1
                k -= 1
            prod *= nums[i]
            prod %= 10**9+7
        if k == 0:
            return prod

        heapq.heapify(nums)
        while k:
            heapq.heappush(nums, heapq.heappop(nums)+1)
            k -= 1
        prod = 1
        for n in nums:
            prod *= n
            prod %= 10**9+7
        return prod


"""
hmm... this cannot pass
so it is actually the problem of first loop....
add the mod
            prod *= nums[i]
            prod %= 10**9+7

oh.. shit... it passes...
so this mod operation can sometimes mess up with correctness..
best to leave it to happen at one single pass..

don't mess up with it many times..
"""

if __name__ == '__main__':
    s = Solution()
    print(s.maximumProduct([0, 4], 5))
    print(s.maximumProduct(nums=[6, 3, 3, 2], k=2))
