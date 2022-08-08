"""
https://leetcode.com/problems/combination-sum-ii/

backtracking?
"""


from typing import Counter, List


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        results = set()
        candidates.sort()

        def backtracking(i, run, S):
            if S == target:
                results.add(tuple(run))
                return

            if S > target:
                return

            for j in range(i+1, len(candidates)):
                run.append(candidates[j])
                backtracking(j, run, S+candidates[j])
                run.pop()

        for i in range(len(candidates)):
            if i > 0 and candidates[i] == candidates[i-1]:
                continue
            backtracking(i, [candidates[i]], candidates[i])

        res = []
        for r in results:
            res.append(list(r))
        return res


"""
will work.. but
[[1, 2, 5], [1, 7], [1, 6, 1], [2, 6], [2, 1, 5], [7, 1]]

duplicated..
maybe sort first
still
[[1, 1, 6], [1, 2, 5], [1, 7], [1, 2, 5], [1, 7], [2, 6]]

then add this
                if j > 0 and candidates[j] == candidates[j-1] and len(run) < j:
                    continue
        reject using the duplicated number at the "top" level (len(run) < j)?
        not sure if this is right
not right..
failed here

[3,1,3,5,1,1]
8
Output
[[1,1,1,5],[3,5]]
Expected
[[1,1,1,5],[1,1,3,3],[3,5]]

used set() tuple to deduplicate and failed here
[1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1]
30
>>> len(a)
100

so I really to dive into change to hashmap mode
for each key...

and then it turns into a recursive sub-problem..
"""


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        C = Counter(candidates)

        def helper(t, used):
            if t == 0:
                return [[]]
            combs = []
            for k, v in C.items():
                if k in used:
                    continue
                for i in range(1, min(v, t//k)+1):
                    left = [k]*i
                    used.add(k)
                    for r in helper(t-k*i, used.copy()):
                        combs.append(left + r)

            return combs
        return helper(target, set())


"""
Runtime: 101 ms, faster than 63.86% of Python3 online submissions for Combination Sum II.
Memory Usage: 13.9 MB, less than 93.22% of Python3 online submissions for Combination Sum II.

Runtime: 82 ms, faster than 78.14% of Python3 online submissions for Combination Sum II.
Memory Usage: 13.8 MB, less than 93.22% of Python3 online submissions for Combination Sum II.

the most tricky part is to deal with duplicates..
first solution working on the straight array.. it will easily duplicate and timeout
2nd move the battle to work on the hash keys..

then there is some critial point
    for i in range(1, min(v, t//k)+1):
                   ^ # this must start with 1, not 0

if you start with 0, you are introducing duplicated..
because it will become a DFS..
for example
[1 1 1 2] 2
it will use 0 1s, and get [2]
then it use 2 1s..

then it move on to use 1 [2].. and that is duplicate

but since the DFS also finds the combinations..
can I modify this a bit to use that too
let me try
huh... very hard...

let me learn now
hmm... honestly I don't see a very elegant way
most of them use backtracking but dealt with the duplicates in some way, 6 paramters in the function..
very hard to understand..

I'll accept my solution

let me try to crack up that DP solution
terrible...

full of duplicates...
"""


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        def isLegit(m1, m2):
            for k, v in m2.items():
                if v > m1[k]:
                    return False
            return True

        dp = [[]]*(target+1)
        for i in range(target+1):
            dp[i] = []

        C = Counter(candidates)
        for c in C:
            if c <= target:
                dp[c].append([c])

        for i in range(1, target+1):
            for j in range(1, i//2+1):
                if dp[j] and dp[i-j] and (len(dp[j]) == 1 or len(dp[i-j]) == 1):
                    for l in dp[j]:
                        for r in dp[i-j]:
                            if isLegit(C, Counter(l+r)):
                                dp[i].append(l+r)

        return dp[target]


"""
with something learned from 4sum problem..
let me revisit the backtracking
"""


class Solution:
    def combinationSum2(self, candidates: List[int], target: int) -> List[List[int]]:
        candidates.sort()

        res = []

        def backtracking(i, run, S):
            if S == target:
                res.append(run.copy())
                return

            if S > target:
                return

            for j in range(i, len(candidates)):
                if j > i and candidates[j] == candidates[j-1]:
                    continue
                run.append(candidates[j])
                backtracking(j+1, run, S+candidates[j])
                run.pop()

        backtracking(0, [], 0)
        return res


"""
Runtime: 95 ms, faster than 68.09% of Python3 online submissions for Combination Sum II.
Memory Usage: 14 MB, less than 58.94% of Python3 online submissions for Combination Sum II.

the key to de-duplication is
                if j > i and candidates[j] == candidates[j-1]:
                    continue
                because i is the first desinated usage of this element... 
                on the same layer.. I shall not re-used
                but on the next layer, i should be a different value (usually +1)
"""

if __name__ == '__main__':
    print(Solution().combinationSum2([1, 1, 1, 2], 2))
    print(Solution().combinationSum2([1, 1, 1, 3, 3, 5], 8))
    print(Solution().combinationSum2([10, 1, 2, 7, 6, 1, 5], 8))
    print(Solution().combinationSum2([2, 5, 2, 1, 2], 5))
