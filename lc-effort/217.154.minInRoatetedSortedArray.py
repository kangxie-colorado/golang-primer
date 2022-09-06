"""
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

I kind of see a recursion solution
"""


from tkinter import N
from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        def helper(l, r):
            if l == r:
                return nums[l]

            m = l+(r-l)//2
            if nums[l] < nums[m] < nums[r]:
                return nums[l]

            if nums[l] < nums[m]:
                return min(nums[l], helper(m+1, r))
            if nums[m] < nums[r]:
                return helper(l, m)

            return min(nums[l:r+1])

        return helper(0, len(nums)-1)


"""
179 / 193 test cases passed.

not too bad.. much better than I thougt
however failed here

[3,3,1,3]
it is actually this case
            if nums[l] <= nums[m] <= nums[r]:   
                return nums[l]

so I remove = from left half.. it passes..
will it pass the whole set?
            if nums[l] < nums[m] <= nums[r]:
                return nums[l]
not it failed here
[10,1,10,10,10]

so I see what is going on
need to be this
            if nums[l] < nums[m] < nums[r]:
                return nums[l]

the only situation you cannot tell right/left is when they all eqaul
because 
[10 1 10 10 10]
1 can appear in idx-1, idx-3.. and the test above will fail to differentiate
so okay... need to take the equal case away

Runtime: 61 ms, faster than 85.82% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.
Memory Usage: 14.9 MB, less than 9.68% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.

so it passes...
still kind of unsure why the above three if are comprehensive already

            if nums[l] < nums[m] < nums[r]:
                return nums[l]
                ^ this is the lucky case, capture it.. 


            if nums[l] < nums[m]:
                return min(nums[l], helper(m+1, r))
                ^ this is not ambiguous either

            if nums[m] < nums[r]:
                return helper(l, m)
                ^ this is not ambiguous either

            return min(nums[l:r+1])
                ^ should be more tests or this is all

what if nums[l] >= nums[m] and nums[r] > nums[m], this seems to be captured by 
            if nums[m] < nums[r]: # without the equal case
what if nums[m] <= nums[r] and nums[l] > nums[m]
like [4 5 6 7 0 1 2 4] the 1... it can only be in [l:m] inclusive
this is again captured by above same test but without the equal case

nums[l]>nums[m]>nums[r] apparently doesn't stand valid 
so it kind of become okay?

not sure.. let me check the disccusions now.
okay
https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/discuss/167981/Beats-100-Binary-Search-with-Explanations

this dancing programmer one-upped me
there are actually another way to look at it..

instead of looking at sort part.. how about unsorted part..
and isn't it fun, she moves l/r by 1 towards m
"""


class Solution:
    def findMin(self, nums: List[int]) -> int:
        def helper(l, r):
            if l == r:
                return nums[l]

            m = l+(r-l)//2
            if nums[l] < nums[m] < nums[r]:
                return nums[l]

            if nums[l] < nums[m]:
                return min(nums[l], helper(m+1, r))
            elif nums[l] > nums[m]:
                return helper(l+1, m)
            elif nums[m] < nums[r]:
                return helper(l, m)
            elif nums[m] > nums[r]:
                return helper(m+1, r)
            else:
                return min(nums[l:r+1])

        return helper(0, len(nums)-1)


"""
Runtime: 61 ms, faster than 85.82% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.
Memory Usage: 14.7 MB, less than 9.68% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.

okay.. this is kind of the same but ugly.. ugly

this move r by one position https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/discuss/48808/My-pretty-simple-code-to-solve-it
is O(n) but it kind of shows the idea, get rid of the invalid entries one at a time

yeah... if you can get rid of half.. you get rid of half.. if you cannot, get rid of one..
woow....

let me just type it for practice

"""


class Solution:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1

        while l < r:
            m = l+(r-l)//2

            if nums[m] > nums[r]:
                # get rid of the left
                l = m+1
            elif nums[m] < nums[r]:
                # get rid of the right
                r = m
            else:  # nums[m] == nums[r]
                # get rid of the right element
                r -= 1

        return nums[l]


"""
Runtime: 86 ms, faster than 48.47% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.
Memory Usage: 14.4 MB, less than 42.95% of Python3 online submissions for Find Minimum in Rotated Sorted Array II.

I think by symmetry I can do left as well
well.. maybe not...
"""


class Solution_oops:
    def findMin(self, nums: List[int]) -> int:
        l, r = 0, len(nums)-1

        while l < r:
            m = l+(r-l)//2

            if nums[m] < nums[l]:
                # get rid of the right
                r = m
            elif nums[m] > nums[l]:
                # get rid of the right again?????
                r = m
            else:  # nums[m] == nums[r]
                # get rid of the right element
                r -= 1

        return nums[l]


if __name__ == '__main__':
    s = Solution()
    print(s.findMin([2, 2, 1]))
    print(s.findMin([2, 2, 0, 1]))
    print(s.findMin([4, 5, 6, 6, 6, 7, 0, 0, 1, 2, 3, 3, 3, 4]))
