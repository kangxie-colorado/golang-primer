"""
https://leetcode.com/problems/max-consecutive-ones-iii/

I did this one before... but back then I don't know sliding window as now
so I struggled and awed at others' solutions.

now I know a bit better so I want to redo it.. since I come across it..
do a classical multiple times is the strategy right?

so this one, actually inspired by last night video's most frequent after K OPs..
I can use the similar policy for sliding/shrinking

when windlen <= sum+k, then I can grow
that mean while windLen > sum+k, shrink

yes...
"""


from typing import List


class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        maxLen = 0
        i, j = 0, 0
        sum = 0

        while j < len(nums):
            sum += nums[j]

            while j-i+1 > sum+k:
                sum -= nums[i]
                i += 1

            maxLen = max(maxLen, j-i+1)
            j += 1

        return maxLen


"""
Runtime: 1133 ms, faster than 21.55% of Python3 online submissions for Max Consecutive Ones III.
Memory Usage: 14.7 MB, less than 21.07% of Python3 online submissions for Max Consecutive Ones III.

okay... whatever...
there will be a non-shrinking solution too
"""


class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        i, j = 0, 0
        sum = 0

        while j < len(nums):
            sum += nums[j]

            if j-i+1 > sum+k:
                sum -= nums[i]
                i += 1

            j += 1

        return j-i


"""
Runtime: 1084 ms, faster than 26.77% of Python3 online submissions for Max Consecutive Ones III.
Memory Usage: 14.7 MB, less than 21.07% of Python3 online submissions for Max Consecutive Ones III.

this non-shrinking is counter intuitive and not as widely applicable as the shrinking solution
but it is smart and here
what it means is when there a j,i meeting the condition, it will stay that way.. it will only becomes bigger
if j grows more while i don't

then lee's code is 
    def longestOnes(self, A, K):
        i = 0
        for j in xrange(len(A)):
            K -= 1 - A[j]
            if K < 0:
                K += 1 - A[i]
                i += 1
        return j - i + 1


this guy, likes to do minus but you have to admit this is cleaner
so it focus on the zeros.. the cost of admitting one ZERO is 1-A[j], which is 1-0=1..
on the opposite, the cost of admitting a ONE is 1-1=0, so no cost...

also when shrinking, the gain from rid of a ZEOR or a ONE is the same...
and then of course, he is using non-shrinking..
let me change it to shrinking and it should still pass
class Solution:
    def longestOnes(self, A, K):
        i = 0
        maxLen=0
        for j in range(len(A)):
            K -= 1 - A[j]
            while K < 0:
                K += 1 - A[i]
                i += 1
            maxLen=max(maxLen,j-i+1)
        return maxLen

Runtime: 1042 ms, faster than 30.87% of Python3 online submissions for Max Consecutive Ones III.
Memory Usage: 14.7 MB, less than 60.58% of Python3 online submissions for Max Consecutive Ones III.

also notice my non-shrinking vs his, i returned j-i, he returned j-i+1
might be the xrange()??

nah.. I think my j can grow to len(nums)
his can only grow to len(nums)-1
"""

if __name__ == '__main__':
    s = Solution()
    print(s.longestOnes([1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0], 2))
    print(s.longestOnes([0, 0, 1, 1, 0, 0, 1, 1,
          1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1], 3))
