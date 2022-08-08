"""
https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

binary search to start? I am not quite sure how the duplicates could impact things
I'll search and see which side is sorted... and make a call there..
"""


from typing import List


class Solution_Failed:
    def search(self, nums: List[int], target: int) -> bool:
        l, r = 0, len(nums)-1
        while l < r:
            mid = l+(r-l)//2

            if nums[mid] == target or nums[l] == target or nums[r] == target:
                return True

            if nums[mid] >= nums[l]:
                # left already sorted
                if nums[l] < target < nums[mid]:
                    r = mid-1
                else:
                    l = mid+1

            if nums[mid] <= nums[r]:
                # right already sorted
                if nums[mid] < target < nums[r]:
                    l = mid+1
                else:
                    r = mid-1

        return False


"""
I knew this is not going to be right but not disappointed at all
269 / 280 test cases passed.

failed this.. let me see
[1,0,1,1,1]
0

with duplicate values.. 
1,0,1   will meet nums[mid] >= nums[l].. but it is not sorted anymore...
l   mid
"""


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        l, r = 0, len(nums)-1
        while l < r:
            mid = l+(r-l)//2

            if nums[mid] == target or nums[l] == target or nums[r] == target:
                return True

            if nums[mid] > nums[l]:
                # left already sorted
                if nums[l] < target < nums[mid]:
                    r = mid-1
                else:
                    l = mid+1

            if nums[mid] < nums[r]:
                # right already sorted
                if nums[mid] < target < nums[r]:
                    l = mid+1
                else:
                    r = mid-1

            if nums[mid] == nums[l] or nums[mid] == nums[r]:
                # cannot tell... if it is sorted or not
                # only able to do linear serach now
                return True if any(i == target for i in nums[l:r+1]) else False
        return nums[l] == target


"""
279 / 280 test cases passed.

failed 
[1]
1

are you fucking ...
just change 
        return nums[l] == target

Runtime: 102 ms, faster than 22.60% of Python3 online submissions for Search in Rotated Sorted Array II.
Memory Usage: 14.6 MB, less than 22.46% of Python3 online submissions for Search in Rotated Sorted Array II.

it is very ugly let me think
"""


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        l, r = 0, len(nums)-1
        while l < r:
            mid = l+(r-l)//2

            if nums[mid] == target or nums[l] == target or nums[r] == target:
                return True

            if nums[mid] == nums[l] or nums[mid] == nums[r]:
                # cannot tell... if it is sorted or not
                # only able to do linear serach now
                return True if any(i == target for i in nums[l:r+1]) else False

            if (nums[mid] > nums[l] and nums[l] < target < nums[mid]) or \
                    (nums[mid] < nums[r] and not (nums[mid] < target < nums[r])):
                r = mid-1
            else:
                l = mid+1

        return nums[l] == target


"""
Runtime: 64 ms, faster than 80.69% of Python3 online submissions for Search in Rotated Sorted Array II.
Memory Usage: 14.4 MB, less than 59.18% of Python3 online submissions for Search in Rotated Sorted Array II.
"""


class Solution:
    def search(self, nums: List[int], target: int) -> bool:
        l, r = 0, len(nums)-1
        while l < r:
            mid = l+(r-l)//2

            if nums[mid] == target or nums[l] == target or nums[r] == target:
                return True

            if nums[mid] != nums[l] and nums[mid] != nums[r]:
                if (nums[mid] > nums[l] and nums[l] < target < nums[mid]) or \
                        (nums[mid] < nums[r] and not (nums[mid] < target < nums[r])):
                    r = mid-1
                else:
                    l = mid+1
            elif nums[mid] == nums[l] and (nums[mid] < nums[r] and nums[mid] < target < nums[r]):
                l = mid+1
            elif nums[mid] == nums[r] and (nums[mid] > nums[l] and nums[l] < target < nums[mid]):
                r = mid-1
            else:
                # cannot tell... if it is sorted or not
                # only able to do linear serach now
                return True if any(i == target for i in nums[l:r+1]) else False

        return nums[l] == target


"""
Runtime: 65 ms, faster than 79.26% of Python3 online submissions for Search in Rotated Sorted Array II.
Memory Usage: 14.5 MB, less than 59.18% of Python3 online submissions for Search in Rotated Sorted Array II.
"""


class Solution_Fail:
    def search(self, nums: List[int], target: int) -> bool:
        l, r = 0, len(nums)-1
        while l < r:
            mid = l+(r-l)//2

            if nums[mid] == target or nums[l] == target or nums[r] == target:
                return True

            # this below will not cover [3,1,2,2,2] 1
            if nums[mid] == nums[l] == nums[r]:
                return True if any(i == target for i in nums[l:r+1]) else False
            else:
                if (nums[mid] > nums[l] and nums[l] < target < nums[mid]) or \
                        (nums[mid] < nums[r] and not (nums[mid] < target < nums[r])):
                    r = mid-1
                else:
                    l = mid+1

        return nums[l] == target


if __name__ == '__main__':

    print(Solution().search([1, 0, 1, 1, 1], 0))
