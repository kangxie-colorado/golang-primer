"""
https://leetcode.com/problems/kth-largest-element-in-an-array/

O(n)? hmm... did I do this before???
then use memory

bucket sort

notice 
-10^4 <= nums[i] <= 10^4... so only 20001 numbers
"""


from typing import List


class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        buckets = [0] * 20001
        for n in nums:
            buckets[n+10000] += 1

        for i in range(20000, -1, -1):
            if k <= buckets[i]:
                return i-10000
            k -= buckets[i]


"""
Runtime: 1095 ms, faster than 61.88% of Python3 online submissions for Kth Largest Element in an Array.
Memory Usage: 27.1 MB, less than 86.61% of Python3 online submissions for Kth Largest Element in an Array.


okay.. quick select based on the partition...

okay.. this guy explains it really clear
https://leetcode.com/problems/kth-largest-element-in-an-array/discuss/1019513/Python-QuickSelect-average-O(n)-explained

the code is even more clear
class Solution:
    def findKthLargest(self, nums, k):
        if not nums: return
        pivot = random.choice(nums)
        left =  [x for x in nums if x > pivot]
        mid  =  [x for x in nums if x == pivot]
        right = [x for x in nums if x < pivot]
        
        L, M = len(left), len(mid)
        
        if k <= L:
            return self.findKthLargest(left, k)
        elif k > L + M:
            return self.findKthLargest(right, k - L - M)
        else:
            return mid[0]

it is 3*(o(n)) but thats cool
otherwise, just implement the partition function to partition in-place

actually my bucket sort is better in performance...
"""
