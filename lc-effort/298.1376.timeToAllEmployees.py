"""
https://leetcode.com/problems/time-needed-to-inform-all-employees/

solved it using top-down
trying to solve it bottom-up.. but not recursive

"""


from typing import List


class Solution:
    def numOfMinutes(self, n: int, headID: int, manager: List[int], informTime: List[int]) -> int:
        # try dp bottom up
        timeToMe = [0]*n
        timeToMe[headID] = 0
        res = 0

        for empId, mgr in enumerate(manager):
            if empId == headID:
                continue
            stack = [empId]
            upMgr = mgr
            while upMgr != headID and timeToMe[upMgr] == 0:
                stack.append(upMgr)
                upMgr = manager[upMgr]

            time = timeToMe[upMgr]
            while stack:
                emp = stack.pop()
                time += informTime[manager[emp]]
                timeToMe[emp] = time

            res = max(res, timeToMe[empId])

        return res


"""
Runtime: 3434 ms, faster than 30.86% of Python3 online submissions for Time Needed to Inform All Employees.
Memory Usage: 30.7 MB, less than 98.65% of Python3 online submissions for Time Needed to Inform All Employees.

alright...
"""


if __name__ == '__main__':
    s = Solution()
    print(s.numOfMinutes(n=6, headID=2, manager=[
          2, 2, -1, 2, 2, 2], informTime=[0, 0, 1, 0, 0, 0]))
