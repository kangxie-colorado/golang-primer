"""
https://leetcode.com/problems/divide-array-in-sets-of-k-consecutive-numbers/

this should be almost the same as 659
so hey.. so maybe if I ran into a problem 
I did some due diligence I cannot figure out, look at someone's solution
when do a related problem, I also learn enough...
"""


from typing import Counter, List


class Solution:
    def isPossibleDivide(self, nums: List[int], k: int) -> bool:
        if len(nums) % k != 0:
            return False

        unused = Counter(nums)
        for n in sorted(unused.keys()):
            if not unused[n]:
                continue

            for i in range(1, k):
                if not unused[n+i]:
                    return False
                unused[i+n] -= unused[n]
            unused[n] = 0

        return True


"""
ah..

the tricky part is
1. sort
2. must use all copies...

Runtime: 666 ms, faster than 52.72% of Python3 online submissions for Divide Array in Sets of K Consecutive Numbers.
Memory Usage: 31.5 MB, less than 59.89% of Python3 online submissions for Divide Array in Sets of K Consecutive Numbers.
"""

"""
another solution inspired by the consecutive lenghth of the the line.. 
forgot what that problem is?


but idea is to look for the starting point instead of sort
so it becomes O(n)
"""


class Solution:
    def isPossibleDivide(self, nums: List[int], k: int) -> bool:
        if len(nums) % k != 0:
            return False

        unused = Counter(nums)
        for n in nums:
            if not unused[n]:
                continue

            while unused[n-1]:
                n -= 1

            for i in range(k-1, -1, -1):
                if not unused[n+i]:
                    return False
                unused[i+n] -= unused[n]

        return all([not unused[i] for i in unused])


"""
huh.. cannot win so easy? can you?

even I have the grand idea, I still miss dealing with the necessary edges case
1. unused not used up
2. the k==1
[5,7,8,8,7,4,3,6]
1

so             while unused[n-1]:
                n -= 1
        needs a make up? to account for k...
        no cannot do
        think [2 3 4 5] k=2, when 4 is here.. go back 1 to 3.. and take it away... leaving 2 5... un attended
    so maybe 1 is just a special case

"""


class Solution:
    def isPossibleDivide(self, nums: List[int], k: int) -> bool:
        if k == 1:
            return True

        if len(nums) % k != 0:
            return False

        unused = Counter(nums)
        for n in nums:
            if not unused[n]:
                continue

            while unused[n-1]:
                n -= 1

            for i in range(k-1, -1, -1):
                if not unused[n+i]:
                    return False
                unused[i+n] -= unused[n]

        return all([not unused[i] for i in unused])


"""
Runtime: 428 ms, faster than 99.07% of Python3 online submissions for Divide Array in Sets of K Consecutive Numbers.
Memory Usage: 31.5 MB, less than 59.89% of Python3 online submissions for Divide Array in Sets of K Consecutive Numbers.

wonder how the smart man handles the edge cases...

def isPossibleDivide(self, nums, k):
    ctr = collections.Counter(nums)
    for num in nums:
        start = num
        while ctr[start - 1]:
            start -= 1
        while start <= num:
            while ctr[start]:
                for victim in range(start, start + k):
                    if not ctr[victim]:
                        return False
                    ctr[victim] -= 1
            start += 1
    return True
    anyway... he used the inner loop to take care of all start<=num... 
    ....
    anyway... coding like this need some edge case handling
"""

if __name__ == "__main__":
    s = Solution()
    #assert s.isPossibleDivide(nums=[1, 2, 3, 3, 4, 4, 5, 6], k=4)
    assert not s.isPossibleDivide(nums=[3, 2, 8, 1, 8, 5, 7, 6], k=2)
