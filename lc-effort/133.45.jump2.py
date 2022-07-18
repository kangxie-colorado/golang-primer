"""
https://leetcode.com/problems/jump-game-ii/

O(n^2) is visble to me... maynot pass but let me try
"""


from ast import List
import heapq
from tkinter import N


class Solution_n_square:
    def jump(self, nums) -> int:
        nums[len(nums)-1] = 0
        for i in range(len(nums)-2, -1, -1):
            nums[i] = 1 + min(nums[i+1:min(len(nums), i+nums[i]+1)]
                              ) if nums[i] != 0 else 1000

        return nums[0]


"""
Runtime: 1297 ms, faster than 31.69% of Python3 online submissions for Jump Game II.
Memory Usage: 14.7 MB, less than 99.29% of Python3 online submissions for Jump Game II.

alright not as bad as I thought..

then I see the dijstra solution looming around..
but I need to priority queue to prioritize on the idx+jump value
https://stackoverflow.com/questions/3954530/how-to-make-heapq-evaluate-the-heap-off-of-a-specific-attribute
this mentions using a tuple and let it sort on the first value.. then cool
let me repl a bit

>>> import heapq
>>> q=[]
>>> heapq.heappush(q, (1,2,3))
>>> heapq.heappush(q, (-1,2,3))
>>> q
[(-1, 2, 3), (1, 2, 3)]
"""


class Solution_dijstra:
    def jump(self, nums) -> int:
        # tuple: (-abs(idx+nums[idx]), idx)
        # tuple[0] used for sorting
        # tuple[1] used for bfs
        # using -abs for max heap
        q = []
        # just to make sure when it actually touches the end, it prioritize
        nums[len(nums)-1] = 1001

        heapq.heappush(q, (-(0+nums[0]), 0))
        steps = 0
        while True:
            n = heapq.heappop(q)
            if n[1] == len(nums) - 1:
                return steps

            # careful, you can at most jump to len(nums)-1
            # so j can at most be len(nums)-1.. and don't mix it with the idx+nums[idx]
            # ugh.. I am so easy to make mistakes..
            for j in range(n[1]+1, min(len(nums), n[1]+nums[n[1]]+1)):
                heapq.heappush(q, (-(j+nums[j]), j))

            steps += 1


"""
Runtime: 272 ms, faster than 42.30% of Python3 online submissions for Jump Game II.
Memory Usage: 15.3 MB, less than 12.59% of Python3 online submissions for Jump Game II.

okay.. think this is first tier

Runtime: 145 ms, faster than 86.93% of Python3 online submissions for Jump Game II.
Memory Usage: 15.7 MB, less than 9.22% of Python3 online submissions for Jump Game II.

and.. there is an O(n) solution...
using the facts that the BFS is naturally extending to 1 direction
kinda like the frontline extension.. 
step1: start->end ===> next start:end+1 -> next end:the furthest reach
"""


class Solution:
    def jump(self, nums) -> int:
        start = 0
        end = 1
        furthest = 1
        steps = 0
        while end < len(nums):
            steps += 1
            for i in range(start, end):
                furthest = max(furthest, i+nums[i])
                if furthest >= len(nums)-1:
                    return steps
            start, end = end, furthest+1

        return steps


"""
[1,2,3] TLE..
ugh... so painful I cannot convert idea into code correctly on fewest tries
had to debug 
probably still not walked thru enough examples

best still stick to left close and right open [i:j) and make adjustment around to do this...

hmm..

Runtime: 289 ms, faster than 40.33% of Python3 online submissions for Jump Game II.
Memory Usage: 15.1 MB, less than 58.61% of Python3 online submissions for Jump Game II.

why

"""

if __name__ == '__main__':
    s = Solution()
    print(s.jump([9]))
    print(s.jump([]))
    print(s.jump([1, 2, 3]))
    print(s.jump([9, 3, 1, 1, 4]))
    print(s.jump([2, 3, 1, 1, 4]))
    print(s.jump([2, 3, 0, 1, 4]))
