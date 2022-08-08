
"""
https://leetcode.com/problems/shortest-subarray-with-sum-at-least-k/

at first sight, this is same to https://leetcode.com/problems/minimum-size-subarray-sum/
but it has negative

I don't see why negative could impact the results.. then I ran it
it passed 61 tests and failed this one
[84,-37,32,40,95]
167
expect 3. output 5

at first sight.. this doesn't make sense.. why it can fail..
let me find out what is special about this


I see, when S-=84, the -37 becomes a burden
so first thinking is to skip all negative as well... when shrinking by doing this 
            while S >= k or (i<len(nums) and nums[i] <= 0):
                if S >= k:
                    res = min(j-i+1, res)

then it fails 
[45,95,97,-34,-42]
21

so negative and zeroes does complicate things a lot

one thing I can do is to add 10^5+1 to make it pure positive
but then you need to offset the sum part.. it will be added windowLen*10^5+1

but then the target turns into a moving target so it is not gonna work..
what it can do is turn the presum into a fixed target...

still not gonna cut.. pause...

how about I just get rid of all the negative in the back cause they don't contribute to nothing
also the largest of a [neg,pos].. can be replaced by a neg+pos..
why I am doing this is because 1 <= k <= 10^9, so I do need to stay in the positive territory



"""


from typing import List


class Solution:
    def shortestSubarray(self, nums: List[int], k: int) -> int:
        # find last positive integer
        lastPostive = len(nums)-1
        for i in range(len(nums)-1, -1, -1):
            if nums[i] > 0:
                break
            lastPostive = i-1

        nums = nums[:lastPostive+1]
        i, j = 0, 0
        res = len(nums)+1
        S = 0

        while j < len(nums):
            S += nums[j]
            if S <= 0:
                while nums[j] <= 0:
                    j += 1
                S = nums[j]
                i = j
            while i <= j and (S >= k or nums[i] <= 0):
                if S >= k:
                    res = min(j-i+1, res)
                S -= nums[i]
                i += 1
            j += 1

        posSum = 0
        i, j = 0, 0
        while j < len(nums):
            posSum += nums[j]
            if nums[j] <= 0:
                posSum = 0
                i = j

            while i <= j and (posSum >= k or nums[i] <= 0):
                if posSum >= k:
                    res = min(j-i+1, res)
                if nums[i] > 0:
                    posSum -= nums[i]
                i += 1
            j += 1

        if res == len(nums) + 1:
            return -1
        return res


"""
failed here
[18,83,-43,-18,60]
18

output -2.. shit... apparently this is 1

okay.. just let i<=j, don't let i go over j..

then failed here 4 vs 2
[-34,37,51,3,-12,-50,51,100,-47,99,34,14,-13,89,31,-14,-44,23,-38,6]
151

okay.. I do twice 
once i do the general sum.. 2nd time I do the positive only
then it failed here

[-36,10,-28,-42,17,83,87,79,51,-26,33,53,63,61,76,34,57,68,1,-30]
484

then skipped negative presum or numbers
then failed here
[58701,23101,6562,60667,20458,-14545,74421,54590,84780,63295,33238,
-10143,-35830,-9881,67268,90746,9220,-15611,23957,29506,-33103,-14322,
19079,-34950,-38551,51786,-48668,-17133,5163,15122,5463,74527,41111,
-3281,73035,-28736,32910,17414,4080,-42435,66106,48271,69638,14500,37084,
-9978,85748,-43017,75337,-27963,-34333,-25360,82454,87290,87019,84272,
17540,60178,51154,19646,54249,-3863,38665,13101,59494,37172,-16950,-30560,
-11334,27620,73388,34019,-35695,98999,79086,-28003,87339,2448,66248,81817,
73620,28714,-46807,51901,-23618,-29498,35427,11159,59803,95266,20307,-3756,
67993,-31414,11468,-28307,45126,77892,77226,79433]
1677903
54 vs 48..

okay... I think my random trial will not cut it...
need to really think now

NOT SOLVED YET
"""

if __name__ == '__main__':

    s = Solution()
    print(s.shortestSubarray([-36, 10, -28, -42, 17, 83, 87, 79,
          51, -26, 33, 53, 63, 61, 76, 34, 57, 68, 1, -30], 484))

    print(s.shortestSubarray([-34, 37, 51, 3, -12, -50, 51, 100, -
          47, 99, 34, 14, -13, 89, 31, -14, -44, 23, -38, 6], 151))
    print(s.shortestSubarray([18, 83, -43, -18, 60], 18))
    print(s.shortestSubarray([84, -37, 32, 40, 95], 167))
    print(s.shortestSubarray([45, 95, 97, -34, -42], 21))
