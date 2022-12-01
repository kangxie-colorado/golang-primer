"""
https://leetcode.com/problems/decode-ways-ii/

a hard but fate has it lets have it
if stuck just try to learn the solution

no need to bang my head
"""

"""
So I actually did the decode ways I, which is without the * to mess things up
but even I cannot remember how to do that

since I am here.. 
why I don't go to redo that for a refresher...  
"""


# decode ways I
"""
just thinking which way to start and how to bind two digits together
still feel like I need to start from the end
do I?

... 1 2 7
s[i] = 0 initialize to 0 then do following
 - s[i] == 0, can only bind to prev digit, s[i] += s[i-2] if s[i-1:i+1] <= 20 (10 or 20)
 - s[i-1:i+1] <= 26, e.g. 1,2, s[i] += s[i-2], binding with s[i-1]; or not, +=s[i-1]
 - s[i-1:i+1] > 26, e.g 2,7, cannot bind with s[i-1], so s[i] = s[i-1]



"""




from ast import Num
class Solution:
    def numDecodings(self, s: str) -> int:
        dp = [0]*(len(s)+1)
        dp[0] = 1
        dp[1] = 1 if s[0] != "0" else 0

        # dp[i] <==> s[i-1]
        for i in range(1, len(s)):
            dp[i+1] = 0
            if s[i-1] == "0":
                dp[i+1] += dp[i] if s[i] != "0" else 0
            else:
                if s[i] == "0":
                    if s[i-1] in "12":
                        dp[i+1] += dp[i-1]

                elif s[i-1:i+1] <= "26":
                    dp[i+1] += dp[i] + dp[i-1]
                elif s[i-1:i+1] > "26":
                    dp[i+1] += dp[i]

            if dp[i+1] == 0:
                return 0

        return dp[len(s)]


"""
Runtime: 50 ms, faster than 54.41% of Python3 online submissions for Decode Ways.
Memory Usage: 13.8 MB, less than 80.45% of Python3 online submissions for Decode Ways.

alright.. finally
I know that 30 is a problem but I didn't change the code
the other edge cases are two 00s..

            dp[i+1] += dp[i] if s[i] != "0" and dp[i] else 0
                        ^ need to add this number, not 1
                                                ^ this dp[i] is redundant now
                                        
Runtime: 35 ms, faster than 90.51% of Python3 online submissions for Decode Ways.
Memory Usage: 13.9 MB, less than 80.45% of Python3 online submissions for Decode Ways.

the i+1 is quite messy
let me see if I can make it a bit better
"""


class Solution:
    def numDecodings(self, s: str) -> int:
        dp = [0]*(len(s)+1)
        dp[0] = 1
        dp[1] = 1 if s[0] != "0" else 0

        # dp[i] <==> s[i-1]
        for i in range(2, len(dp)):
            dp[i] = 0
            if s[i-2] == "0":
                dp[i] += dp[i-1] if s[i-1] != "0" else 0
            else:
                if s[i-1] == "0":
                    if s[i-2] in "12":
                        dp[i] += dp[i-2]

                elif s[i-2:i] <= "26":
                    dp[i] += dp[i-1] + dp[i-2]
                elif s[i-2:i] > "26":
                    dp[i] += dp[i-1]

            if dp[i] == 0:
                return 0

        return dp[len(s)]


"""
pass
but equally messy..
"""

# with the "*"


class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == "0":
            return 0

        dp = [0]*(len(s)+1)
        dp[0] = 1
        dp[1] = 1
        if s[0] == "*":
            dp[1] = 9

        # dp[i] <==> s[i-1]
        for i in range(2, len(dp)):
            j = i-1
            if s[j] == "*":
                if s[j-1] == "0":
                    dp[i] = dp[i-1]*9
                elif s[j-1] == "1":
                    # alone and combine s[j-1]
                    dp[i] = dp[i-1]*9 + dp[i-2]*9
                elif s[j-1] == "2":
                    dp[i] = dp[i-1]*9 + dp[i-2]*6
                elif s[j-1] >= "3":
                    dp[i] = dp[i-1]*9
                elif s[j-1] == "*":
                    dp[i] = dp[i-1]*9 + dp[i-2]*15
            else:
                if s[j-1] == "*":
                    if s[j] == "0":
                        dp[i] = dp[i-2]*2
                    elif "0" < s[j] <= "6":
                        dp[i] = dp[i-1] + dp[i-2]*2
                    elif s[j] > "6":
                        # alone or combine
                        # only combine for 1
                        dp[i] = dp[i-1] + dp[i-2]
                else:
                    # regular cases
                    if s[j-1] == "0":
                        dp[i] += dp[i-1] if s[j] != "0" else 0
                    else:
                        if s[j] == "0":
                            if s[j-1] in "12":
                                dp[i] += dp[i-2]

                        elif s[j-1:j+1] <= "26":
                            dp[i] += dp[i-1] + dp[i-2]
                        elif s[j-1:j+1] > "26":
                            dp[i] += dp[i-1]

            if dp[i] == 0:
                return 0

        return dp[len(s)] % (10**9+7)


"""
super messy
but first try not too bad 195 / 217 test cases passed.

"1*72*"
expected 285 got 270
okay here
                    elif s[j] > "6":
                        # alone or combine
                        # only combine for 1
                        dp[i] = dp[i-1] + dp[i-2]

208 / 217 test cases passed.
Input: "1*1*22*19"
Output: 17012
Expected: 19064

let me see up to which substring I am still ok

throw in 
"1*1*"
"1*1*22"
"1*1*22*" <== okay.. wrong here already, so s[j-1] == 2 case
"1*1*22*19"

so 
input           expected        output
"1*1*2"         382             382
"1*1*22"        724             724
"1*1*22*"       8808            7782

let me debug

okay...

memory limit
time limit..

but maybe the cases are right now
I also see here is a generic way to do this

basically it is function of dp[i-2] and dp[i-1], the different is the coefficient
"""

N = 10**9+7


class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == "0":
            return 0

        dp = [0]*(len(s)+1)
        dp[0] = 1
        dp[1] = 1
        if s[0] == "*":
            dp[1] = 9

        # dp[i] <==> s[i-1]
        for i in range(2, len(dp)):
            j = i-1
            if s[j] == "*":
                if s[j-1] == "0":
                    dp[i] = dp[i-1]*9
                elif s[j-1] == "1":
                    # alone and combine s[j-1]
                    dp[i] = (dp[i-1]*9 + dp[i-2]*9) % N
                elif s[j-1] == "2":
                    dp[i] = (dp[i-1]*9 + dp[i-2]*6) % N
                elif s[j-1] >= "3":
                    dp[i] = dp[i-1]*9 % N
                elif s[j-1] == "*":
                    dp[i] = (dp[i-1]*9 + dp[i-2]*15) % N
            else:
                if s[j-1] == "*":
                    if s[j] == "0":
                        dp[i] = dp[i-2]*2 % N
                    elif "0" < s[j] <= "6":
                        dp[i] = (dp[i-1] + dp[i-2]*2) % N
                    elif s[j] > "6":
                        # alone or combine
                        # only combine for 1
                        dp[i] = (dp[i-1] + dp[i-2]) % N
                else:
                    # regular cases
                    if s[j-1] == "0":
                        dp[i] += dp[i-1] if s[j] != "0" else 0
                    else:
                        if s[j] == "0":
                            if s[j-1] in "12":
                                dp[i] += dp[i-2]

                        elif s[j-1:j+1] <= "26":
                            dp[i] += dp[i-1] + dp[i-2]
                        elif s[j-1:j+1] > "26":
                            dp[i] += dp[i-1]

            if dp[i] == 0:
                return 0

        return dp[len(s)] % (10**9+7)


class Solution:
    def numDecodings(self, s: str) -> int:
        if s[0] == "0":
            return 0

        dp2 = 1
        dp1 = 1
        if s[0] == "*":
            dp1 = 9

        def getCoefficients(comb):
            if "*" in comb:
                if comb[1] == "*":
                    if comb[0] == "0" or comb[0] >= "3":
                        c1, c2 = 9, 0
                    elif comb[0] == "1":
                        c1, c2 = 9, 9
                    elif comb[0] == "2":
                        c1, c2 = 9, 6
                    elif comb[0] == "*":
                        c1, c2 = 9, 15
                else:
                    if comb[1] == "0":
                        c1, c2 = 0, 2
                    elif "1" <= comb[1] <= "6":
                        c1, c2 = 1, 2
                    elif comb[1] > "6":
                        c1, c2 = 1, 1
            else:
                if comb == "00" or comb[0] > "2" and comb[1] == "0":
                    c1, c2 = 0, 0
                elif comb in ("10", "20"):
                    c1, c2 = 0, 1
                elif "01" <= comb <= "09" or comb > "26":
                    c1, c2 = 1, 0
                elif comb <= "26":
                    c1, c2 = 1, 1

            return c1, c2

        # next = dp2*c2 + dp1*c2
        # then
        #   dp2,dp1=dp1,next
        for i in range(2, len(s)+1):
            j = i-1
            comb = s[j-1:j+1]
            c1, c2 = getCoefficients(comb)
            next = (dp2*c2 + dp1*c1) % (10**9+7)
            dp2, dp1 = dp1, next

        return dp1


"""
Runtime: 365 ms, faster than 96.19% of Python3 online submissions for Decode Ways II.
Memory Usage: 18.6 MB, less than 59.17% of Python3 online submissions for Decode Ways II.

wowo......
"""

if __name__ == "__main__":
    s = Solution()
    print(s.numDecodings("10011"))
    print(s.numDecodings("12"))
    print(s.numDecodings("226"))
    print(s.numDecodings("06"))
    print(s.numDecodings("2611055971756562"))
    print(s.numDecodings("230"))

    print(s.numDecodings("1*"))
    print(s.numDecodings("1**"))
    print(s.numDecodings("1***"))
    print(s.numDecodings("2*"))
    print(s.numDecodings("1**6"))
    print(s.numDecodings("1*72*"))
    print(s.numDecodings("1*1*22*"))
