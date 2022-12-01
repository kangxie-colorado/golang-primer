"""
https://leetcode.com/problems/longest-mountain-in-array/

didn't do this before
let me think

naive and intuitive is find the increasing seq then decreasing seq
then I also think taking a page from the palindrome.. 

start from any number.. go two ways.. see how far it decreasing..
but this will be O(n2)

maybe not.. it can jump to the front
"""

from typing import List


class Solution:
    def longestMountain(self, arr: List[int]) -> int:
        res = 0
        i = 1

        while i < len(arr)-1:

            j1, j2 = i, i-1
            k1, k2 = i, i+1

            isMountain = False
            while j2 >= 0 and arr[j1] > arr[j2] and k2 < len(arr) and arr[k1] > arr[k2]:
                j1, j2 = j1-1, j2-1
                k1, k2 = k1+1, k2+1
                isMountain = True

            if isMountain:
                while j2 >= 0 and arr[j1] > arr[j2]:
                    j1, j2 = j1-1, j2-1

                while k2 < len(arr) and arr[k1] > arr[k2]:
                    k1, k2 = k1+1, k2+1

                res = max(res, k1-j1+1)
            i = max(k1, i+1)
        return res


"""
Runtime: 453 ms, faster than 5.08% of Python3 online submissions for Longest Mountain in Array.
Memory Usage: 15.4 MB, less than 16.13% of Python3 online submissions for Longest Mountain in Array.
Next challenges:

okay. truly mirrable 
let me do that intuitive solution
"""


class Solution:
    def longestMountain(self, A: List[int]) -> int:
        res = 0
        i = 0

        while i < len(A)-1:
            j = i+1
            while j < len(A) and A[j] > A[j-1]:
                # increasing seq
                # end when A[j] <= A[j-1]
                j += 1
            if j == len(A):
                break

            if j > i+1:
                while j < len(A) and A[j] < A[j-1]:
                    # decreasing seq
                    # end when A[k] >= A[k-1]
                    j += 1
                    res = max(res, j-i)
            i = max(i+1, j-1)
        return res


if __name__ == "__main__":
    s = Solution()
    print(s.longestMountain([2, 1, 4, 7, 3, 2, 5]))
