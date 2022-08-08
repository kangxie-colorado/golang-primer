"""
https://leetcode.com/problems/build-array-from-permutation/

of course there is a O(N) memory solution 
but it wants a O(1) memory solution

then it can only be done by swapping
I think I did some similar problems before... by following a cycle..
but it is still kind of hard to figure it out

[0,2,1,5,3,4]

there are 3 cycles..
0 itself 
2 1 
5 3 4

use 5 3 4 for example
cycleStartIdx=i=3
nextIdx=nums[i]=5 -> save t=5, nums[i] (nums[3]) = 4 (nums[nextIdx])
i = nextIdx = 5
-
nextIdx = nums[i]=4 -> nums[i]=nums[nextIdx] => nums[5] = nums[4] = 3
i = nextIdx = 4
-
i=4
nextIdx=3
nextIdx hits cycleStart.. so pull in the t value, nums[4] = 5
it will end having 4 5 3

problem is how do I know this cycle is sorted now I do O(1) memory?
so after sorting out a cycle
let me see.. 
0 2 1

after sorting 2 1
it will be 0 1 2
i will be at idx-2 (2)
cycleStart will be 2... next idx will 2... so it naturally ends
but how about 5 3 4
it will become 4 5 3... i will be at 4
cycleStartIdx=4 naturally
nextIdx =3 .... hmmm... it looks like it is setting stage for another round of shifting????

maybe the max i, this cycle encounters should be the start of next cycle (max-i + 1 I mean)
nah..

6 0 1 5 3 4 2 7
      ^^^^^     <- this is a cycle
^^^^^       ^   <- tihs is another cycle

you see.. two cycles.. and first cycle is max 6

so the problem is how to know cycle has been ended here?
I recall the previous problem will end up like idx=val condition to end the cycle..
but here??? actually I cannot even find previous problme maybe I made it up???

anyway.. I am now stuck..
let me read
https://leetcode.com/problems/build-array-from-permutation/discuss/1315926/Python-O(n)-Time-O(1)-Space-w-Full-Explanation

that is a brilliant thing
then reading this
https://leetcode.com/problems/build-array-from-permutation/discuss/1316619/Python3-O(n)-Time-O(1)-Space-Using-FastSlow-Pointer-Method-From-LeetCode-287

oh... I totally forgot I can covert the number to another place to mark it sorted... 
like +len(nums) or make it negative...

then convert back... 
ouch.. I totally didn't come up with any of these ideas..

now I can code it up
"""


from typing import List


class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        for i in range(len(nums)):
            if nums[i] >= len(nums):
                continue

            cycleStart = i
            t = nums[cycleStart]
            nextI = nums[i]
            while nextI != cycleStart:
                # the order is pain
                # 1. save up the to-modify index in last
                lastI = i
                # 2. modify nums[i] first time.. pull in the next value
                nums[i] = nums[nextI]
                # 3. update i and nextI
                i = nextI
                nextI = nums[i]
                # 4. mark it visited
                nums[lastI] += len(nums)
            nums[i] = t + len(nums)

        for i in range(len(nums)):
            nums[i] -= len(nums)
        return nums


"""
Runtime: 143 ms, faster than 80.01% of Python3 online submissions for Build Array from Permutation.
Memory Usage: 14.1 MB, less than 89.48% of Python3 online submissions for Build Array from Permutation.

did I save memory????
"""

"""
anyway... looking at this now
https://en.wikipedia.org/wiki/Euclidean_division

it is kind of like an encoding 

  # turn the array into a=qb+r
  for i in range(len(nums)):
	r = nums[i]
	
	# retrieve correct value from potentially already processed element
	# i.e. get r from something maybe already in the form a = qb + r
	# if it isnt already processed (doesnt have qb yet), that's ok b/c
	# r < q, so r % q will return the same value
	b = nums[nums[i]] % q
	
    # put it all together
	nums[i] = q*b + r

so turn nums[i] into q*b+r
some it retain its own information in r.. 
but pull its ought-to-be value/information in b.. 

also pay attention to 
    b = nums[nums[i]] % q
    ^ it is not plain nums[nums[i]]... because the value could be already processed...
    if nums[i] (lets say j) is processed, then nums[j] will be in q*b+r form
    we only need its raw information which is r...
    gosh.. why are people so smart

this is very cool... but I don't expect myself to be so smart
just a little fun to code it up though..

"""


class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        q = len(nums)
        for i, c in enumerate(nums):
            nums[i] += q * (nums[c] % q)
        for i in range(len(nums)):
            nums[i] //= q

        return nums


"""
Runtime: 138 ms, faster than 82.34% of Python3 online submissions for Build Array from Permutation.
Memory Usage: 14.1 MB, less than 89.48% of Python3 online submissions for Build Array from Permutation.

pretty good...
"""


if __name__ == '__main__':
    s = Solution()
    print(s.buildArray([0, 2, 1, 5, 3, 4]))
    print(s.buildArray([5, 0, 1, 2, 3, 4]))
