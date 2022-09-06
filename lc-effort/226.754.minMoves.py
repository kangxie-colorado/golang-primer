"""
https://leetcode.com/problems/reach-a-number/

n=1, possibilities = (-1.1)
n=2, possibilities = (-3,-1,1,3)
n=3, possibilities = (-6, -4, -2, 0, 2, 4, 6)
it is symmetric on two sides of 0 at least

n=4, possibilites = (1,3,6,10), then (-10,-6,-3,-1)
    from (-6,-3,-1) taking 4th move, (-2,1,3)
    from (-3,-1) taking 3rd move to right (0, 2) then 4th (4, 6) (-4, -2)..

    so this is exploding.. and hard to figure out..
    (wrong numbers above for n=4 but the idea is same)
    or basically from the results of n=3
    (-6, -4, -2, 0, 2, 4, 6), you can take 4th from any of these 
    (-10, -2, -8, 0, -6, 2, -4, 4, -2, 6, 0, 8, 2, 10)
    >>> sorted(set(a))
    [-10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10]

n=5, [-15, -13, -11, -9, -7, -5, -3, -1, 1, 3, 5, 7, 9, 11, 13, 15]

so f(4) absolutly contains f(3).. so yeah this is monotonic
maybe can adopt the binary search
with n=5, this is not that plain simple

n=6, [-21, -19, -17, -15, -13, -11, -9, -7, -5, -3, -1, 1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21]
so the range is [-n*(n+1)/2:n*(n+1)/2+1:2]
n=8
>>> list(range(-n*(n+1)//2, n*(n+1)//2+1, 2))
[-36, -34, -32, -30, -28, -26, -24, -22, -20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36]

n=7
>>> list(range(-n*(n+1)//2, n*(n+1)//2+1, 2))
[-28, -26, -24, -22, -20, -18, -16, -14, -12, -10, -8, -6, -4, -2, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28]

so to end up with even number 
n or n+1 needs to be multiple times of 4
so n*(n+1)//2 can be even
"""


from math import remainder


class Solution:
    def reachNumber(self, target: int) -> int:
        target = abs(target)
        if target % 2 == 0:
            # even number, n or n+1 must be 4*x
            # just to see target falls between which n and n+1
            # 3 or 4; 7 or 8; 11 or 12...
            # actually decide >4, >7 >12 first, then decide the finer scope
            l, r = 1, 10**9//4
            while l < r:
                m = l+(r-l)//2
                if 4*m*(4*m+1)//2 < target:
                    l = m+1
                else:
                    r = m

            if 4*l*(4*l-1)//2 >= target:
                return 4*l-1
            return 4*l

        else:
            l, r = 1, 10**9
            while l < r:
                m = l+(r-l)//2
                if m*(m+1)//2 < target:
                    l = m+1
                else:
                    r = m

            if l % 4 == 0:
                return l+1
            if (l+1) % 4 == 0:
                return l+2
            return l


"""
Runtime: 64 ms, faster than 76.50% of Python3 online submissions for Reach a Number.
Memory Usage: 14 MB, less than 13.36% of Python3 online submissions for Reach a Number.

let me try simplify the code 
"""


class Solution:
    def reachNumber(self, target: int) -> int:
        target = abs(target)

        l, r = 1, 10**9
        while l < r:
            m = l+(r-l)//2
            if m*(m+1)//2 < target:
                l = m+1
            else:
                r = m

        if target % 2 == 0:
            if l % 4 != 0 and (l+1) % 4 != 0:
                if l % 4 == 1:
                    # e.g. l=5, need to return 7
                    if (l+2)*(l+3) >= target:
                        return l+2
                    return l+3
                elif l % 4 == 2:
                    # e.g l=2, need to return 3
                    if (l+1)*(l+2) >= target:
                        return l+1
                    return l+2
        else:
            if l % 4 == 0:
                return l+1
            if (l+1) % 4 == 0:
                return l+2

        return l


"""
Runtime: 49 ms, faster than 81.57% of Python3 online submissions for Reach a Number.
Memory Usage: 14 MB, less than 13.36% of Python3 online submissions for Reach a Number.

Runtime: 38 ms, faster than 92.17% of Python3 online submissions for Reach a Number.
Memory Usage: 13.8 MB, less than 62.21% of Python3 online submissions for Reach a Number.
"""


class Solution:
    def reachNumber(self, target: int) -> int:
        target = abs(target)

        l, r = 1, 10**9
        while l < r:
            m = l+(r-l)//2
            if m*(m+1)//2 < target:
                l = m+1
            else:
                r = m

        if target % 2 == 0:
            if l % 4 != 0 and (l+1) % 4 != 0:
                if l % 4 == 1:
                    # e.g. l=5, need to return 7
                    if (l+2)*(l+3) >= target:
                        return l+2
                    return l+3
                elif l % 4 == 2:
                    # e.g l=2, need to return 3
                    if (l+1)*(l+2) >= target:
                        return l+1
                    return l+2
            # return l # can be replaced by one return at the end
        else:
            if l % 4 == 0:
                return l+1
            if (l+1) % 4 == 0:
                return l+2
            # return l # can be replaced by one return at the end
        return l


if __name__ == "__main__":
    s = Solution()
    print(s.reachNumber(2))
    print(s.reachNumber(4))
    print(s.reachNumber(6))
    print(s.reachNumber(8))

    print(s.reachNumber(5))
    print(s.reachNumber(15))
    print(s.reachNumber(13))
    print(s.reachNumber(23))
