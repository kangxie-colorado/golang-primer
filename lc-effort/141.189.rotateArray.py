"""
https://leetcode.com/problems/rotate-array/

as many solutions as I can?

okay.. first thing coming to mind is save a buffer then swap
there probably will be some clever tricks... after the first one

not swap both ways i<-->j.. swap forward... i->i+k-->i+k+k
each time swap one position then k times..  O(n*k) k is k%n

"""


import math


class Solution1:
    def rotate(self, nums, k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """

        def rotateOnce():
            t = nums[len(nums)-1]
            for i in range(len(nums)-1, 0, -1):
                nums[i] = nums[i-1]

            nums[0] = t

        for i in range(k % len(nums)):
            rotateOnce()


"""
this of course TLE

then I see actually you can just swap forward N times.. it will naturally form a cycle
so same trick I can use but without k times
"""


class Solution2_wrong:  # wrong but the comment is important
    def rotate(self, nums, k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # like swap by 1 position, which I saved up the last position
        # this should also save up the last position, which will jump to 0
        # or save up the 0 idx, but start from k.. that could work too
        # hmm... a bit confusing now
        # think [1,2,3,4,5] k=3.. the cycle is 1-4-2-5-3-1
        # save up 1, then start with 1; 1 suck in 3...
        # save up 5, then start with 5; 5 suck in 2...
        # i see equivilence between them.. just save 1
        # the sucking in will hapen n-1 time...
        # after n-1 suckings, the (reverse) 1-3-5-2- have theirs values
        # 4 should have the saved up
        n = len(nums)
        k = k % n
        t = nums[0]
        toIdx = 0
        fromIdx = (n - k) % n

        for i in range(len(nums)-1):
            nums[toIdx] = nums[fromIdx]
            toIdx = fromIdx
            fromIdx = (fromIdx + n-k) % n

        nums[0+k] = t


"""
will fail
[-1,-100,3,99]
2

different cycles...
then the algorithm needs some refining.

examples show when n%k==0, there will be n/k cycles.. otherwise, only one cycle
is that true?

let me find out
"""


class Solution:
    def rotate(self, nums, k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k = k % n

        def rotateOneCycle(cyclestart, cyclelen):
            t = nums[cyclestart]
            toIdx = cyclestart
            fromIdx = (n - k+cyclestart) % n

            for i in range(cyclelen-1):
                nums[toIdx] = nums[fromIdx]
                toIdx = fromIdx
                fromIdx = (fromIdx + n-k) % n

            nums[(cyclestart+k) % n] = t

        cycles = math.gcd(n, k)
        for c in range(cycles):
            rotateOneCycle(c, n//cycles)


"""
worked on a lot more cases but failed at 
[1,2,3,4,5,6]
4

which would have actually two cycles..
some LCD(6,4) = 2 (actually GCD.. greatest common divisor)

I see.. this code can be simplified as 
        if math.gcd(n, k) == 1:
            # this takes care k==1 and n%k!=0 (some cases, n=8 k=6 will be as in below)
            rotateOneCycle(0, n)
        else:
            cycles = math.gcd(n, k)
            for c in range(cycles):
                rotateOneCycle(c, n//cycles)
        
to
        cycles = math.gcd(n, k)
        for c in range(cycles):
            rotateOneCycle(c, n//cycles)

Runtime: 421 ms, faster than 33.03% of Python3 online submissions for Rotate Array.
Memory Usage: 25.5 MB, less than 27.87% of Python3 online submissions for Rotate Array.

Runtime: 350 ms, faster than 51.92% of Python3 online submissions for Rotate Array.
Memory Usage: 25.3 MB, less than 75.31% of Python3 online submissions for Rotate Array.

I didn't use a lot of memory.. but speed is not fastest
I think there might be another solution..

the rotate by reverse is pure evil but genius

"""


class Solution:
    def rotate(self, nums, k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        def reverse(lst, l, r):
            while l < r:
                lst[l], lst[r] = lst[r], lst[l]
                l, r = l+1, r-1

        k %= len(nums)
        reverse(nums, 0, len(nums)-1)
        reverse(nums, 0, k-1)
        reverse(nums, k, len(nums)-1)


"""
Runtime: 383 ms, faster than 42.19% of Python3 online submissions for Rotate Array.
Memory Usage: 25.4 MB, less than 75.31% of Python3 online submissions for Rotate Array.
"""


if __name__ == '__main__':
    s = Solution()
    l = [i for i in range(1, 8)]
    s.rotate(l, 3)
    print(l)
