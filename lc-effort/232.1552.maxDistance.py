"""
https://leetcode.com/problems/magnetic-force-between-two-balls/

if I don't come here via binary search.. I would be difficult to think it is a binary search problem
how the monotonic attribute is at when a m meets the condition, so does m' <=m...

but how to write the test() function
walk thru and see it can be arranged?
"""


from turtle import pos
from typing import List


class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        position.sort()

        def ok(dist):
            placed = 1
            lastPos = position[0]
            i = 1
            while True:
                while i < len(position) and position[i]-lastPos < dist:
                    i += 1
                if i == len(position):
                    break

                placed += 1
                if placed >= m:
                    return True
                lastPos = position[i]
            return False

        l, r = 1, position[-1] - position[0]

        while l < r:
            mid = r-(r-l)//2
            if ok(mid):
                l = mid
            else:
                r = mid-1
        return l


"""
best I can achieve
Runtime: 2271 ms, faster than 43.86% of Python3 online submissions for Magnetic Force Between Two Balls.
Memory Usage: 28.1 MB, less than 26.32% of Python3 online submissions for Magnetic Force Between Two Balls.
hmm.. what is the improvement

heap???

nah.. same complexity


https://leetcode.com/problems/magnetic-force-between-two-balls/discuss/794070/PythonC%2B%2B-Binary-search-solution-with-explanation-and-similar-questions
hmm indeed I can simply the whil eloop
"""


class Solution:
    def maxDistance(self, position: List[int], m: int) -> int:
        position.sort()

        def ok(dist):
            placed = 1
            lastPos = position[0]
            i = 1
            for i in range(1, len(position)):
                if position[i] - lastPos >= dist:
                    lastPos = position[i]
                    placed += 1

            return placed >= m

        l, r = 1, position[-1] - position[0]

        while l < r:
            mid = r-(r-l)//2
            if ok(mid):
                l = mid
            else:
                r = mid-1
        return l


""""
Runtime: 1499 ms, faster than 82.68% of Python3 online submissions for Magnetic Force Between Two Balls.
Memory Usage: 27.7 MB, less than 91.67% of Python3 online submissions for Magnetic Force Between Two Balls.

why?  so much different?
"""

if __name__ == "__main__":
    s = Solution()
    assert 6 == s.maxDistance(position=[1, 2, 3, 4, 7], m=2)
    assert 3 == s.maxDistance(position=[1, 2, 3, 4, 7], m=3)
    assert 999999999 == s.maxDistance(
        position=[5, 4, 3, 2, 1, 1000000000], m=2)
    assert 4 == s.maxDistance(
        position=[5, 4, 3, 2, 1, 1000000000], m=3)
    assert 2 == s.maxDistance(
        position=[5, 4, 3, 2, 1, 1000000000], m=4)
    assert 4 == s.maxDistance(
        position=[5, 4, 3, 2, 1, 1000, 1000000000], m=4)
