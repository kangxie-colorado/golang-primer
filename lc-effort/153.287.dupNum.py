"""
https://leetcode.com/problems/find-the-duplicate-number/

I remember that using xor to find out the missing number
will it also able to find out the duplciate?

1-n xor 1-n will be zero
the result xor the dup number will just be dup number?

"""


from gettext import find


class Solution(object):
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        n = len(nums)-1
        xor = 0

        for i, n in enumerate(nums[:-1]):
            xor ^= (i+1)
            xor ^= n

        return xor ^ nums[-1]


"""
[2,2,2,2,2] expect 2, got 6

haha.. it is not each number for every 1-n 

when you can modify -- like sorting, this is trivial
when you can use o(n) memory, like a hash map.. this is trivial
when you can just do O(n^2) search, like two loop, this is also trivial

but the requirement declines the first 2, the constraint 1 <= n <= 10^5 denies the 3rd
so have to think something else

so I came up with this
1-n as the possible number, binary search it
for each m, have to iterate the array to find the number that is <= this m

if more than m+1, then you find a range.. to the left
     you can also keep a equal case to accelerate the best case... but not necessary
otherwise, to the right..
"""


class Solution(object):
    def findDuplicate(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        l, r = 1, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            smallerEqual = sum([1 for n in nums if n <= m])
            if smallerEqual > m:
                r = m
            else:
                l = m+1

        return l

        """
        Runtime: 735 ms, faster than 63.87% of Python online submissions for Find the Duplicate Number.
        Memory Usage: 25.3 MB, less than 57.59% of Python online submissions for Find the Duplicate Number.
        """


"""
that O(N) using cycle detection is super cool and I think I do understand it now
let me code it up
"""


class Solution(object):
    def findDuplicate(self, nums):
        slow = nums[0]
        fast = nums[nums[0]]
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]

        finder = 0
        while finder != slow:
            finder = nums[finder]
            slow = nums[slow]

        return finder
