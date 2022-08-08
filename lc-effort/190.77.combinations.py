"""
https://leetcode.com/problems/combinations/

I see the backtrcking.. 
and I ran into that duplicate scenarios..

but then I figured out for next i, start with i+1.. don't look back
"""


from typing import List


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def backtracking(last, r):
            if len(r) == k:
                res.append(list(r))

            for i in range(last+1, n+1):
                if i in r:
                    continue
                r.add(i)
                backtracking(i, r)
                r.remove(i)

        backtracking(0, set())
        return res


"""
Runtime: 804 ms, faster than 22.03% of Python3 online submissions for Combinations.
Memory Usage: 16.1 MB, less than 20.22% of Python3 online submissions for Combinations.

very obviously there are two tiers of solutions..
let me see what is the better way

actually... because I am starting from last+1..
no need to test i in r
"""


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        res = []

        def backtracking(last, r):
            if len(r) == k:
                res.append(r.copy())

            for i in range(last+1, n+1):
                r.append(i)
                backtracking(i, r)
                r.pop()

        backtracking(0, [])
        return res


"""
Runtime: 718 ms, faster than 33.04% of Python3 online submissions for Combinations.
Memory Usage: 16 MB, less than 52.65% of Python3 online submissions for Combinations.

this is obviously a DP problem? or recursive problem? not that DP
"""


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        def getCombs(N, K):
            if K == 1:
                return [[i] for i in range(1, N+1)]
            res = []
            for i in range(N, K-1, -1):
                subCombs = getCombs(i-1, K-1)
                res.extend([[i]+j for j in subCombs])

            return res
        return getCombs(n, k)


"""
Runtime: 175 ms, faster than 80.42% of Python3 online submissions for Combinations.
Memory Usage: 16.4 MB, less than 6.51% of Python3 online submissions for Combinations.

why this is much quicker?
"""


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        m = {}

        def getCombs(N, K):
            nonlocal m
            if K == 1:
                return [[i] for i in range(1, N+1)]
            if (N, K) in m:
                return m[(N, K)]
            res = []
            for i in range(N, K-1, -1):
                subCombs = getCombs(i-1, K-1)
                res.extend([[i]+j for j in subCombs])
            m[(N, K)] = res
            return res
        return getCombs(n, k)


"""
Runtime: 148 ms, faster than 84.52% of Python3 online submissions for Combinations.
Memory Usage: 19.7 MB, less than 5.02% of Python3 online submissions for Combinations.

there should be an iterative ones
"""


class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        combs = [[i] for i in range(1, n+1)]
        for i in range(2, k+1):
            newCombs = []
            for c in combs:
                for j in range(c[-1]+1, n+1):
                    newCombs.append(c+[j])
            combs = newCombs
        return combs


"""
Runtime: 1563 ms, faster than 5.00% of Python3 online submissions for Combinations.
Memory Usage: 69 MB, less than 5.02% of Python3 online submissions for Combinations.

Runtime: 703 ms, faster than 35.05% of Python3 online submissions for Combinations.
Memory Usage: 69 MB, less than 5.02% of Python3 online submissions for Combinations.

wel...
"""


if __name__ == '__main__':
    print(Solution().combine(4, 2))
    print(Solution().combine(4, 1))
    print(Solution().combine(4, 3))
    print(Solution().combine(4, 4))
    print(Solution().combine(5, 3))
