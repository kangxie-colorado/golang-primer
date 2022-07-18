"""
https://leetcode.com/problems/sort-colors/

not much idea..
maybe on 0, swap to left
on 2, swap to right?

and maintain l-idx, right idx...
"""


from ast import List


class Solution:
    def sortColors(self, nums) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        l, r = 0, len(nums)-1
        i = 0
        while i <= r:
            if nums[i] == 1:
                i += 1
            elif nums[i] == 0:
                # think this as partition:
                # [0:l] is all zero.. and idx-l will be swapped to zero
                # whatsoever, [:l] cannot have 2... so no need to worry about I swapped a 2 out of left side
                # but for the right side, this is very valid
                # yeah, somehow I felt it... but I did't prove it and this is it
                nums[i], nums[l] = nums[l], nums[i]
                l += 1
                i += 1  # i should also walk forward
            else:
                nums[i], nums[r] = nums[r], nums[i]
                r -= 1


"""
Runtime: 32 ms, faster than 96.08% of Python3 online submissions for Sort Colors.
Memory Usage: 13.9 MB, less than 63.66% of Python3 online submissions for Sort Colors.

yeah.. it is a bit tricky to figure out how to walk the pointers..
the important part is to make it generic and simplify

to do that, it is important to notice when walking l, i should walk too
otherwise... it might not get right

also l/r is only the place holder for next 0/2.. it is not necessarily 0/2 at l/r position until the swap happens
so l/r should not be skipped by any means
"""

if __name__ == '__main__':
    s = Solution()
    l = [2, 0, 2, 1, 1, 0]
    s.sortColors(l)
    print(l)
    l = [2, 0, 1]
    s.sortColors(l)
    print(l)
    l = [0, 0]
    s.sortColors(l)
    print(l)
    l = [1, 0]
    s.sortColors(l)
    print(l)
