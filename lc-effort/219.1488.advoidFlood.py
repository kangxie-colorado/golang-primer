"""
https://leetcode.com/problems/avoid-flood-in-the-city/

I don't have an idea..
a backtracking.. maybe.. 
cannot pass but just to check my thinking
"""


from codecs import backslashreplace_errors
from collections import defaultdict
from typing import Counter, List


class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        res = [0]*len(rains)

        def backtrack(i, fullLakes):
            if i == len(rains):
                return True

            if rains[i] != 0:
                res[i] = -1
                if rains[i] in fullLakes:
                    return False
                fullLakes.add(rains[i])
                return backtrack(i+1, fullLakes.copy())  # need to return here
            else:
                if len(fullLakes):
                    for k in list(fullLakes):
                        fullLakes.remove(k)
                        res[i] = k
                        if backtrack(i+1, fullLakes.copy()):
                            # if one option ends up true, need to return here
                            return True
                        fullLakes.add(k)
                else:
                    res[i] = 1
                    return backtrack(i+1, fullLakes.copy())

        dried = backtrack(0, set())
        return res if dried else []


"""
failed here
[69,0,0,0,69]

okay.. copy the set() everywhere and the correctness may have been established
45 / 81 test cases passed.

but of course
Status: Time Limit Exceeded

>>> len(a)
228
"""


class Solution:
    def avoidFlood(self, rains: List[int]) -> List[int]:
        C = Counter(rains)
        for i, r in enumerate(rains):
            if r != 0 and C[r] == 1:
                rains[i] = -1

        res = [0]*len(rains)

        def backtrack(i, fullLakes):
            if i == len(rains):
                return True

            if rains[i] != 0:
                res[i] = -1
                if rains[i] in fullLakes:
                    return False
                if rains[i] != -1:
                    fullLakes.add(rains[i])
                return backtrack(i+1, fullLakes.copy())  # need to return here
            else:
                if len(fullLakes):
                    for k in list(fullLakes):
                        fullLakes.remove(k)
                        res[i] = k
                        if backtrack(i+1, fullLakes.copy()):
                            # if one option ends up true, need to return here
                            return True
                        fullLakes.add(k)
                else:
                    res[i] = 1
                    return backtrack(i+1, fullLakes.copy())

        dried = backtrack(0, set())
        return res if dried else []


"""
I briefly checked a couple discussions..
so the issue is I cannot quite figure out the theory for solving this.. hahhaha
so anyway.. just try understand them and code it myself, no need to stay stuck here...
"""

if __name__ == '__main__':
    s = Solution()

    a = [0, 38176, 0, 26125, 0, 0, 0, 0, 59074, 0, 63638, 51866, 73765, 96691, 71758, 46542, 0, 0, 0, 0, 27875, 0, 0, 30371, 68853, 0, 0, 27875, 46542, 59074, 0, 75794, 18368, 46542, 0, 0, 99196, 38176, 96691, 0, 0, 0, 52956, 0, 18368, 0, 0, 52956, 0, 0, 0, 38176, 68853, 0, 0, 0, 73765, 75794, 59074, 59074, 0, 0, 0, 0, 0, 0, 0, 51866, 94805, 0, 75325, 0, 75325, 82237, 0, 0, 51866, 59074, 71758, 51866, 0, 30371, 52956, 0, 75325, 96691, 0, 0, 0, 0, 0, 94805, 30371, 46168, 56780, 0, 70395, 40371, 0, 51866, 63638, 94805, 63638, 18368, 0, 96691, 75794, 0, 68853, 26125, 0, 71758, 14052,
         0, 59074, 75794, 30371, 0, 0, 63638, 71758, 75325, 13194, 18368, 0, 68853, 0, 13194, 52956, 0, 71758, 14052, 27875, 0, 70395, 0, 27875, 0, 52956, 73765, 18368, 75794, 0, 0, 0, 0, 52956, 70395, 0, 73765, 0, 0, 75325, 0, 73765, 70395, 0, 0, 75325, 0, 0, 70395, 68853, 0, 0, 0, 27875, 0, 0, 51866, 0, 18368, 0, 0, 0, 14052, 18368, 30371, 0, 0, 0, 0, 27875, 73765, 75794, 27875, 0, 73765, 59074, 26125, 68853, 38176, 46542, 0, 0, 0, 0, 0, 38176, 75794, 0, 51866, 0, 0, 0, 0, 96691, 94805, 0, 0, 0, 63638, 0, 96691, 0, 94805, 0, 0, 38176, 94805, 96691, 0, 0, 0, 0, 13194, 13194, 0]
    print(s.avoidFlood(a))

    print(s.avoidFlood([3, 5, 4, 0, 1, 0, 1, 5, 2, 8, 9]))
    print(s.avoidFlood([1, 2, 0, 0, 2, 1]))
    print(s.avoidFlood([1, 2, 0, 2, 0, 1]))
    print(s.avoidFlood([1, 2, 0, 2, 1]))
    print(s.avoidFlood([69, 0, 0, 0, 69]))
