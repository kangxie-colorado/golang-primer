"""
https://leetcode.com/problems/valid-palindrome-ii/

this is categorized in greedy

"""


class Solution:
    def validPalindrome(self, s: str) -> bool:
        def isPalindrome(s, delIdx):
            i, j = 0, len(s)-1
            while i < j:
                if i == delIdx:
                    i += 1
                    continue
                if j == delIdx:
                    j -= 1
                    continue

                if s[i] != s[j]:
                    return False
                i, j = i+1, j-1
            return True

        for delIdx in range(-1, len(s)):
            if isPalindrome(s, delIdx):
                return True

        return False


"""
this TLE...
let me think the greedy one

so just start with this string..
if it is a palindrom then fine

otherwise, on the first mismatch.. just do two branches..
"""


class Solution:
    def validPalindrome(self, s: str) -> bool:
        def isPalindrome(s):
            i, j = 0, len(s)-1
            while i < j:
                if s[i] != s[j]:
                    return False
                i, j = i+1, j-1
            return True
        i, j = 0, len(s)-1
        while i < j:
            if s[i] != s[j]:
                return isPalindrome(s[i:j]) or isPalindrome(s[i+1:j+1])
            i, j = i+1, j-1
        return True


"""
Runtime: 327 ms, faster than 17.76% of Python3 online submissions for Valid Palindrome II.
Memory Usage: 14.5 MB, less than 88.73% of Python3 online submissions for Valid Palindrome II.

Runtime: 215 ms, faster than 51.46% of Python3 online submissions for Valid Palindrome II.
Memory Usage: 14.5 MB, less than 45.64% of Python3 online submissions for Valid Palindrome II.
"""
