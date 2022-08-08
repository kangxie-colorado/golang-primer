from collections import defaultdict


class Solution:
    def longestConsecutive(self, nums) -> int:
        nums.sort()
        maxLen = 0
        i = 0
        while i < len(nums):
            l = 1
            while i+1 < len(nums):
                if nums[i] < nums[i+1] - 1:
                    break
                if nums[i]+1 == nums[i+1]:
                    l += 1
                i += 1

            i += 1
            maxLen = max(maxLen, l)

        return maxLen


class Solution:
    def longestConsecutive(self, nums) -> int:
        m = defaultdict(int)
        maxLen = 0
        for n in nums:
            if m[n] != 0:
                continue
            m[n] = 1
            if n-1 in m:
                m[n] += m[n-1]
            if n+1 in m:
                m[n] += m[n+1]
            m[n+m[n+1]] = m[n]
            m[n - m[n-1]] = m[n]
            maxLen = max(maxLen, m[n])

        return maxLen


if __name__ == '__main__':
    s = Solution()
    print(s.longestConsecutive([100, 4, 200, 1, 3, 2]))
    print(s.longestConsecutive([0, 3, 7, 2, 5, 8, 4, 6, 0, 1]))
