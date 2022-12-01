from typing import List


class Solution:
    def diStringMatch(self, s: str) -> List[int]:
        # sort and wiggle
        res = [i for i in range(len(s)+1)]

        i = 0
        while i < len(s):
            j = i+1
            while j < len(s) and s[j] == s[i]:
                j += 1
            if s[i] == "I":
                res[i:j+1] = sorted(res[i:j+1])
            else:
                res[i:j+1] = sorted(res[i:j+1], reverse=True)
            i = j

        return res


if __name__ == '__main__':
    s = Solution()
    print(s.diStringMatch("IDID"))
