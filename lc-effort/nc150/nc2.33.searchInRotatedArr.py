"""
https://leetcode.com/problems/search-in-rotated-sorted-array/


done this before 
so this will be my quota of done-problems

okay... so I know because it is sorted and rotated, I just need to compare with the right side
or more verbosally, both side to figure out which side is sorted or not..

let me try to recover the code
"""


from typing import List


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            if nums[m] == target:
                return m
            # [4,5,6,7,0,1,2]
            if nums[m] <= nums[r]:
                # right side sorted
                if nums[m] <= target and target <= nums[r]:
                    l = m
                else:
                    r = m-1

            else:
                # left side sorted
                if nums[l] <= target <= nums[m]:
                    r = m
                else:
                    l = m+1
        return l if nums[l] == target else -1


"""
TLE at 
[1,3]
2

yeah I felt the coverging is less thought of
"""


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            if nums[m] == target:
                return m
            # [4,5,6,7,0,1,2]
            if nums[m] <= nums[r]:
                # right side sorted
                if nums[m] < target and target <= nums[r]:
                    l = m+1
                else:
                    r = m-1

            else:
                # left side sorted
                if nums[l] <= target < nums[m]:
                    r = m-1
                else:
                    l = m+1
        return l if nums[l] == target else -1


"""
Runtime: 94 ms, faster than 5.31% of Python3 online submissions for Search in Rotated Sorted Array.
Memory Usage: 14.2 MB, less than 55.48% of Python3 online submissions for Search in Rotated Sorted Array.

wow...
the expression is not anywhere pretty

let me think a bit more
messy still .. but I see others posting the same ideas
so anyway

and there is this smart but quite good thinking.. find the true minimal then just do regular binary search
let me try code it up

that is indeed comparing to right
"""


class Solution_Wrong:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            if nums[m] < nums[r]:
                # from m to r, sorted
                r = m
            else:
                l = m+1

        p = l

        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            m = (m+p) % len(nums)
            if nums[m] < target:
                l = m+1
            else:
                r = m
        return l if nums[l] == target else -1


"""
what is wrong here ^

the m and the realMid is kind of mixed up
m is still the index in the array 
realMid is the rotated position 

but to walk the array still m..
"""


class Solution:
    def search(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            if nums[m] < nums[r]:
                # from m to r, sorted
                r = m
            else:
                l = m+1

        p = l

        l, r = 0, len(nums)-1
        while l < r:
            m = l+(r-l)//2
            realM = (m+p) % len(nums)
            if nums[realM] < target:
                l = m+1
            else:
                r = m

        # there is a mapping (rotated mapping relationship)
        # so l is actually pointing to (l+p)%n
        l = (l+p) % len(nums)
        return l if nums[l] == target else -1


if __name__ == '__main__':
    s = Solution()
    print(s.search([4, 5, 6, 7, 0, 1, 2], 0))
