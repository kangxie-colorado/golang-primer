from collections import Counter
from typing import List


class Solution:
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        # a greedy idea.. not sure..
        # at most 6 buckets for each arrary

        def canMerge(src, dest, target, need):
            if need == 0:
                return True
            for i, j in zip(src, dest):
                if i == target and j != target:
                    need -= 1
                if need <= 0:
                    return True
            return False

        topBuckets = sorted(list([(v, k) for k, v in Counter(tops).items()]))
        bottomBuckets = sorted(
            list([(v, k) for k, v in Counter(bottoms).items()]))

        i, j = len(topBuckets)-1, len(bottomBuckets)-1
        while i >= 0 or j >= 0:
            if i >= 0 and j >= 0:
                if topBuckets[i][0] > bottomBuckets[j][0]:
                    # try to flip i from bottoms
                    # by merging
                    if canMerge(bottoms, tops, topBuckets[i][1], len(tops)-topBuckets[i][0]):
                        return len(tops)-topBuckets[i][0]
                    i -= 1
                else:
                    if canMerge(tops, bottoms, bottomBuckets[j][1], len(tops)-bottomBuckets[j][0]):
                        return len(tops)-bottomBuckets[j][0]
                    j -= 1
            else:
                if i >= 0:
                    if canMerge(bottoms, tops, topBuckets[i][1], len(tops)-topBuckets[i][0]):
                        return len(tops)-topBuckets[i][0]
                    i -= 1
                else:
                    if canMerge(tops, bottoms, bottomBuckets[j][1], len(tops)-bottomBuckets[j][0]):
                        return len(tops)-bottomBuckets[j][0]
                    j -= 1
        return -1


"""
okay.. code is ugly..
passed 
Runtime: 2820 ms, faster than 9.55% of Python3 online submissions for Minimum Domino Rotations For Equal Row.
Memory Usage: 15.2 MB, less than 20.94% of Python3 online submissions for Minimum Domino Rotations For Equal Row.

but not good in any way..
let me see how to make it better


Runtime: 1466 ms, faster than 68.01% of Python3 online submissions for Minimum Domino Rotations For Equal Row.
Memory Usage: 15.2 MB, less than 20.94% of Python3 online submissions for Minimum Domino Rotations For Equal Row.
"""


class Solution:
    def minDominoRotations(self, tops: List[int], bottoms: List[int]) -> int:
        # a greedy idea.. not sure..
        # at most 6 buckets for each arrary

        def canMerge(src, dest, target, need):
            # edge case here: need=0..
            #
            flips = 0
            for i, j in zip(src, dest):
                if i == target and j != target:
                    flips += 1
                    if flips >= need:
                        break

            return flips >= need

        topBuckets = sorted(list([(v, k) for k, v in Counter(tops).items()]))
        bottomBuckets = sorted(
            list([(v, k) for k, v in Counter(bottoms).items()]))

        i, j = len(topBuckets)-1, len(bottomBuckets)-1
        while i >= 0 and j >= 0:
            if topBuckets[i][0] > bottomBuckets[j][0]:
                # try to flip i from bottoms
                # by merging
                if canMerge(bottoms, tops, topBuckets[i][1], len(tops)-topBuckets[i][0]):
                    return len(tops)-topBuckets[i][0]
                i -= 1
            else:
                if canMerge(tops, bottoms, bottomBuckets[j][1], len(tops)-bottomBuckets[j][0]):
                    return len(tops)-bottomBuckets[j][0]
                j -= 1
        return -1


"""
Runtime: 1107 ms, faster than 96.31% of Python3 online submissions for Minimum Domino Rotations For Equal Row.
Memory Usage: 15 MB, less than 93.13% of Python3 online submissions for Minimum Domino Rotations For Equal Row.

this code gets rid of the one bucket running out scenario
I felt it but cannot put it into words...

just gave a try and it saves lot of time
thinking about 
I have exhausted the possible combinations for sure... 

okay... after reading Lee's post.. I think I actually over think it..
take the break and come back to this later..
"""


if __name__ == "__main__":
    s = Solution()
    tops = [1, 1, 1, 1, 1, 1, 1, 1]
    bottoms = [1, 1, 1, 1, 1, 1, 1, 1]
    print(s.minDominoRotations(tops, bottoms))

    tops = [3, 5, 1, 2, 3]
    bottoms = [3, 6, 3, 3, 4]
    print(s.minDominoRotations(tops, bottoms))

    tops = [2, 1, 2, 4, 2, 2]
    bottoms = [5, 2, 6, 2, 3, 2]

    print(s.minDominoRotations(tops, bottoms))
