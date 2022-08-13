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


Solution().repeatedSubstringPattern('ababab')
