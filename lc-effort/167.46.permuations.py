"""
https://leetcode.com/problems/permutations/

the first idea is backtracking..
or course there is this next-permutation

then this swap driving 
I'll just do backtracking here
"""


from typing import List


class Solution:
    def permute(self, nums: List[int]) -> List[List[int]]:
        res = []

        def backtracking(run, idx):
            if idx == len(nums):
                res.append(run.copy())
                return

            for n in nums:
                if n not in run:
                    run.append(n)
                    backtracking(run, idx+1)
                    run.pop()

        backtracking([], 0)
        return res


"""
Runtime: 74 ms, faster than 24.92% of Python3 online submissions for Permutations.
Memory Usage: 14.1 MB, less than 57.64% of Python3 online submissions for Permutations.
"""


if __name__ == '__main__':
    s = Solution()
    print(s.permute([1, 2, 3, 4, 5, 6]))
