"""
https://leetcode.com/problems/count-primes/


"""


from collections import defaultdict


class Solution:
    def countPrimes(self, n: int) -> int:
        A = [1 for i in range(max(2, n+1))]
        A[0] = A[1] = 0

        i = 2
        count = 0
        while i < n:
            if A[i] == 0:
                i += 1
                continue

            count += 1
            p = i*2
            while i*t <= n:
                A[t*i] = 0
                t += 1
            i += 1

        return count


"""
TLE as expected
maybe use hashmap more memory to give more time
huh.. stupid.. no diff

so just divide all the previous prime.. if not dividable by any, its a prime..
"""


class Solution:
    def countPrimes(self, n: int) -> int:
        if n <= 2:
            return 0

        i = 3
        primes = [2]
        count = 1
        while i < n:
            for p in primes:
                if i % p == 0:
                    break
            else:
                count += 1
                primes.append(i)

            i += 1

        return count


"""
Still TLE

In fact, we can mark off multiples of 5 starting at 5 × 5 = 25, 
because 5 × 2 = 10 was already marked off by multiple of 2, 
similarly 5 × 3 = 15 was already marked off by multiple of 3. 
Therefore, if the current number is p, we can always mark off 
multiples of p starting at p^2, then in increments of 
p: p^2 + p, p^2 + 2p, ... Now what should be the terminating loop condition?

"""


class Solution:
    def countPrimes(self, n: int) -> int:
        A = [1 for i in range(max(2, n+1))]
        A[0] = A[1] = 0

        i = 2
        count = 0
        while i < n:
            if A[i] == 0:
                i += 1
                continue

            count += 1
            p = i*i
            while p <= n:
                A[p] = 0
                p += i
            i += 1

        return count


"""
okay.. still TLE
867896

more hints
It is easy to say that the terminating loop condition is p < n,
 which is certainly correct but not efficient. Do you still remember Hint #3?

 ^^ okay.. so when a prime's square is over n, then the loop can be done.. 
 because all non-primes should have been marked by smaller primes.. daman good
"""


class Solution:
    def countPrimes(self, n: int) -> int:
        A = [1 for i in range(max(2, n+1))]
        A[0] = A[1] = 0

        i = 2
        while i*i < n:
            if A[i] == 0:
                i += 1
                continue

            p = i*i
            while p < n:
                A[p] = 0
                p += i
            i += 1

        count = 0
        for i in range(n):
            count += A[i]

        return count


"""
Runtime: 9019 ms, faster than 6.24% of Python3 online submissions for Count Primes.
Memory Usage: 67.7 MB, less than 50.61% of Python3 online submissions for Count Primes.
"""

if __name__ == '__main__':
    s = Solution()
    print(s.countPrimes(2))
    print(s.countPrimes(3))
    print(s.countPrimes(6))
    print(s.countPrimes(469193))
