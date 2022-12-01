from functools import cmp_to_key
from typing import List


def twoSum(nums, target):
    i, j = 0, len(nums)-1
    res = []

    while i < j:
        S = nums[i]+nums[j]
        if S == target:
            res.append([nums[i], nums[j]])
            while i < j and nums[i] == nums[i+1]:
                i += 1
            while i < j and nums[j] == nums[j-1]:
                j -= 1
            i += 1
            j -= 1

        elif S < target:
            i += 1
        else:
            j -= 1
    return res


# https: // leetcode.com/problems/repeated-substring-pattern/
class Solution:
    def repeatedSubstringPattern(self, s: str) -> bool:
        bitmap = [0]*26
        for c in s:
            bitmap[ord(c) - ord('a')] ^= 1

        pat = ""
        for b in bitmap:
            pat = chr(c+ord('a'))

        return pat == s[:len(pat)]


# Solution().repeatedSubstringPattern('ababab')

def sqDigits(n):
    res = 0
    while n:
        res += (n % 10)**2
        n //= 10
    return res


def test_sqDigit():
    next = 2
    for i in range(20):
        next = sqDigits(next)
        print(next)


def compare(x1, x2):
    return -1 if x1 < x2 else 1


def compareWithDigitsAndChars(x1, x2):
    # so I think x1 is the later elemet
    # that actually makes sense, deciding a later element against a previous element
    # also here, if a lot of repeating code... instinctly refactor it into a function
    def isDigit(c):
        return '0' <= c <= '9'

    # if in the problem statement, the requirements are clear that either char or digit
    # one boolean function will be enough

    # x1 is behind x2... remember that
    if isDigit(x1) or isDigit(x2):
        if isDigit(x1) and isDigit(x2):
            # keep the order
            return 1
        elif isDigit(x1):
            # keep the order
            # a chance to refactor here
            return 1
        else:
            return -1
    else:
        return -1 if x1 < x2 else 1


class MyBase:
    def __init__(self, value) -> None:
        print("MyBase")
        self.value = value


class TimesTwo:
    def __init__(self) -> None:
        print("TimesTwo")
        self.value *= 2


class PlusFive:
    def __init__(self) -> None:
        print("PlusFive")
        self.value += 5


class MyCalss(MyBase,  PlusFive, TimesTwo):
    def __init__(self, value) -> None:
        print("MyClass")
        MyBase.__init__(self, value)
        TimesTwo.__init__(self)
        PlusFive.__init__(self)


if __name__ == '__main__':
    myobj = MyCalss(5)
    print(myobj.value)
