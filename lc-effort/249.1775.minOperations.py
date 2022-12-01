"""
https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/

kind of see this is a binary search problem
but given k, can it be done? this is the sub problem to solve

maybe using greedy..
A->B... B->A...
not a clear path yet...

so it is decided by the sum difference
diff_sum//5+1???
tried a few examples.. it checks out.. but not sure at all let me try

well.. if that is the way, then the answer is just that ^^? cannot be of course, let me encounter an error case
"""


from typing import Counter, List


class Solution_dumb_wrong:
    def minOperations(self, A: List[int], B: List[int]) -> int:
        if len(A) < len(B):
            A, B = B, A

        if len(A)*1 > len(B)*6:
            return -1

        sum_diff = abs(sum(A)-sum(B))
        return sum_diff//5+1 if sum_diff % 5 != 0 else sum_diff//5


"""
[5,2,1,5,2,2,2,2,4,3,3,5]
[1,4,5,5,6,3,1,3,3]
failed here... got 2 expected 1
so I am thinking maybe 5's multiples or not?
        sum_diff = abs(sum(A)-sum(B))
        return sum_diff//5+1 if sum_diff % 5 != 0 else sum_diff//5

then failed here
[3,3,2,4,2,6,2]
[6,2,6,6,1,1,4,6,4,6,2,5,4,2,1]

okay.. now to study

let me see.. to use 5.. you got have 1
to use 4 you got to have 2
...

>>> A = [3, 3, 2, 4, 2, 6, 2]
>>> B = [6, 2, 6, 6, 1, 1, 4, 6, 4, 6, 2, 5, 4, 2, 1]
>>> sum(A)
22
>>> sum(B)
56

so whatever to do.. is to make 34 out of the change
you can add 17 to A, minus 17 from B
you can add 34 to A, minus 0 from  B

you can add 35 to A, add 1 to B? if that is possible
e.g. sumA=7, all 1s; sumB=35, whatevers..

A + delta
B - (sum_diff-delta)


"""


class Solution_naive_binary_search:
    def minOperations(self, A: List[int], B: List[int]) -> int:
        m, n = len(A), len(B)
        if m > n*6 or n > m*6:
            return -1

        sumA, sumB = sum(A), sum(B)
        if sumA > sumB:
            A, B = B, A
            sumA, sumB = sum(A), sum(B)
            m, n = len(A), len(B)

        # A can add m*6 - sumA at most => delta
        # B can minus sum_diff - delata
        # see how many steps would be used
        # should I worry A to become smaller..
        # maybe not, because that means B has to cover sumDiff + delta
        # that is defintely more distance to cover and mightbe more steps..
        # also A to become smaller.. is also a step...
        counterA = Counter(A)
        counterB = Counter(B)

        def steps(counter, delta):
            if delta > 0:
                start, end, dir = 1, 6, 1
            else:
                start, end, dir = 6, 1, -1

            steps = 0
            while start != end and delta*dir > 0:
                if dir > 0:
                    gain = 6 - start
                    if delta <= gain*counter[start]:
                        steps += (delta-1)//gain+1
                        break
                    else:
                        steps += counter[start]
                        delta -= counter[start]*gain
                else:
                    loss = start-1
                    if abs(delta) <= loss * counter[start]:
                        steps += (abs(delta)-1)//loss + 1
                        break
                    else:
                        steps += counter[start]
                        delta += counter[start]*loss

                start += dir
            return steps

        l, r = 0, m*6-sumA
        res = 100000
        while l < r:
            delta = l+(r-l)//2
            stepsA = steps(counterA, delta)
            stepsB = steps(counterB, -(sumB-sumA-delta))
            if stepsA < stepsB:
                l = delta+1
            else:
                r = delta
            res = min(res, stepsA+stepsB)

        return res


"""
passed 53 tests
[6,5,6,3,1,5,4,1,1,5,3,5,5,4,5,6,5,6,4,5,2,1,5,5,6,2,1,5,4,6,4,6,1,1,4,4,2,2,3,3,1,3,1,2,6,6,4,6,1,4,2,4,1,2,1,4,6,6,1,1,3,2,3,3,1,6,2,4,2,1,1,4,5,2,4,2,6,2,6,1,6,3,5,1,3,5,1,2,2,3,1,1,2,6,3,4,3,1,6,5,1,3,3,2,1,2,5,1,3,2,3,2,5,4,1,2,3,3,4,5,2,4,3,3,3,5,6,5,1,1,1,1,1,3,1,5,4,2,2,5,2,1,2,3,6,1,2,3,3,3,5,5,3,5,5,2,1,5,6,2,5,5,5,1,4,4,6,1,4,6,3,4,5,5,4,2,5,4,4,4,1,5,3,6,3,1,1,4,1,6,1,4,3,1,5,6,6,4,5,6,2,5,4,4,6,6,5,4,6,1,2,3,6,2,2,5,6,1,2,4,2,6,2,1,5,1,3,3,2,1,3,2,5,5,1,2,3,1,3,4,5,5,5,3,2,5,2,1,2,5,5,6,3,5,2,5,1,2,3,6,4,3,3,5,4,1,2,4,5,4,4,1,2,2,3,4,3,4,2,2,1,6,1,2,2,5,3,3,5,5,5,2,1,5,5,3,2,5,5,4,4,5,6,2,3,3,3,6,6,6,1,6,2,2,4,6,6,5,5,3,2,3,4,1,1,3,2,3,1,2,6,1,4,2,1,6,4,5,1,1,5,3,2,5,3,5,2,5,2,4,2,1,4,3,4,5,2,4,6,5,1,2,2,1,4,3,1,4,4,2,5,4,5,3,4,2,2,6,4,6,1,3,1,5,5,1,6,2,4,2,3,4,3,6,4,3,1,6,6,2,2,2,6,5,1,6,6,3,2,6,3,5,4,1,5,2,5,4,5,5,1,1,3,3,6,3,2,6,4,3,3,1,3,5,3,4,4,3,2,2,6,3,4,3,3,1,5,6,3,2,6,4,5,1,3,4,3,1,1,2,1,3,2,4,3,6,3,1,5,5,1,6,6,4,2,3,6,1,5,5,4,5,5,3,3,4,6,6,6,4,2,5,4,1,4,6,4,1,5,3,3,5,4,6,5,6,5,1,4,2,2,1,4,5,2,5,1,4,5,3,2,6,6,5,3,4,4,1,3,5,5,2,2,4,3,5,2,4,3,5,2,6,1,6,3,4,3,4,2,4,6,2,6,6,4,1,2,6,1,1,5,4,6,5,6,5,2,3,4,4,5,6,2,1,4,5,2,4,3,1,2,4,1,4,4,2,4,4,4,2,4,4,2,5,4,2,5,4,6,2,6,6,2,6,3,4,4,5,3,3,1,5,6,4,3,3,3,2,4,5,1,1,2,4,6,4,1,4,5,5,1,1,1,1,1,6,2,5,5,1,3,5,2]
[3,6,5,6,1,5,3,3,4,5,6,2,4,4,2,2,1,5,6,5,2,5,2,4,3,4,5,4,1,1,2,4,4,2,4,4,3,1,4,1,3,5,4,1,1,6,2,5,2,6,2,2,2,4,6,4,1,3,2,2,1,3,1,4,2,5,1,1,2,3,1,4,6,1,4,6,2,1,3,1,2,2,6,6,1,5,4,5,6,4,6,2,3,4,6,5,3,5,2,5,1,4,4,1,1,6,4,4,6,2,5,5,4,3,5,5,4,4,6,5,5,4,4,6,1,2,5,2,3,4,4,1,6,6,1,5,2,3,1,6,5,2,2,5,4,4,5,6,1,3,4,5,5,2,3,5,4,6,3,2,2,1,5,3,2,1,2,6,5,3,1,2,2,2,2,6,2,5,1,4,3,5,6,4,4,6,3,4,4,6,4,3,4,3,3,4,1,1,1,6,3,4,6,6,5,4,6,1,4,5,2,1,5,1,2,4,1,2,2,4,4,3,4,1,4,1,3,5,6,4,2,2,1,4,1,1,5,6,5,6,5,1,2,4,4,3,2,6,4,3,2,2,2,2,2,5,2,6,2,5,2,5,6,1,4,4,2,6,1,6,4,4,3,6,3,6,4,1,1,1,2,3,1,4,1,4,4,5,6,6,5,6,5,2,1,2,5,6,1,2,1,1,4,1,1,2,6,3,2,4,5,3,2,6,4,4,6,6,5,2,1,3,4,2,2,1,3,1,5,5,2,3,6,6,3,1,1,6,5,2,3,1,6,3,1,3,6,2,6,1,1,1,1,3,4,5,2,3,3,2,4,5,1,3,2,2,2,6,6,5,1,4,1,5,2,3,4,3,5,2,4,4,1,2,2,1,2,5,4,3,5,6,6,3,2,4,6,5,2,3,5,3,4,3,2,6,4,3,3,4,6,2,5,1,6,1,6,4,5,6,5,5,4,5,6,2,3,2,6,2,4,5,5,3,3,2,5,6,2,5,5,1,4,2,1,6,1,5,3,4,6,5,3,6,3,2,6,2,2,3,4,6,4,4,2,6,4,2,4,5,3,2,1,5,6,4,2,6,3,2,6,1,1,3,4,2,3,2,1,2,4,1,3,2,4,2,5,5,2,3,4,6,1,4,4,5,6,3,5,6,2,2,2,1,2,3,6,5,6,2,6,5,4,5,1,3,1,1,6,3,1,1,4,2,6,5,4,5,5,3,1,3,6,4,5,2,4,2,1,4,1,5,3,5,6,5,3,6,5,3,2,4,4,1,3,6,6,2,4,3,4,6,4,4,4,6,3,2,1,6,5,1,2,5,5,4,6,3,4,4,1,6,5,4,2,5,2,1,2,3,5,1,6,4,6,6,1,6,6,5,2,3,6,1,1,6,1,3,5,2,1,1,3,2,5,4,1,5,3,1,2,3,2,3,2,4,4,6,3,2,5,3,3,5,5,1,6,3,2,1,2,2,3,5,3,6,3,3,4,5,3,5,2,2,4,5,4,4,3,4,1,2,2,3,5,4,5,1,4,2,4,5,3,1,3,2,2,6,5,3,5,3,3,4,4,3,4,6,1,6,3,3,2,2,1,5,1,6,4,4,6,1,4,6,1,2,3,4,3,5,6,5,1,1,3,4,5,3,6,6,5,4,3,1,2,2,4,2,4,5,1,4,3,6,4,6,5,2,4,2,4,3,5,1,1,1,1,3,3,6,6,1,4,4,2,2,3,5,5,6,5,3,6,5,3,1,5,6,1,4,3,4,2,4,6,4,5,5,5,1,2,4,4,5,1,4,6,6,5,3,2,2,5,2,2,1,4,1,2,4,2,4,5,4,5,6,3,3,6,5,6,5,5,2,3,2,1,5,2,1,6,4,2,2,6,5,6,3,1,2,4,3,1,4,6,3,3,4,3,6,5,1,6,5,5,5,3,1,3,2,5,6,2,5,1,4,6,6,2,1,5,1,5,3,3,5,5,3,4,1,4,4,4,3,5,6,1,6,4,2,5,3,6,6,2,5,1,1,5,4,5,5,2,2,5,1,5,6,1,3,4,4,6,3,2,5,6,4,1,6,4,1,6,5,5,6,6,5,6,1,5,6,4,2,1,4,4,5,6,6,3,1,4,2,3,5,2,5,4,4,3,2,3,5,6,5,1,2,3,5,1,1,2,5,6,2,4,2,5,2,6,4,5,1,4,2,2,4,3,5,4,3,1,2,1,4,2,2,4,4,3,5,6,1,2,4]

output 259
expect 255

this is way too hard to debug.. maybe time to give up now

maybe no binary search just serach all the deltas..
"""

"""
even this I got 259..
so maybe my function is not perfect..

that worthy debugging post I take a nap
"""


class Solution:
    def minOperations(self, A: List[int], B: List[int]) -> int:
        m, n = len(A), len(B)
        if m > n*6 or n > m*6:
            return -1

        sumA, sumB = sum(A), sum(B)
        if sumA > sumB:
            A, B = B, A
            sumA, sumB = sum(A), sum(B)
            m, n = len(A), len(B)

        # A can add m*6 - sumA at most => delta
        # B can minus sum_diff - delata
        # see how many steps would be used
        # should I worry A to become smaller..
        # maybe not, because that means B has to cover sumDiff + delta
        # that is defintely more distance to cover and mightbe more steps..
        # also A to become smaller.. is also a step...
        counterA = Counter(A)
        counterB = Counter(B)

        def steps(counter, delta):
            if delta > 0:
                start, end, dir = 1, 6, 1
            else:
                start, end, dir = 6, 1, -1

            steps = 0
            while start != end and delta*dir > 0:
                if dir > 0:
                    gain = 6 - start
                    if delta <= gain*counter[start]:
                        steps += (delta-1)//gain+1
                        delta = 0
                        break
                    else:
                        steps += counter[start]
                        delta -= counter[start]*gain
                else:
                    loss = start-1
                    if abs(delta) <= loss * counter[start]:
                        steps += (abs(delta)-1)//loss + 1
                        delta = 0
                        break
                    else:
                        steps += counter[start]
                        delta += counter[start]*loss

                start += dir
            return steps if delta*dir <= 0 else 10000000

        delta, maxDelta = 0, m*6-sumA
        sumDiff = sumB - sumA
        res = 10000000
        while delta <= min(maxDelta, sumDiff):
            stepsA = steps(counterA, delta)
            stepsB = steps(counterB, -(sumB-sumA-delta))
            res = min(res, stepsA+stepsB)
            delta += 1

        return res


"""
Runtime: 6452 ms, faster than 5.02% of Python3 online submissions for Equal Sum Arrays With Minimum Number of Operations.
Memory Usage: 18.4 MB, less than 60.40% of Python3 online submissions for Equal Sum Arrays With Minimum Number of Operations.
Next challenges:

okay.. at least.. this is ac solution
maybe I can then bring the binary search back
"""


class Solution:
    def minOperations(self, A: List[int], B: List[int]) -> int:
        m, n = len(A), len(B)
        if m > n*6 or n > m*6:
            return -1

        sumA, sumB = sum(A), sum(B)
        if sumA > sumB:
            A, B = B, A
            sumA, sumB = sum(A), sum(B)
            m, n = len(A), len(B)

        counterA = Counter(A)
        counterB = Counter(B)

        def steps(counter, delta):
            if delta > 0:
                start, end, dir = 1, 6, 1
            else:
                start, end, dir = 6, 1, -1

            steps = 0
            while start != end and delta*dir > 0:
                if dir > 0:
                    gain = 6 - start
                    if delta <= gain*counter[start]:
                        steps += (delta-1)//gain+1
                        delta = 0
                        break
                    else:
                        steps += counter[start]
                        delta -= counter[start]*gain
                else:
                    loss = start-1
                    if abs(delta) <= loss * counter[start]:
                        steps += (abs(delta)-1)//loss + 1
                        delta = 0
                        break
                    else:
                        steps += counter[start]
                        delta += counter[start]*loss

                start += dir
            return steps if delta*dir <= 0 else 10000000

        l, r = 0, min(m*6-sumA, sumB-sumA)

        res = 10000000
        while l < r:
            delta = l + (r-l)//2
            stepsA = steps(counterA, delta)
            stepsB = steps(counterB, -(sumB-sumA-delta))
            res = min(res, stepsA+stepsB)
            if stepsA < stepsB:
                l = delta+1
            elif stepsA > stepsB:
                r = delta-1
            else:
                ...  # ?

        return res


"""
off by 1??
yeah.. this is not right application of binary search in the root

so reading this
https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/discuss/1085859/Python-O(nlogn)-(faster-than-100.00)-with-explanation

sort the gain..
yeah.. I thought of putting all number together since modify every single number bears the same value/cost
then I lost in thinking, didn't come through of putting the gain together.. if I don't do that, I would need to 
treat some number as 1->6 and other as 6->1.. just mess

always one step away from the correctness. sign..
but with this idea, I can optimize it abit
use the counters.. I can put it back in..
"""


class Solution:
    def minOperations(self, A: List[int], B: List[int]) -> int:
        m, n = len(A), len(B)
        if m > n*6 or n > m*6:
            return -1

        sumA, sumB = sum(A), sum(B)
        if sumA > sumB:
            A, B = B, A

        counterA = Counter(A)
        counterB = Counter(B)

        gains = []  # (gain, available)
        for i in range(1, 6):
            j = 7-i
            available = counterA[i] + counterB[j]
            gains.append((6-i, available))

        sumDiff = abs(sumB-sumA)
        steps = 0
        for gain, available in gains:
            if sumDiff <= gain*available:
                steps += (sumDiff-1)//gain+1
                sumDiff = 0
                break
            else:
                steps += available
                sumDiff -= available*gain
        return steps


"""
Runtime: 1715 ms, faster than 72.81% of Python3 online submissions for Equal Sum Arrays With Minimum Number of Operations.
Memory Usage: 18.1 MB, less than 87.59% of Python3 online submissions for Equal Sum Arrays With Minimum Number of Operations.

woow... so the lesson is 
1. yes, after thinking a lot.. the numbers can be put together to be treated
2. but maybe not the raw numbers, the value.. (some function of numbers, f(num))
3. the counter usage increased the performance somehow..
"""

if __name__ == "__main__":
    s = Solution()

    A = [5, 2, 1, 5, 2, 2, 2, 2, 4, 3, 3, 5]
    B = [1, 4, 5, 5, 6, 3, 1, 3, 3]
    print(s.minOperations(A, B))

    A = [6, 6]
    B = [1]
    print(s.minOperations(A, B))

    A = [3, 3, 2, 4, 2, 6, 2]
    B = [6, 2, 6, 6, 1, 1, 4, 6, 4, 6, 2, 5, 4, 2, 1]
    print(s.minOperations(A, B))

    A = [6, 5, 6, 3, 1, 5, 4, 1, 1, 5, 3, 5, 5, 4, 5, 6, 5, 6, 4, 5, 2, 1, 5, 5, 6, 2, 1, 5, 4, 6, 4, 6, 1, 1, 4, 4, 2, 2, 3, 3, 1, 3, 1, 2, 6, 6, 4, 6, 1, 4, 2, 4, 1, 2, 1, 4, 6, 6, 1, 1, 3, 2, 3, 3, 1, 6, 2, 4, 2, 1, 1, 4, 5, 2, 4, 2, 6, 2, 6, 1, 6, 3, 5, 1, 3, 5, 1, 2, 2, 3, 1, 1, 2, 6, 3, 4, 3, 1, 6, 5, 1, 3, 3, 2, 1, 2, 5, 1, 3, 2, 3, 2, 5, 4, 1, 2, 3, 3, 4, 5, 2, 4, 3, 3, 3, 5, 6, 5, 1, 1, 1, 1, 1, 3, 1, 5, 4, 2, 2, 5, 2, 1, 2, 3, 6, 1, 2, 3, 3, 3, 5, 5, 3, 5, 5, 2, 1, 5, 6, 2, 5, 5, 5, 1, 4, 4, 6, 1, 4, 6, 3, 4, 5, 5, 4, 2, 5, 4, 4, 4, 1, 5, 3, 6, 3, 1, 1, 4, 1, 6, 1, 4, 3, 1, 5, 6, 6, 4, 5, 6, 2, 5, 4, 4, 6, 6, 5, 4, 6, 1, 2, 3, 6, 2, 2, 5, 6, 1, 2, 4, 2, 6, 2, 1, 5, 1, 3, 3, 2, 1, 3, 2, 5, 5, 1, 2, 3, 1, 3, 4, 5, 5, 5, 3, 2, 5, 2, 1, 2, 5, 5, 6, 3, 5, 2, 5, 1, 2, 3, 6, 4, 3, 3, 5, 4, 1, 2, 4, 5, 4, 4, 1, 2, 2, 3, 4, 3, 4, 2, 2, 1, 6, 1, 2, 2, 5, 3, 3, 5, 5, 5, 2, 1, 5, 5, 3, 2, 5, 5, 4, 4, 5, 6, 2, 3, 3, 3, 6, 6, 6, 1, 6, 2, 2, 4, 6, 6, 5, 5, 3, 2,
         3, 4, 1, 1, 3, 2, 3, 1, 2, 6, 1, 4, 2, 1, 6, 4, 5, 1, 1, 5, 3, 2, 5, 3, 5, 2, 5, 2, 4, 2, 1, 4, 3, 4, 5, 2, 4, 6, 5, 1, 2, 2, 1, 4, 3, 1, 4, 4, 2, 5, 4, 5, 3, 4, 2, 2, 6, 4, 6, 1, 3, 1, 5, 5, 1, 6, 2, 4, 2, 3, 4, 3, 6, 4, 3, 1, 6, 6, 2, 2, 2, 6, 5, 1, 6, 6, 3, 2, 6, 3, 5, 4, 1, 5, 2, 5, 4, 5, 5, 1, 1, 3, 3, 6, 3, 2, 6, 4, 3, 3, 1, 3, 5, 3, 4, 4, 3, 2, 2, 6, 3, 4, 3, 3, 1, 5, 6, 3, 2, 6, 4, 5, 1, 3, 4, 3, 1, 1, 2, 1, 3, 2, 4, 3, 6, 3, 1, 5, 5, 1, 6, 6, 4, 2, 3, 6, 1, 5, 5, 4, 5, 5, 3, 3, 4, 6, 6, 6, 4, 2, 5, 4, 1, 4, 6, 4, 1, 5, 3, 3, 5, 4, 6, 5, 6, 5, 1, 4, 2, 2, 1, 4, 5, 2, 5, 1, 4, 5, 3, 2, 6, 6, 5, 3, 4, 4, 1, 3, 5, 5, 2, 2, 4, 3, 5, 2, 4, 3, 5, 2, 6, 1, 6, 3, 4, 3, 4, 2, 4, 6, 2, 6, 6, 4, 1, 2, 6, 1, 1, 5, 4, 6, 5, 6, 5, 2, 3, 4, 4, 5, 6, 2, 1, 4, 5, 2, 4, 3, 1, 2, 4, 1, 4, 4, 2, 4, 4, 4, 2, 4, 4, 2, 5, 4, 2, 5, 4, 6, 2, 6, 6, 2, 6, 3, 4, 4, 5, 3, 3, 1, 5, 6, 4, 3, 3, 3, 2, 4, 5, 1, 1, 2, 4, 6, 4, 1, 4, 5, 5, 1, 1, 1, 1, 1, 6, 2, 5, 5, 1, 3, 5, 2]
    B = [3, 6, 5, 6, 1, 5, 3, 3, 4, 5, 6, 2, 4, 4, 2, 2, 1, 5, 6, 5, 2, 5, 2, 4, 3, 4, 5, 4, 1, 1, 2, 4, 4, 2, 4, 4, 3, 1, 4, 1, 3, 5, 4, 1, 1, 6, 2, 5, 2, 6, 2, 2, 2, 4, 6, 4, 1, 3, 2, 2, 1, 3, 1, 4, 2, 5, 1, 1, 2, 3, 1, 4, 6, 1, 4, 6, 2, 1, 3, 1, 2, 2, 6, 6, 1, 5, 4, 5, 6, 4, 6, 2, 3, 4, 6, 5, 3, 5, 2, 5, 1, 4, 4, 1, 1, 6, 4, 4, 6, 2, 5, 5, 4, 3, 5, 5, 4, 4, 6, 5, 5, 4, 4, 6, 1, 2, 5, 2, 3, 4, 4, 1, 6, 6, 1, 5, 2, 3, 1, 6, 5, 2, 2, 5, 4, 4, 5, 6, 1, 3, 4, 5, 5, 2, 3, 5, 4, 6, 3, 2, 2, 1, 5, 3, 2, 1, 2, 6, 5, 3, 1, 2, 2, 2, 2, 6, 2, 5, 1, 4, 3, 5, 6, 4, 4, 6, 3, 4, 4, 6, 4, 3, 4, 3, 3, 4, 1, 1, 1, 6, 3, 4, 6, 6, 5, 4, 6, 1, 4, 5, 2, 1, 5, 1, 2, 4, 1, 2, 2, 4, 4, 3, 4, 1, 4, 1, 3, 5, 6, 4, 2, 2, 1, 4, 1, 1, 5, 6, 5, 6, 5, 1, 2, 4, 4, 3, 2, 6, 4, 3, 2, 2, 2, 2, 2, 5, 2, 6, 2, 5, 2, 5, 6, 1, 4, 4, 2, 6, 1, 6, 4, 4, 3, 6, 3, 6, 4, 1, 1, 1, 2, 3, 1, 4, 1, 4, 4, 5, 6, 6, 5, 6, 5, 2, 1, 2, 5, 6, 1, 2, 1, 1, 4, 1, 1, 2, 6, 3, 2, 4, 5, 3, 2, 6, 4, 4, 6, 6, 5, 2, 1, 3, 4, 2, 2, 1, 3, 1, 5, 5, 2, 3, 6, 6, 3, 1, 1, 6, 5, 2, 3, 1, 6, 3, 1, 3, 6, 2, 6, 1, 1, 1, 1, 3, 4, 5, 2, 3, 3, 2, 4, 5, 1, 3, 2, 2, 2, 6, 6, 5, 1, 4, 1, 5, 2, 3, 4, 3, 5, 2, 4, 4, 1, 2, 2, 1, 2, 5, 4, 3, 5, 6, 6, 3, 2, 4, 6, 5, 2, 3, 5, 3, 4, 3, 2, 6, 4, 3, 3, 4, 6, 2, 5, 1, 6, 1, 6, 4, 5, 6, 5, 5, 4, 5, 6, 2, 3, 2, 6, 2, 4, 5, 5, 3, 3, 2, 5, 6, 2, 5, 5, 1, 4, 2, 1, 6, 1, 5, 3, 4, 6, 5, 3, 6, 3, 2, 6, 2, 2, 3, 4, 6, 4, 4, 2, 6, 4, 2, 4, 5, 3, 2, 1, 5, 6, 4, 2, 6, 3, 2, 6, 1, 1, 3, 4, 2, 3, 2, 1, 2, 4, 1, 3, 2, 4, 2, 5,
         5, 2, 3, 4, 6, 1, 4, 4, 5, 6, 3, 5, 6, 2, 2, 2, 1, 2, 3, 6, 5, 6, 2, 6, 5, 4, 5, 1, 3, 1, 1, 6, 3, 1, 1, 4, 2, 6, 5, 4, 5, 5, 3, 1, 3, 6, 4, 5, 2, 4, 2, 1, 4, 1, 5, 3, 5, 6, 5, 3, 6, 5, 3, 2, 4, 4, 1, 3, 6, 6, 2, 4, 3, 4, 6, 4, 4, 4, 6, 3, 2, 1, 6, 5, 1, 2, 5, 5, 4, 6, 3, 4, 4, 1, 6, 5, 4, 2, 5, 2, 1, 2, 3, 5, 1, 6, 4, 6, 6, 1, 6, 6, 5, 2, 3, 6, 1, 1, 6, 1, 3, 5, 2, 1, 1, 3, 2, 5, 4, 1, 5, 3, 1, 2, 3, 2, 3, 2, 4, 4, 6, 3, 2, 5, 3, 3, 5, 5, 1, 6, 3, 2, 1, 2, 2, 3, 5, 3, 6, 3, 3, 4, 5, 3, 5, 2, 2, 4, 5, 4, 4, 3, 4, 1, 2, 2, 3, 5, 4, 5, 1, 4, 2, 4, 5, 3, 1, 3, 2, 2, 6, 5, 3, 5, 3, 3, 4, 4, 3, 4, 6, 1, 6, 3, 3, 2, 2, 1, 5, 1, 6, 4, 4, 6, 1, 4, 6, 1, 2, 3, 4, 3, 5, 6, 5, 1, 1, 3, 4, 5, 3, 6, 6, 5, 4, 3, 1, 2, 2, 4, 2, 4, 5, 1, 4, 3, 6, 4, 6, 5, 2, 4, 2, 4, 3, 5, 1, 1, 1, 1, 3, 3, 6, 6, 1, 4, 4, 2, 2, 3, 5, 5, 6, 5, 3, 6, 5, 3, 1, 5, 6, 1, 4, 3, 4, 2, 4, 6, 4, 5, 5, 5, 1, 2, 4, 4, 5, 1, 4, 6, 6, 5, 3, 2, 2, 5, 2, 2, 1, 4, 1, 2, 4, 2, 4, 5, 4, 5, 6, 3, 3, 6, 5, 6, 5, 5, 2, 3, 2, 1, 5, 2, 1, 6, 4, 2, 2, 6, 5, 6, 3, 1, 2, 4, 3, 1, 4, 6, 3, 3, 4, 3, 6, 5, 1, 6, 5, 5, 5, 3, 1, 3, 2, 5, 6, 2, 5, 1, 4, 6, 6, 2, 1, 5, 1, 5, 3, 3, 5, 5, 3, 4, 1, 4, 4, 4, 3, 5, 6, 1, 6, 4, 2, 5, 3, 6, 6, 2, 5, 1, 1, 5, 4, 5, 5, 2, 2, 5, 1, 5, 6, 1, 3, 4, 4, 6, 3, 2, 5, 6, 4, 1, 6, 4, 1, 6, 5, 5, 6, 6, 5, 6, 1, 5, 6, 4, 2, 1, 4, 4, 5, 6, 6, 3, 1, 4, 2, 3, 5, 2, 5, 4, 4, 3, 2, 3, 5, 6, 5, 1, 2, 3, 5, 1, 1, 2, 5, 6, 2, 4, 2, 5, 2, 6, 4, 5, 1, 4, 2, 2, 4, 3, 5, 4, 3, 1, 2, 1, 4, 2, 2, 4, 4, 3, 5, 6, 1, 2, 4]

    print(s.minOperations(A, B))
