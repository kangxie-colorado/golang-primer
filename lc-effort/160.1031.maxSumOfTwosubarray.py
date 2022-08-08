"""
https://leetcode.com/problems/maximum-sum-of-two-non-overlapping-subarrays/

not much idea
but the space is smaller... so let me see brute force first
hope it can jog some thoughts
"""


from ast import List
from curses import nonl
from heapq import heapify
import heapq


class Solution:
    def maxSumTwoNoOverlap(self, nums, firstLen: int, secondLen: int) -> int:
        l1, l2 = firstLen, secondLen
        if l1 < l2:
            l1, l2 = l2, l1

        presum = 0
        for i in range(len(nums)):
            nums[i] += presum
            presum = nums[i]

        # get two window sums
        # notice there are only len - windownSize + 1 sums..
        def getWindsum(l):
            presum = 0
            windSums = [0]*len(nums)
            for i in range(len(nums)):
                if i >= l:
                    presum = nums[i-l]
                windSums[i] = nums[i] - presum
            return windSums

        windSums1 = getWindsum(l1)[l1-1:]
        windSums2 = getWindsum(l2)[l2-1:]

        def overlap(r1, r2):
            return not (r1[1] < r2[0] or r1[0] > r2[1])

        # cross over two window Sums find out the max..
        maxSum = 0
        for i, s1 in enumerate(windSums1):
            for j, s2 in enumerate(windSums2):
                # end-start = l-1: so knowing end, the start=end-l+1
                # then notice, i starts with l1-1; j starts with l2-1
                # so also offset that, start=end-l+1 + l-1 = end; and end becomes end+l-1
                if not overlap([i, i+l1-1], [j, j+l2-1]):  # [:] .. close at right/left
                    maxSum = max(maxSum, windSums1[i]+windSums2[j])
        return maxSum


"""
Runtime: 1516 ms, faster than 6.91% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.
Memory Usage: 14.2 MB, less than 25.75% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.

apparently there is something better out there..
but I cannot see

I can see a possible max-heap solution
push two window sums into the heap..

then pop.. if they don't overlap.. cool.. return
otherwise, discard one and get another... it then becomes DP problem

I feel it has a slight chance to be right.. but not sure.. let me try
"""


class Solution:
    def maxSumTwoNoOverlap(self, nums, firstLen: int, secondLen: int) -> int:
        l1, l2 = firstLen, secondLen
        presum = 0
        for i in range(len(nums)):
            nums[i] += presum
            presum = nums[i]

        # get two window sums
        # notice there are only len - windownSize + 1 sums..
        def getWindsum(l):
            presum = 0
            windSums = [0]*len(nums)
            for i in range(len(nums)):
                if i >= l:
                    presum = nums[i-l]
                windSums[i] = nums[i] - presum
            return windSums

        windSums1 = getWindsum(l1)[l1-1:]
        windSums2 = getWindsum(l2)[l2-1:]

        def overlap(r1, r2):
            return not (r1[1] < r2[0] or r1[0] > r2[1])

        def heapify(windSums):
            h = []
            for i, n in enumerate(windSums):
                heapq.heappush(h, (-n, i))
            return h

        maxHeap1 = heapify(windSums1)
        maxHeap2 = heapify(windSums2)
        memtable = [[]]*(len(maxHeap1)+1)
        for i in range(len(memtable)):
            memtable[i] = [-1]*(len(maxHeap2)+1)

        def getMaxLen(heap1, heap2):
            if len(heap1) == 0 or len(heap2) == 0:
                return 0

            if memtable[len(heap1)][len(heap2)] != -1:
                return memtable[len(heap1)][len(heap2)]

            s1 = heapq.heappop(heap1)
            s2 = heapq.heappop(heap2)
            nonlocal l1, l2

            if not overlap([s1[1], s1[1]+l1-1], [s2[1], s2[1]+l2-1]):
                heapq.heappush(heap1, s1)
                heapq.heappush(heap2, s2)
                memtable[len(heap1)][len(heap2)] = -(s1[0]+s2[0])
            else:
                heapq.heappush(heap1, s1)
                heapq.heappush(heap2, s2)

                heapq.heappop(heap2)
                len1 = getMaxLen(heap1, heap2)
                heapq.heappush(heap2, s2)

                heapq.heappop(heap1)
                len2 = getMaxLen(heap1, heap2)
                heapq.heappush(heap1, s1)

                memtable[len(heap1)][len(heap2)] = max(len1, len2)

            return memtable[len(heap1)][len(heap2)]
        return getMaxLen(maxHeap1, maxHeap2)


"""
TLE... hmm..
worse that brute force...

Runtime: 4918 ms, faster than 5.02% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.
Memory Usage: 29.1 MB, less than 8.63% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.

so congratulations.. I got another horrible solution

alright.. I am just SREII... 
I wonder if the senior/staff can solve this easily

leave this to tomorrow.. I must be missing some foundation
"""

"""
okay.. oncall week and I took a break from leetcode
now lets resume it..


"""


class Solution:
    def maxSumTwoNoOverlap(self, nums, firstLen: int, secondLen: int) -> int:
        L, M = firstLen, secondLen

        presum = 0
        for i in range(len(nums)):
            nums[i] += presum
            presum = nums[i]

        maxL, maxM, res = nums[L-1], nums[M-1], nums[L+M-1]
        for i in range(L+M, len(nums)):
            maxL = max(maxL, nums[i-M]-nums[i-M-L])
            maxM = max(maxM, nums[i-L]-nums[i-L-M])

            newM = nums[i]-nums[i-M]
            newL = nums[i]-nums[i-L]

            res = max(res, newM+maxL, newL+maxM)

        return res


"""
Runtime: 88 ms, faster than 52.43% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.
Memory Usage: 14 MB, less than 89.33% of Python3 online submissions for Maximum Sum of Two Non-Overlapping Subarrays.

cool...
"""

if __name__ == '__main__':
    s = Solution()
    print(s.maxSumTwoNoOverlap([0, 6, 5, 2, 2, 5, 1, 9, 4], 1, 2))
    print(s.maxSumTwoNoOverlap([0, 6, 5, 2, 2, 5, 1, 9, 4], 4, 2))
