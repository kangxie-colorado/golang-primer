"""
https://leetcode.com/problems/count-unhappy-friends/

"""


from typing import List


class Solution:
    def unhappyFriends(self, n: int, A: List[List[int]], pairs: List[List[int]]) -> int:

        pairMap = {}
        for a, b in pairs:
            pairMap[a] = b
            pairMap[b] = a
        res = 0
        for a, b in pairMap.items():

            # a paired with b
            # check how many more-preferred ahead of b
            morePrefers = []
            for p in A[a]:
                if p == b:
                    break
                morePrefers.append(p)

            # check more-preferred for their prefers

            for mp in morePrefers:
                for prefer in A[mp]:
                    if prefer == pairMap[mp]:
                        break
                    if prefer == a:
                        res += 2
                        break

        return res//2


if __name__ == "__main__":
    s = Solution()

    print(s.unhappyFriends(n=4, A=[[1, 3, 2], [2, 3, 0], [
          1, 3, 0], [0, 2, 1]], pairs=[[1, 3], [0, 2]]))
    print(s.unhappyFriends(n=4,
                           A=[
                               [1, 2, 3], [3, 2, 0], [3, 1, 0], [1, 2, 0]],
                           pairs=[[0, 1], [2, 3]]))
