
"""
https://leetcode.com/problems/next-permutation/

I didn't do this before?
this is well known to me now

from right, find the increasing sequence, where it ends..
where it ends will be the swap point then sort the right..

if it increasing all the way... then just sort
"""


class Solution(object):
    def nextPermutation(self, nums):
        """
        :type nums: List[int]
        :rtype: None Do not return anything, modify nums in-place instead.
        """
        for i in range(len(nums)-1, -2, -1):
            if i >= 0 and i+1 < len(nums) and nums[i] < nums[i+1]:
                break
        # where the i will end up to
        # i is breaking point for above condition.. so i should be at the left swap idx
        # it is easy to mistake it as the right swap idx.. ugh
        # but yet... to the end of cliff... i==0 is another break point..
        # so I need to allow it goes off the edge.. thus i==-1
        if i == -1:
            nums.sort()
        else:
            # nah... cannot swap the biggest from right to left
            # swap the first number bigger than left
            for j in range(len(nums)-1, i, -1):
                if nums[j] > nums[i]:
                    nums[j], nums[i] = nums[i], nums[j]
                    break
            a = nums[i+1:]
            nums[i+1:] = sorted(a)


"""
Runtime: 25 ms, faster than 94.30% of Python online submissions for Next Permutation.
Memory Usage: 13.2 MB, less than 91.30% of Python online submissions for Next Permutation.

considering this... I call it a win yeh?
Accepted 853,772 Submissions 2,328,372

so now.. think why searching for increasing(non-decreasing sequence)
because if the right part is non-decreasing that means there is no next permutaion in that part
of course, when the whole array is non-decreasing.. that is a different edge case

I was lucky I didn''t overthink these... otherwise I might be dragged down by my thoughts...

when you found the partion that is already at its max permutaion, now to the left is the one to swap to make next.. 
but which number to swap? the biggest one???
of course not... the one that is just bigger that left from the right
"""

if __name__ == '__main__':
    s = Solution()
    l = [1, 2, 3]
    s.nextPermutation(l)
    print(l)
    l = [3, 2, 1]
    s.nextPermutation(l)
    print(l)
    l = [1, 3, 2, 1]
    s.nextPermutation(l)
    print(l)
