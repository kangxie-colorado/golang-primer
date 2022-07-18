"""
https://leetcode.com/problems/increasing-triplet-subsequence/

there was that problem of 132 pattern.. and this simplifies to 123 pattern
seems easier to do

so I just need to keep track of last minimum before me...

"""


from typing import Tuple


class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        lastMin = (0, nums[0])
        numJ = (-1, -1)
        for i in range(1, len(nums)):
            if nums[i] <= lastMin[1]:
                lastMin = (i, nums[i])
            elif (nums[i] > lastMin[1] and numJ[0] == -1) or \
                 (numJ[0] != -1 and nums[i] <= numJ[1]):
                numJ = (i, nums[i])
            elif nums[i] > numJ[1]:
                return True
        return False


"""
Runtime: 503 ms, faster than 87.28% of Python online submissions for Increasing Triplet Subsequence.
Memory Usage: 33.7 MB, less than 19.34% of Python online submissions for Increasing Triplet Subsequence.

yeah.. I still used the same swap card algorithm from 132 pattern
the idea is
1. when < min, swap min
    think [2,4,1,5] --- 1 can replace 2, then 5 comes, it forms 1,4,5..
    although that is not really the sequence but it doesn't change the logical
    becase 2<4, and 1<2.. replacing doesn't change the semantics
2. when < numJ, swap numJ -- this is a subtle part..
    think [2,1,5,0,4,5]
    after 0 replaces 1, you need to also replace 5 with the 4...
    otherwise, you are left with 0 5 5...

    so yeah, this is kind of tricky.. but still very fun
    I don't know if I have over-think of not...

    public boolean increasingTriplet(int[] nums) {
        // start with two largest values, as soon as we find a number bigger than both, while both have been updated, return true.
        int small = Integer.MAX_VALUE, big = Integer.MAX_VALUE;
        for (int n : nums) {
            if (n <= small) { small = n; } // update small if n is smaller than both
            else if (n <= big) { big = n; } // update big only if greater than small but smaller than big
            else return true; // return if you find a number bigger than both
        }
        return false;
    }

    same idea but kind of cleaner ... rewrite my own as follow

"""


class Solution(object):
    def increasingTriplet(self, nums):
        """
        :type nums: List[int]
        :rtype: bool
        """
        lastMin = nums[0]
        numJ = 2**31-1
        for i in range(1, len(nums)):
            if nums[i] <= lastMin:
                lastMin = nums[i]
            elif nums[i] <= numJ:
                numJ = nums[i]
            else:
                return True
        return False
