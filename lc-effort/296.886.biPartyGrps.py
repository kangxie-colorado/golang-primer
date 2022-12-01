"""
https://leetcode.com/problems/possible-bipartition/

1. tried union-find like technique --  wrong
2. tried back-track like  -- wrong

now I think it is a dfs problem
"""


from collections import defaultdict
from typing import List


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        grpA = set()
        grpB = set()

        adjSets = defaultdict(set)
        for a, b in dislikes:
            adjSets[a].add(b)
            adjSets[b].add(a)

        grps = [grpA, grpB]

        def dfs(a, grpChoice):
            if a in grps[grpChoice]:
                return True
            grps[grpChoice].add(a)
            grpChoice ^= 1
            if a in grps[grpChoice]:
                return False
            for b in adjSets[a]:
                if not dfs(b, grpChoice):
                    return False

            return True

        for a, b in dislikes:
            if a not in grpA and a not in grpB:
                if not dfs(a, 0):
                    return False
        return True


"""
as a dfs
uni-directional, wrong
bi-directional two groups... right

now this is a bi-directional graph, it just become a circle detection
"""


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        adjSets = defaultdict(set)
        for a, b in dislikes:
            adjSets[a].add(b)
            adjSets[b].add(a)

        visited = set()

        # wrong... wrong... wrong
        # let me try to make it right
        # okay.. think I fixed the cycle detection code but this problem is no longer a cycle detection issue
        # wait.. not really..
        # al.. okay.. this logical is not right
        # 1 -- 2... this is a cycle? no, but it is kind of a cycle... because 1 to 2 and 2 to 1
        # so cycle detection is kind of a uni-directional thing???????
        # okay.. that is arguable
        # but let me brush up the cycle detection algorithm since I am already here
        # https://www.geeksforgeeks.org/detect-cycle-undirected-graph/
        # key is using a parent to prevent back-flip
        # Else if the adjacent node is visited and not the parent of the current node then return true.
        # let me try fix the cycle detection code
        """
            at least correct answers to these questions: False/True/True
            print(s.possibleBipartition(
                n=4, dislikes=[[1, 2], [1, 3], [2, 4], [3, 4]]))
            print(s.possibleBipartition(n=4, dislikes=[[1, 2], [1, 3], [2, 4]]))
            print(s.possibleBipartition(10,
                                        [[5, 9], [5, 10], [5, 6], [5, 7], [1, 5], [4, 5], [2, 5], [5, 8], [3, 5]]))
        """

        def cycle(a, parent, path):
            if a in path:
                return True

            path.add(a)
            for b in adjSets[a]:
                if b != parent and cycle(b, a, path):
                    return True
            visited.add(a)
            return False

        for a in adjSets:
            if a not in visited and cycle(a, -1, set()):
                return False
        return True


"""
okay stop this... it is not cycle detection either... let alone the code is wrong anyway
  2
 / \
1   4
 \ /
  3

think this... it has a cycle.. but can be grouped as [1,4] and [2,3]

but now I learned the trick to use a parent to prevent back-flip
let me improve the code...

hmm... still not right.. (code deleted)
the code I used to detection is right... just ugly

okay.. I remember this is actually the coloring problem??

"""


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        colors = [0]*(n+1)

        adjSets = defaultdict(set)
        for a, b in dislikes:
            adjSets[a].add(b)
            adjSets[b].add(a)

        def dfs(a, color):
            if colors[a] != 0:
                return colors[a] == color

            colors[a] = color
            for b in adjSets[a]:
                if not dfs(b, -color):
                    return False
            return True

        for a in adjSets:
            if colors[a] == 0 and not dfs(a, 1):
                return False

        return True


"""
dfs can work, can I bfs
"""


class Solution:
    def possibleBipartition(self, n: int, dislikes: List[List[int]]) -> bool:
        colors = [0]*(n+1)

        adjSets = defaultdict(set)
        for a, b in dislikes:
            adjSets[a].add(b)
            adjSets[b].add(a)

        def bfs(a, color):
            q = [(a, color)]
            while q:
                a, color = q[0]
                q = q[1:]
                if colors[a] == -color:
                    return False
                if colors[a] == color:
                    continue
                colors[a] = color
                for b in adjSets[a]:
                    q.append((b, -color))
            return True

        for a in adjSets:
            if colors[a] == 0 and not bfs(a, 1):
                return False

        return True


"""
Runtime: 7721 ms, faster than 5.01% of Python3 online submissions for Possible Bipartition.
Memory Usage: 20.9 MB, less than 47.38% of Python3 online submissions for Possible Bipartition.

okay.. apparently the dfs is faster in detection such confliction 
that is understandablewe
"""


if __name__ == "__main__":
    s = Solution()
    print(s.possibleBipartition(
        n=4, dislikes=[[1, 2], [1, 3], [2, 4], [3, 4]]))
    print(s.possibleBipartition(n=4, dislikes=[[1, 2], [1, 3], [2, 4]]))
    print(s.possibleBipartition(10,
                                [[5, 9], [5, 10], [5, 6], [5, 7], [1, 5], [4, 5], [2, 5], [5, 8], [3, 5]]))
