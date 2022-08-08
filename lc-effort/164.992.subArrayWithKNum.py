"""
https://leetcode.com/problems/subarrays-with-k-different-integers/

wow..
only by thinking of it I feel stuck..

then Lee came out to say, you need to change the problem to solve at-most K
and then exactly-K = at-most-K - at-most(K-1)

what can I say!!!
I actually came here from https://leetcode.com/problems/count-number-of-nice-subarrays/

which I thought I figured out two solutions
1. use a queue and sliding window
2. odd number idx and just simple math..

then I see this one in the discussion.. what a fucking god!
so let me try to solve this at most k

notice that
[1,2,1,2,2,3] k=2

when 1, it is at-most-2, so +1
when sliding to 2, it is at-most-2, so +?... should +2, not just 1
    why? because [1,2] will be one for sure
                and [2] itself will be one for sure too
                and kind of see why at most is much more solvable than exactly k
                for exact-2, [1,2,1] vs [2,1,1] will melt your brain
                but for at-most-2, they can be easily solved (easier)

pray and bless, let me try

"""


from collections import defaultdict
from typing import List


class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:

        def atMost(t):
            i, j = 0, 0
            count = 0
            numMap = defaultdict(int)
            while j < len(nums):
                numMap[nums[j]] += 1
                while len(numMap) > t:
                    numMap[nums[i]] -= 1
                    if numMap[nums[i]] == 0:
                        numMap.pop(nums[i])
                    i += 1

                count += j-i+1
                j += 1

            return count

        return atMost(k) - atMost(k-1)


"""
Runtime: 763 ms, faster than 47.26% of Python3 online submissions for Subarrays with K Different Integers.
Memory Usage: 16.7 MB, less than 53.89% of Python3 online submissions for Subarrays with K Different Integers.

now I go back to the medium and rewrite this...
"""


"""
https://leetcode.com/problems/subarrays-with-k-different-integers/discuss/235235/C%2B%2BJava-with-picture-prefixed-sliding-window

reading this.. it is brilliant but so many steps to follow.
what is the key is I think the while() loop...
int subarraysWithKDistinct(vector<int>& A, int K, int res = 0) {
  vector<int> m(A.size() + 1);
  for(auto i = 0, j = 0, prefix = 0, cnt = 0; i < A.size(); ++i) {
    if (m[A[i]]++ == 0) ++cnt;
    if (cnt > K) --m[A[j++]], --cnt, prefix = 0;
    // ^ because of the while() loop, this will ensure taking off an uniq number
    while (m[A[j]] > 1) ++prefix, --m[A[j++]];      
    // ^ this keeps the left elements unqiue to the subarray and turn duplicates into prefix
    if (cnt == K) res += prefix + 1;
  }
  return res;
}
"""

if __name__ == '__main__':
    s = Solution()
    print(s.subarraysWithKDistinct([1, 2, 1, 2, 3], 2))
