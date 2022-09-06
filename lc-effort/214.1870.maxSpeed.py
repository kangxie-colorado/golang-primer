"""
https://leetcode.com/problems/minimum-speed-to-arrive-on-time/

"""


from typing import List


class Solution:
    def minSpeedOnTime(self, dist: List[int], hour: float) -> int:
        def isOkay(m):
            hrs = 0.0
            for d in dist[:-1]:
                hrs += (d+m-1)//m
            hrs += dist[-1]/m
            return hrs <= hour
        l, r = 1, 10**7+1
        while l < r:
            m = l+(r-l)//2
            if isOkay(m):
                r = m
            else:
                l = m+1

        return l if l <= 10**7 else -1


"""
Runtime: 5443 ms, faster than 56.11% of Python3 online submissions for Minimum Speed to Arrive on Time.
Memory Usage: 28.1 MB, less than 73.07% of Python3 online submissions for Minimum Speed to Arrive on Time.
"""

if __name__ == '__main__':
    s = Solution()

    print(s.minSpeedOnTime([1, 1, 100000], 2.01))
    print(s.minSpeedOnTime([1, 3, 2], 6))
