"""
https://leetcode.com/problems/single-number-ii/

haha.. iq test

the SUM is 3*M + n
so I am thinking I can use this 

if (SUM-n) is 3's multiples then n is candidation, if n%3!=0, then must be it
tricky is if n%3==0, how to tell.. 
then lets add 1 or minus 1 to make n-1%3 != 0

"""


from typing import List


class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        S = sum(nums)
        if S % 3 != 0:
            for n in nums:
                if n % 3 != 0 and (S-n) % 3 == 0 and (S-n*3) % 3 != 0:
                    return n
        else:
            S = sum([i//3 for i in nums]) + len(nums)
            for n in nums:
                n = n//3 + 1
                if n % 3 != 0 and (S-n) % 3 == 0 and (S-n*3) % 3 != 0:
                    return n


"""
the 2nd case is not right
when n%3==0, not handled right

this remains a problem to solve
n%3==0

"""


class Solution:
    # k = 3, p =2
    def singleNumber(self, nums: List[int]) -> int:
        def printBin(x, name):
            xBin = "".join(bin(x)[2:])
            print(f"{name:<6}", f"{x:<4}", f"{xBin:0>{m}}")

        def bindigits(n, bits):
            s = bin(n & int("1"*bits, 2))[2:]
            return ("{0:0>%s}" % (bits)).format(s)

        def printMaskBin(mask, bits=8):
            if mask > 0:
                maskBin = bin(mask)[2:]
            else:
                maskBin = bindigits(mask, bits)

            print(f"{'mask':<6}", f"{mask:<4}", f"{maskBin:0>{m}}")

        x1 = x2 = x3 = mask = 0
        m = 8
        for n in nums:
            print(f"Incoming:{n}, x2:{x2}, x1:{x1}")
            x3 ^= x2 & x1 & n
            x2 ^= x1 & n
            x1 ^= n
            mask = ~(x3 & ~x2 % ~x1)

            printBin(x3, "x3")
            printBin(x2, "x2")
            printBin(x1, "x1")
            printMaskBin(mask)

            x3 &= mask
            x2 &= mask
            x1 &= mask

            printBin(x3, "x3")
            printBin(x2, "x2")
            printBin(x1, "x1")
            print()

        return x1 | x2


class Solution:
    # k = 3, p =1
    def singleNumber(self, nums: List[int]) -> int:
        def printBin(x, name):
            xBin = "".join(bin(x)[2:])
            print(f"{name:<6}", f"{x:<4}", f"{xBin:0>{m}}")

        def bindigits(n, bits):
            s = bin(n & int("1"*bits, 2))[2:]
            return ("{0:0>%s}" % (bits)).format(s)

        def printMaskBin(mask, bits=8):
            if mask > 0:
                maskBin = bin(mask)[2:]
            else:
                maskBin = bindigits(mask, bits)

            print(f"{'mask':<6}", f"{mask:<4}", f"{maskBin:0>{m}}")

        x1 = x2 = mask = 0
        m = 8
        for n in nums:
            print(f"Incoming:{n}, x2:{x2}, x1:{x1}")
            x2 ^= x1 & n
            x1 ^= n
            mask = ~(x2 & x1)

            printBin(x2, "x2")
            printBin(x1, "x1")
            printMaskBin(mask)

            x2 &= mask
            x1 &= mask

            printBin(x2, "x2")
            printBin(x1, "x1")
            print()

        return x1 | x2


"""
https://leetcode.com/problems/single-number-ii/discuss/43297/Java-O(n)-easy-to-understand-solution-easily-extended-to-any-times-of-occurance

aha.. this one is equally smart...
"""

"""
wonder simple number 3 can be solved similarly?
only two occurences.. so only an x1... not place for two resuls...

hmm...
"""

if __name__ == "__main__":
    s = Solution()
    #assert 4 == s.singleNumber([2, 3, 4, 2, 3, 2, 3])
    #assert 4 == s.singleNumber([2, 3, 4, 2, 3, 2, 3, 4])
    assert 9 == s.singleNumber([3, 3, 3, 2, 2, 2, 9])
    assert 9 == s.singleNumber([2, 3, 2, 3, 3, 2, 9])

    #assert 99 == s.singleNumber([0, 1, 0, 1, 0, 1, 99])
