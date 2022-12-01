"""
https://leetcode.com/problems/happy-number/

I did this before but I have a misunderstanding 
I was unsure how the terminal condition could present itself

loop 100 times? 1000 times?
but turn out, just loop until it becomes 1 or re-appears...
"""


class Solution:
    def isHappy(self, n: int) -> bool:
        def getNext(n):
            res = 0
            while n:
                res += (n % 10)**2
                n //= 10
            return res

        s = set()
        nextN = getNext(n)
        while nextN != 1 and nextN not in s:
            s.add(nextN)
            nextN = getNext(nextN)

        return nextN == 1


"""
haha.. I was wondering this is tagged two pointer
because there is a cycle... cycle detection can be done via fast/slow pointers...
something new indeed
"""


class Solution:
    def isHappy(self, n: int) -> bool:
        def getNext(n):
            res = 0
            while n:
                res += (n % 10)**2
                n //= 10
            return res

        slow = getNext(n)
        fast = getNext(getNext(n))

        while slow != fast:
            slow = getNext(slow)
            fast = getNext(getNext(fast))

        return slow == 1


"""
cool...
"""
