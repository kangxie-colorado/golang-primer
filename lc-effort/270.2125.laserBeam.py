"""
https://leetcode.com/problems/number-of-laser-beams-in-a-bank/


"""


from collections import Counter
from typing import List


class Solution:
    def numberOfBeams(self, bank: List[str]) -> int:
        res = i = 0
        while i < len(bank):
            count = Counter(bank[i])["1"]
            if count == 0:
                i += 1
                continue

            j = i+1
            nextRowCount = 0
            while j < len(bank):
                nextRowCount = Counter(bank[j])["1"]
                if nextRowCount != 0:
                    break
                j += 1

            res += count*nextRowCount
            i = j

        return res


if __name__ == "__main__":
    s = Solution()
    print(s.numberOfBeams(bank=["011001", "000000", "010100", "001000"]))
