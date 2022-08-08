"""
https://leetcode.com/problems/shuffle-the-array/

I wasn't able to come up a solution with no extra space
https://leetcode.com/problems/shuffle-the-array/discuss/675956/In-Place-O(n)-Time-O(1)-Space-With-Explanation-and-Analysis

then this post talks about combining two numbers in one... 
that gave me an idea..

I don't need to follow his steps
because of 1 <= nums[i] <= 10^3

I just combine x1 y1 as x1*1001 + y1 => z1
then z1//1001 will be x1
z1%1001 will be y1..

I save this at the 0 2 4 indexs... then populate the results
"""


from typing import List


class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        idx = 0
        for n1, n2 in zip(nums[:n], nums[n:]):
            nums[idx] = n1*1001 + n2
            idx += 2

        for i in range(0, len(nums), 2):
            nums[i+1] = nums[i] % 1001
            nums[i] //= 1001
        return nums


"""
Runtime: 113 ms, faster than 25.20% of Python3 online submissions for Shuffle the Array.
Memory Usage: 14.2 MB, less than 41.37% of Python3 online submissions for Shuffle the Array.
"""
