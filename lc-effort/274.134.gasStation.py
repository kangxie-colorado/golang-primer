"""
https://leetcode.com/problems/gas-station/  

"""


from typing import List


class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        gains = [i-j for i, j in zip(gas, cost)]
        gains += gains

        i = 0
        while i <= len(gas):
            if gains[i] < 0:
                i += 1
                continue

            j = i+1
            tank = gains[i]
            while j < i+len(gas):
                tank += gains[j]
                if tank < 0:
                    break
                j += 1

            if j == i+len(gas):
                return i % len(gas)
            # i can jump to j... because it has been negative gain from i to j..
            # and i is positive.. that makes i+1 to j smaller... can skip
            i = j
        return -1


if __name__ == '__main__':
    s = Solution()
    print(s.canCompleteCircuit(gas=[2, 3, 4], cost=[3, 4, 3]))
