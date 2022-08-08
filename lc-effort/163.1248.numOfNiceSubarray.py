"""
https://leetcode.com/problems/count-number-of-nice-subarrays/

this is indeed a sliding window..
so if it is qualifying window or under-qualifying.. slide right.. 
over-qualifying... slide left.. 

but when do the counting, adding the i+1. looking to the left to avoid the tail effect
"""


from ast import List


class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        i, j = 0, 0
        oddNums = 0
        count = 0
        lastOddIdx = -1
        oddQ = []
        while j < len(nums):
            if nums[j] % 2 == 1:
                oddNums += 1
                oddQ.append(j)
            while oddNums > k:
                if nums[i] % 2 == 1:
                    oddNums -= 1
                    lastOddIdx = i
                    oddQ = oddQ[1:]
                i += 1

            if oddNums == k:
                count += oddQ[0]-lastOddIdx

            j += 1

        return count


"""
Runtime: 4846 ms, faster than 5.06% of Python3 online submissions for Count Number of Nice Subarrays.
Memory Usage: 20.6 MB, less than 68.67% of Python3 online submissions for Count Number of Nice Subarrays.


then I kind of see another sliding window solution
extraction the odd number index.. then looking to left.. looking to right..
e..g
[1,2,2,2,1,1,1,2,2] k=2
I got the odd index to be 
[0,4,5,6]

so now I just do fix width window sliding
[0,4] there will be an imaginary left boundary -1
so looking to left... 0-1
looking to right.. the next odd idx is 5, so 5-4 
so (0-(-1))*(5-4) = 1

then [4,5]
looking to left 4-0 = 4
looking right 6-5 = 1 
4*1 = 4

then [5,6]
looking to left 5-4 = 1
looking to right, there is no more, so use imaginary 9, 9-6=3
1*3=3

1+4+3 = 8
let me test .. and that is right



"""


class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        oddIdxs = [-1]
        for i in range(len(nums)):
            if nums[i] % 2 == 1:
                oddIdxs.append(i)
        oddIdxs.append(len(nums))

        # start with k, idx-1 to idx-k, makes it a k-size window..
        count = 0
        for j in range(k, len(oddIdxs)-1):
            i = j-k+1
            count += (oddIdxs[i]-oddIdxs[i-1]) * (oddIdxs[j+1]-oddIdxs[j])

        return count


"""
Runtime: 961 ms, faster than 85.12% of Python3 online submissions for Count Number of Nice Subarrays.
Memory Usage: 20.7 MB, less than 68.67% of Python3 online submissions for Count Number of Nice Subarrays.

wow hoo.

then there is this Lee's god code..
atMost(k) - atMost(k-1)
"""


class Solution:
    def numberOfSubarrays(self, nums: List[int], k: int) -> int:
        def atMost(t):
            i, j = 0, 0
            count = 0
            oddNum = 0
            while j < len(nums):
                oddNum += nums[j] % 2
                while oddNum > t:
                    oddNum -= nums[i] % 2
                    i += 1

                count += j-i+1
                j += 1
            return count

        return atMost(k) - atMost(k-1)
