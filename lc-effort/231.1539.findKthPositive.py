"""
https://leetcode.com/problems/kth-missing-positive-number/

O(n) is trivial..
O(lgn) is interesting

so whatever that is >arr[-1] will be just missing...
so I know A[-1] and know len(A) is
len(A) - A[-1] is what is missed in the array..
if k > len(A) - A[-1], there is an easy answer


        # now binary search to find what have been missing inside array
        # and looks like this is binary search on the range?
        [2,3,4,7,11]
        m = 4.. [2,3,4], missing 1
        then it has to be to right

        [5... 7,11] missing 5 numbers(5,6,8,9,10)..
        how to express this??

        feel like recursive is kind of more natural to me
        and the above conditino is actually the base case
"""


from typing import List


class Solution:
    def findKthPositive(self, arr: List[int], k: int) -> int:

        def helper(A, base, remain):
            if len(A) == 0 or remain > A[-1] - len(A) - base:
                # deal with out of array misses or empty array
                return remain + len(A) + base

            m = len(A)//2
            if A[m] - base - m-1 < remain:
                # what is A[m] - base - m-1
                # A[m] - base is the ought-to-be full length
                # m-1 is the half array's length (actually -(m+1))
                # not in this half
                return helper(A[m+1:], A[m], remain-(A[m] - base - m-1))
            else:
                # why A[:m] not A[:m+1], somehow a bit weird
                # of couse, A[:m+1] will cause dead loop
                # but if between A[m-1] and A[m] we also have gap..
                # do I need to care.. maybe not
                # it will be take care by the base case??
                # [1,2,4,5,6] k=1.. m=idx-2 (4)
                # A[:m] will be [1,2] remain is 1.. it will be
                # remain(1) + len(A)(2) + base(0) = 3...
                # huh...

                return helper(A[:m], base, remain)

        return helper(arr, 0, k)


"""
[8,11,16,20,29,30,32,33,37,39,42,44,46,47,48,50,52,56,60,63,64,65,68,70,72,74,80]
45

expect 67 I got 65...
okay.. forgot to minus base on the m
            if A[m] - base - m-1 < remain:
                # not in this half
                return helper(A[m+1:], A[m], remain-(A[m] - base - m-1))
"""

"""
Runtime: 59 ms, faster than 86.32% of Python3 online submissions for Kth Missing Positive Number.
Memory Usage: 13.9 MB, less than 85.85% of Python3 online submissions for Kth Missing Positive Number.
"""

"""
https://leetcode.com/problems/kth-missing-positive-number/discuss/779999/JavaC%2B%2BPython-O(logN)

Lee's is looking at again in reverval
I was focus on how many missing
he is focusing on not-missing... 

not fully understood
but this first comment is also genius 

    def findKthPositive(self, A, k):
        class Count(object):
            def __getitem__(self, i):
                return A[i] - i - 1
        return k + bisect.bisect_left(Count(), k, 0, len(A))

    the Count class, transform the array into missing-so-far
    [1,3,4,6] to
    [0,1,1,2] at i-th spot, missing B[i] elemenet

    bisect_left will get the insert index, i.e. the first pos that is >= k or the end..
    but it is >=.. so there could be more
    so look back at one position l-1.. then add how many?
    think at A[l-1] it misses B[l-1]
    so add A[l-1] + (k- B[l-1]) = A[l-1] + (k- (A[l-1] -(l-1)-1)) ==> k+l...ahaha...
    this is kind of like what I did for base case

    also returning out of range is kind of like returning +infinite.. 
    so looking back one position

    now people ask what if l=0... you cannot look at idx -1..
    well this is math vs program.. what can I say
    they are not exactly the same.. match checks out and hide that bit

okay.. so my solution was to search biggest range that cannot provide K missing number

his is to search smallest range that CAN provide K missing number.. 
if over the boundary treat it like infinite? 
whatever, then look back one position and get the results... 
let me code it up


"""


class Solution:
    def findKthPositive(self, arr, k):
        A = arr
        l, r = 0, len(A)
        while l < r:
            m = l+(r-l)//2
            if A[m]-m-1 < k:
                l = m+1
            else:
                r = m
        return l+k


"""
Runtime: 43 ms, faster than 99.50% of Python3 online submissions for Kth Missing Positive Number.
Memory Usage: 14 MB, less than 85.85% of Python3 online submissions for Kth Missing Positive Number.

wow....
"""

if __name__ == "__main__":
    s = Solution()
    assert 67 == s.findKthPositive([8, 11, 16, 20, 29, 30, 32, 33, 37, 39, 42,
                                    44, 46, 47, 48, 50, 52, 56, 60, 63, 64, 65, 68, 70, 72, 74, 80], 45)  # 67

    assert 3 == s.findKthPositive(arr=[1, 2,  4, 7, 11], k=1)  # 3

    assert 1 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=1)  # 1
    assert 6 == s.findKthPositive(arr=[1, 2, 3, 4], k=2)
    assert 9 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=5)
    assert 10 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=6)
    assert 12 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=7)
    assert 5 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=2)
    assert 6 == s.findKthPositive(arr=[2, 3, 4, 7, 11], k=3)
    assert 5 == s.findKthPositive(arr=[1, 4, 7, 8], k=3)
