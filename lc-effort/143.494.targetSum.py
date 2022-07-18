"""
https://leetcode.com/problems/target-sum/

backtracking but 2^20 will it pass????
"""


from cProfile import run
from collections import defaultdict
from os import defpath


class Solution(object):
    def findTargetSumWays(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        count = 0

        def backtracking(i, sumSofar):
            nonlocal count
            if i == len(nums):
                if sumSofar == target:
                    count += 1
                return

            backtracking(i+1, sumSofar+nums[i])
            backtracking(i+1, sumSofar-nums[i])

        backtracking(0, 0)
        return count


"""
TLE here..
[33,36,38,40,25,49,1,8,50,13,41,50,29,27,18,25,37,8,0,48]
0

of course

hmm.. I discovered the mirror effect but I still don't feel it is going to pass
just try
"""


class Solution(object):
    def findTargetSumWays(self, nums, target):
        prefix = 0
        for i in range(len(nums)):
            nums[i] += prefix
            prefix = nums[i]

        count = 0
        seed = nums[-1]
        if abs(seed) == abs(target):
            count += 1
        results = [seed]
        for i in range(len(nums)-2, -1, -1):
            newResults = []
            for r in results:
                result = nums[i]*2 - r
                if abs(result) == abs(target):
                    count += 1
                newResults.append(result)

            results.extend(newResults)

        # 0 should always double.. the opposite will have the same 0, not negative
        return count if target else count*2


"""
but still TLE -- passed 109.. better than first version

[5,40,23,47,43,19,36,10,28,46,14,11,5,0,5,22,39,30,50,41]
48

hah. seems like I can reduce by another half by adding two targets...
"""


class Solution(object):
    def findTargetSumWays(self, nums, target):
        if len(nums) <= 4:
            prefix = 0
            for i in range(len(nums)):
                nums[i] += prefix
                prefix = nums[i]

            count = 0
            seed = nums[-1]
            if abs(seed) == abs(target):
                count += 1
            results = [seed]
            for i in range(len(nums)-2, -1, -1):
                newResults = []
                for r in results:
                    result = nums[i]*2 - r
                    if abs(result) == abs(target):
                        count += 1
                    newResults.append(result)

                results.extend(newResults)

            # 0 should always double.. the opposite will have the same 0, not negative
            return count if target else count*2
        else:
            prefix = 0
            for i in range(len(nums)):
                nums[i] += prefix
                prefix = nums[i]

            count = 0
            seed = nums[-1]
            # targets = set(target, -target, nums[0]*2-target, nums[0]*2+target)
            targets = defaultdict(int)
            targets[target] += 1
            targets[-target] += 1
            targets[nums[0]*2-target] += 1
            targets[nums[0]*2+target] += 1

            targets[nums[1]*2-target] += 1
            targets[nums[1]*2+target] += 1
            targets[nums[1]*2-nums[0]*2-target] += 1
            targets[nums[1]*2-nums[0]*2+target] += 1

            count += targets[seed]
            results = [seed]
            for i in range(len(nums)-2, 1, -1):
                newResults = []
                for r in results:
                    result = nums[i]*2 - r
                    count += targets[result]
                    newResults.append(result)

                results.extend(newResults)

            # now the count is factered in..
            # becasue +0, -0 will naturally count twice
            return count


"""
Runtime: 6774 ms, faster than 5.01% of Python3 online submissions for Target Sum.
Memory Usage: 24.2 MB, less than 47.26% of Python3 online submissions for Target Sum.

wow..hohoho..

I see the path.. it converts into a adding targets and reducing searching space game
but if you go all the way.. it is still 2^n..

maybe stike a balance.. just three levels?

"""


class Solution_fold3(object):
    def findTargetSumWays(self, nums, target):
        if len(nums) <= 4:
            prefix = 0
            for i in range(len(nums)):
                nums[i] += prefix
                prefix = nums[i]

            count = 0
            seed = nums[-1]
            if abs(seed) == abs(target):
                count += 1
            results = [seed]
            for i in range(len(nums)-2, -1, -1):
                newResults = []
                for r in results:
                    result = nums[i]*2 - r
                    if abs(result) == abs(target):
                        count += 1
                    newResults.append(result)

                results.extend(newResults)

            # 0 should always double.. the opposite will have the same 0, not negative
            return count if target else count*2
        else:
            prefix = 0
            for i in range(len(nums)):
                nums[i] += prefix
                prefix = nums[i]

            count = 0
            seed = nums[-1]
            # targets = set(target, -target, nums[0]*2-target, nums[0]*2+target)
            targets = defaultdict(int)
            targets[target] += 1

            targets[-target] += 1

            targets[nums[0]*2-target] += 1
            targets[nums[0]*2+target] += 1

            targets[nums[1]*2-target] += 1
            targets[nums[1]*2+target] += 1
            targets[nums[1]*2-nums[0]*2-target] += 1
            targets[nums[1]*2-nums[0]*2+target] += 1

            targets[nums[2]*2+target] += 1
            targets[nums[2]*2-target] += 1
            targets[nums[2]*2-nums[0]*2-target] += 1
            targets[nums[2]*2-nums[0]*2+target] += 1
            targets[nums[2]*2-nums[1]*2-target] += 1
            targets[nums[2]*2-nums[1]*2+target] += 1
            targets[nums[2]*2-nums[1]*2+nums[0]*2-target] += 1
            targets[nums[2]*2-nums[1]*2+nums[0]*2+target] += 1

            count += targets[seed]
            results = [seed]
            for i in range(len(nums)-2, 2, -1):
                newResults = []
                for r in results:
                    result = nums[i]*2 - r
                    count += targets[result]
                    newResults.append(result)

                results.extend(newResults)

            return count


class Solution(object):
    def findTargetSumWays(self, nums, target):
        prefix = 0
        for i in range(len(nums)):
            nums[i] += prefix
            prefix = nums[i]

        def runner(sums, targetsCount):
            count = 0
            seed = sums[-1]
            count += targetsCount[seed]
            results = [seed]
            for i in range(len(sums)-2, -1, -1):
                newResults = []
                for r in results:
                    result = sums[i]*2 - r
                    count += targetsCount[result]
                    newResults.append(result)

                results.extend(newResults)

            return count

        def divisor(sums, depth):
            # return: targets
            targetsCount = defaultdict(int)
            targetsCount[target] += 1
            targetsCount[-target] += 1
            for i in range(depth):
                newDict = targetsCount.copy()
                for k, v in targetsCount.items():
                    # how many newDict[k], how many newDict[sums[i]*2-k]
                    # yes... this is
                    newDict[sums[i]*2-k] += v

                targetsCount = newDict

            return targetsCount

        depth = 8
        if len(nums) <= depth:
            targetsCount = divisor(nums, 0)
            return runner(nums, targetsCount)
        else:
            depth = len(nums)//2
            targetsCount = divisor(nums, depth)
            return runner(nums[depth:], targetsCount)


"""
Runtime: 119 ms, faster than 97.08% of Python3 online submissions for Target Sum.
Memory Usage: 14.1 MB, less than 87.50% of Python3 online submissions for Target Sum.
"""

"""
so they solved like this 

   def findTargetSumWays(self, A, S):
        count = collections.Counter({0: 1})
        for x in A:
            step = collections.Counter()
            for y in count:
                step[y + x] += count[y]
                step[y - x] += count[y]
            count = step
        return count[S]

it is just do the math step by step and keep count of reappering sums..
why this is not 2^n...
because some sum will collide and it will not really end up 2^n.. but subject to the total of different sum
    ^^ yeah.. after reading a few articles..
        the total variations of sum can be at most -SUM to +SUM..
        so it is not going to blow off..
        what I thought about was something else... but still fun..


cont reading 
https://leetcode.com/problems/target-sum/discuss/455024/DP-IS-EASY!-5-Steps-to-Think-Through-DP-Questions.

hmm.. interesting
https://leetcode.com/problems/target-sum/discuss/97334/Java-(15-ms)-C%2B%2B-(3-ms)-O(ns)-iterative-DP-solution-using-subset-sum-with-explanation

    public int subsetSum(int[] nums, int s) {
        int[] dp = new int[s + 1]; 
        dp[0] = 1;
        for (int n : nums)
            for (int i = s; i >= n; i--)
                dp[i] += dp[i - n]; 
        return dp[s];
    } 

this is kind of like my solution
structure-wise.. but not truly.. shit.. this is that one row subset sum.. I have solved a few times

ok.. this is actually the subset sum...
Find a subset P of nums such that sum(P) = (target + sum(nums)) / 2
                                        ^^ this is very good!



"""


if __name__ == '__main__':
    s = Solution()
    print(s.findTargetSumWays([1]*5, 3))
