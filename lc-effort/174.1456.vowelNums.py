"""
https://leetcode.com/problems/maximum-number-of-vowels-in-a-substring-of-given-length/


"""


from curses.ascii import SO


class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        count, res = 0, 0
        i, j = 0, 0

        while j < k:
            count += 'aeiou'.count(s[j])
            j += 1
        res = count
        while j < len(s):
            i = j-k
            count += 'aeiou'.count(s[j])
            count -= 'aeiou'.count(s[i])
            res = max(res, count)
            j += 1

        return res

        """
Runtime: 283 ms, faster than 50.46% of Python3 online submissions for Maximum Number of Vowels in a Substring of Given Length.
Memory Usage: 15 MB, less than 21.52% of Python3 online submissions for Maximum Number of Vowels in a Substring of Given Length.
Next challenges:
        """


if __name__ == '__main__':
    s = Solution()
    print(s.maxVowels('abciiidef', 3))
