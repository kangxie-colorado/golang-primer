"""
https://leetcode.com/problems/split-array-into-consecutive-subsequences/

I stared at this 
trialed a few cases
and blank ideas......

admit when you are beaten..
I was thinking of greedily extending the subsequence but I always got tangled with what should be condition to 
continue or stop... anyway, never a clear path was formed.

then reading Lee's code
https://leetcode.com/problems/split-array-into-consecutive-subsequences/discuss/106514/C%2B%2BPython-Esay-Understand-Solution

aha... yeah, just go one by one.. 
but always take 3 if it is a new start, otherwise, keep extending and maintain the states

now I should be able to code it up
"""


from collections import defaultdict
from email.policy import default
from re import T
from typing import Counter, List


class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        unused = Counter(nums)
        seqsEndedWith = defaultdict(int)

        for n in nums:
            if not unused[n]:
                continue

            if seqsEndedWith[n-1]:
                # n-1 is some sequnce's end,
                # continue to extend one of it
                seqsEndedWith[n-1] -= 1
                seqsEndedWith[n] += 1
                unused[n] -= 1
            elif unused[n+1] and unused[n+2]:
                # n-1 is not there, start a new one
                # if n+1/n+2 both are available
                unused[n] -= 1
                unused[n+1] -= 1
                unused[n+2] -= 1
                seqsEndedWith[n+2] += 1
            else:
                return False
        return True


"""
Runtime: 696 ms, faster than 70.23% of Python3 online submissions for Split Array into Consecutive Subsequences.
Memory Usage: 15.4 MB, less than 7.27% of Python3 online submissions for Split Array into Consecutive Subsequences.

from now on, unless what? don't push self too hard
when it is appropriate to look at answer, look at it...

also what is damnful is the problem is listed as a heap queue problem but where the fuck is the heap....????
"""

if __name__ == '__main__':
    s = Solution()
    assert s.isPossible([1, 2, 3, 3, 4, 5])
