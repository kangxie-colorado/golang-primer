'''
https://leetcode.com/problems/sqrtx/

I use that binary search..
'''


class Solution:
    def mySqrt(self, x: int) -> int:
        left, right = 0, x

        while left < right:
            mid = (left+right)/2
            if mid*mid - x > 0.0001:
                # mid too big
                right = mid
            elif mid*mid - x < 0:
                # mid too small
                left = mid
            else:
                return int(mid)


"""
12321 should be 111.. 
but I got 110

I think the left condition should be 0
I should error on the right side..
            if mid*mid - x > 0.0001:
                ...
            elif mid*mid - x < 0:
                ...
        so I accept mid which produce a diff between mid*mid and x in [0,0.0001]

but this ugliness does bother me..
however I think this guy's code
public int sqrt(int x) {
    if (x == 0)
        return 0;
    int left = 1, right = Integer.MAX_VALUE;
    while (true) {
        int mid = left + (right - left)/2;
        if (mid > x/mid) {
            right = mid - 1;
        } else {
            if (mid + 1 > x/(mid + 1))
                return mid;
            left = mid + 1;
        }
    }
}

left, right move by 1.. will it be right?
let me see
"""


class Solution:
    def mySqrt(self, x: int) -> int:
        if x == 0:
            return 0
        left, right = 1, x
        while True:
            mid = left + (right-left)//2
            if mid > x//mid:
                right = mid-1
            else:
                if mid+1 > x//(mid+1):
                    return int(mid)
                left = mid + 1


"""
Runtime: 49 ms, faster than 76.03% of Python3 online submissions for Sqrt(x).
Memory Usage: 13.8 MB, less than 57.11% of Python3 online submissions for Sqrt(x).

okay.. I see what it is doing
it is trying to find 
                if mid+1 > x/(mid+1):
                    return int(mid)
it is also focuses on integer only so left=1, starts with 1
also all the mid and comparison compuation do int calculations and thus left/right move by 1..

the key is to find mid*mid and (mid+1)*(mid+1)

this is smart.. mine is just a bit ...

"""
if __name__ == '__main__':
    print(Solution().mySqrt(12321))
