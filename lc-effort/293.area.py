from collections import defaultdict
from typing import List


class Solution:
    def minAreaRect(self, points: List[List[int]]) -> int:

        xMap = defaultdict(set)
        yMap = defaultdict(set)

        xCandidates = set()
        yCandidates = set()

        for x, y in points:
            xMap[x].add(y)
            if len(xMap[x]) >= 2:
                xCandidates.add(x)
            yMap[y].add(x)
            if len(yMap[y]) >= 2:
                yCandidates.add(y)

        res = 5*10**9
        xCandidates = list(xCandidates)
        for i in range(len(xCandidates)):
            for j in range(i+1, len(xCandidates)):
                x1, x2 = xCandidates[i], xCandidates[j]
                ySet1 = xMap[x1]
                ySet2 = xMap[x2]
                interSet = list(ySet1.intersection(ySet2))
                for k1 in range(len(interSet)):
                    for k2 in range(k1+1, len(interSet)):
                        y1, y2 = interSet[k1], interSet[k2]
                        res = min(res, abs(y1-y2)*abs(x1-x2))

        if res == 5*10**9:
            return 0
        return res


if __name__ == '__main__':
    s = Solution()
    print(s.minAreaRect([[0, 1], [1, 3], [3, 3], [
          4, 4], [1, 4], [2, 3], [1, 0], [3, 4]]))
