"""
https://leetcode.com/problems/fruit-into-baskets/

so this is basically longest at-most-2-uniq number subarray
"""


from collections import defaultdict
from typing import List


class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        res, uniqs = 0, 0
        i = 0
        m = defaultdict(int)

        for j, f in enumerate(fruits):
            m[f] += 1
            if m[f] == 1:
                uniqs += 1
            while uniqs > 2:
                m[fruits[i]] -= 1
                if m[fruits[i]] == 0:
                    uniqs -= 1
                i += 1

            res = max(res, j-i+1)

        return res


"""
Runtime: 1564 ms, faster than 32.92% of Python3 online submissions for Fruit Into Baskets.
Memory Usage: 20.8 MB, less than 5.30% of Python3 online submissions for Fruit Into Baskets.

Runtime: 996 ms, faster than 88.05% of Python3 online submissions for Fruit Into Baskets.
Memory Usage: 20 MB, less than 88.10% of Python3 online submissions for Fruit Into Baskets.

"""


class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        uniqs = 0
        i = 0
        m = defaultdict(int)

        for j, f in enumerate(fruits):
            m[f] += 1
            if m[f] == 1:
                uniqs += 1
            if uniqs > 2:
                m[fruits[i]] -= 1
                if m[fruits[i]] == 0:
                    uniqs -= 1
                i += 1

        return j-i+1


"""
non-shrinking..
Runtime: 925 ms, faster than 93.62% of Python3 online submissions for Fruit Into Baskets.
Memory Usage: 20 MB, less than 95.76% of Python3 online submissions for Fruit Into Baskets.

"""

"""
okay.. this Lee tells another solution
for a smaller amount of variables, it seems we can always enumerate all the possibilities

Explanation:
Loop all fruit c in tree,
Note that a and b are the last two different types of fruit that we met,
c is the current fruit type,
so it's something like "....aaabbbc..."

Case 1 c == b:
fruit c already in the basket,
and it's same as the last type of fruit
cur += 1
count_b += 1

Case 2 c == a:
fruit c already in the basket,
but it's not same as the last type of fruit
cur += 1
count_b = 1
a = b, b = c

Case 3 c != b && c!= a:
fruit c not in the basket,
cur = count_b + 1
count_b = 1
a = b, b = c

Of course, in each turn we need to update res = max(res, cur)

I didn't try to understand his code and from the description I cannot translate directly to code
but the idea is pretty good.. 
let me try turning idea into code...

"""


class Solution:
    def totalFruit(self, fruits: List[int]) -> int:
        a, b = fruits[0], -1
        run = countB = 0

        i = 0
        while i < len(fruits) and fruits[i] == a:
            run += 1
            i += 1
        if i < len(fruits):
            b = fruits[i]
            run += 1
            countB += 1

        """
        looks like important in the loop is to maintain the countB
        because it will start the next baskets run when c switches into 3rd value
        """

        res = run
        for j in range(i+1, len(fruits)):
            c = fruits[j]
            if c in (a, b):
                run += 1
                if c == a:
                    countB = 1
                    a, b = b, a
                else:
                    countB += 1
            else:
                run = countB + 1
                countB = 1
                a, b = b, c

            res = max(res, run)
        return res


"""
Runtime: 1648 ms, faster than 26.01% of Python3 online submissions for Fruit Into Baskets.
Memory Usage: 19.9 MB, less than 95.76% of Python3 online submissions for Fruit Into Baskets.

okay... now I understand his code
    def totalFruit(self, tree):
        res = cur = count_b = a = b = 0
        for c in tree:
            cur = cur + 1 if c in (a, b) else count_b + 1
            count_b = count_b + 1 if c == b else 1
            if b != c: a, b = b, c
            res = max(res, cur)
        return res

roll my eyes... wo cao..
but basically the same thing as me..
"""

if __name__ == '__main__':
    s = Solution()
    print(s.totalFruit([0, 1, 2, 2]))
    print(s.totalFruit([1, 1, 1]))
