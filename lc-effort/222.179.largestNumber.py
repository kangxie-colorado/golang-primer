"""
https://leetcode.com/problems/largest-number/


"""


from typing import List


def cmp_to_key(mycmp):
    'Convert a cmp= function into a key= function'
    class K:
        def __init__(self, obj, *args):
            self.obj = obj

        def __lt__(self, other):
            return mycmp(self.obj, other.obj) < 0

        def __gt__(self, other):
            return mycmp(self.obj, other.obj) > 0

        def __eq__(self, other):
            return mycmp(self.obj, other.obj) == 0

        def __le__(self, other):
            return mycmp(self.obj, other.obj) <= 0

        def __ge__(self, other):
            return mycmp(self.obj, other.obj) >= 0

        def __ne__(self, other):
            return mycmp(self.obj, other.obj) != 0
    return K


class Solution:
    def largestNumber(self, nums:   List[int]) -> str:
        def compare(x, y):
            x = str(x)
            y = str(y)

            if x+y > y+x:
                return -1
            elif x+y == y+x:
                return 0
            else:
                return 1

        nums = sorted(nums, key=cmp_to_key(compare))
        if nums[0] == 0:
            return "0"
        return "".join([str(n) for n in nums])


"""
Runtime: 59 ms, faster than 63.36% of Python3 online submissions for Largest Number.
Memory Usage: 14.2 MB, less than 22.41% of Python3 online submissions for Largest Number.

Runtime: 52 ms, faster than 75.86% of Python3 online submissions for Largest Number.
Memory Usage: 14.2 MB, less than 22.41% of Python3 online submissions for Largest Number.

"""


if __name__ == '__main__':
    s = Solution()
    print(s.largestNumber(nums=[10, 2]))
