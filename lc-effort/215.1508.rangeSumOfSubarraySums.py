"""
https://leetcode.com/problems/range-sum-of-sorted-subarray-sums/

this looks so desparate at first
then I draw a matrix
1 2 3 4
  0 1 2 3
0 
1
2
3

m[r][c] is the sum of nums[i:j+1] [i,j] inclusive
then notice this can now become a heap 
because m[0][0] < m[0][1] < m[0][2] < m[0][3]
and m[0][3] > m[1][3] > m[2][3] > m[3][3]

so the diagonal line will be the initial value of the candidates, put them into a min heap
then pop up the smallest one.. next one to enter will be either one row above or one column right

m[i][j], i can only decrease to 0; j can only increase to n-1
    get to above, the sum is add the previous number;
    get to right, the sum is add the right number
and of course, when m[i][j] is taken by another predesor, add another direction
so we need a map to mark it...

"""


import heapq
import sys
from typing import List


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        h = []  # use a min heap (sum, i, j)
        used = [[]]*n
        for i in range(n):
            used[i] = [0]*n
            heapq.heappush(h, (nums[i], i, i))
            used[i][i] = 1

        res = []
        while len(res) < right:
            s, i, j = heapq.heappop(h)
            res.append(s)

            if i > 0 and used[i-1][j] == 0:
                heapq.heappush(h, (s+nums[i-1], i-1, j))
                used[i-1][j] = 1
            if j < n-1 and used[i][j+1] == 0:
                heapq.heappush(h, (s+nums[j+1], i, j+1))
                used[i][j+1] = 1

        return sum(res[left-1:])


"""
53 / 54 test cases passed.
oh.. forget to mod 10^9+7
"""


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        h = []  # use a min heap (sum, i, j)
        used = [[]]*n
        for i in range(n):
            used[i] = [0]*n
            heapq.heappush(h, (nums[i], i, i))
            used[i][i] = 1

        res = []
        while len(res) < right:
            s, i, j = heapq.heappop(h)
            res.append(s)

            if i > 0 and used[i-1][j] == 0:
                heapq.heappush(h, (s+nums[i-1], i-1, j))
                used[i-1][j] = 1
            if j < n-1 and used[i][j+1] == 0:
                heapq.heappush(h, (s+nums[j+1], i, j+1))
                used[i][j+1] = 1

        return sum(res[left-1:]) % (10**9+7)


"""
Runtime: 594 ms, faster than 52.32% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 45.5 MB, less than 8.36% of Python3 online submissions for Range Sum of Sorted Subarray Sums.

aha.. I thought it should be muck quicker..
let me see the brute force
"""


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        sums = []

        for i in range(n):
            pathSum = 0
            pathSums = [0]*(n-i)
            for j in range(i, n):
                pathSum += nums[j]
                pathSums[j-i] = pathSum
            sums.extend(pathSums)

        sums.sort()
        return sum(sums[left-1:right]) % (10**9+7)


"""
Runtime: 370 ms, faster than 92.26% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 37.4 MB, less than 30.65% of Python3 online submissions for Range Sum of Sorted Subarray Sums.

what the fuck..
"""


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        h = []  # use a min heap (sum, i, j)
        used = [[]]*n
        for i in range(n):
            used[i] = [0]*n
            heapq.heappush(h, (nums[i], i, i))
            used[i][i] = 1

        res = [0]*right
        idx = 0
        while idx < right:
            s, i, j = heapq.heappop(h)
            res[idx] = s
            idx += 1

            if j < n-1 and used[i][j+1] == 0:
                heapq.heappush(h, (s+nums[j+1], i, j+1))
                used[i][j+1] = 1

        return sum(res[left-1:]) % (10**9+7)


"""
Runtime: 1197 ms, faster than 19.81% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 45.3 MB, less than 8.36% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
"""

"""
just going right, no need to go up..
because it can be taken care by another direction...
"""


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        h = []  # use a min heap (sum, i, j)
        for i in range(n):
            heapq.heappush(h, (nums[i], i, i))

        res = [0]*right
        idx = 0
        while idx < right:
            s, i, j = heapq.heappop(h)
            res[idx] = s
            idx += 1

            if j < n-1:
                heapq.heappush(h, (s+nums[j+1], i, j+1))

        return sum(res[left-1:]) % (10**9+7)


"""
Runtime: 480 ms, faster than 70.90% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 37.4 MB, less than 44.89% of Python3 online submissions for Range Sum of Sorted Subarray Sums.

Runtime: 453 ms, faster than 75.54% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 37.4 MB, less than 44.89% of Python3 online submissions for Range Sum of Sorted Subarray Sums.

better but still not as good as the binary search
study that tomorrow

okay
three/two ideas put together
1. sliding window get countAndSumTotal for all subarray sums that are smaller/equal to a target-sum
    bla..bla..
    this seems to be up by one dimension..
    but think on a plain array, what you use binary search for
    bla..bla..

2. binary search use the above sliding window to fingure out the rightful sum that just sits at the kth slot
    after you get l, the l is a target-sum, and it could be sitting in a repeating sequence 
    so you need to countAndSum once again.. and remove the possible duplicates (count-k)*target
    so l is target sum that sits at kth slot.. but k-1,k,k+1.. could all equal to l...
    so do this one more time to rid of latter duplicates

3. the answer is sum_k_sums(right) - sum_k_sums(left-1)
"""

"""
three layer of sums... this is really complicated 
so lee computes B/C... and this is why

ending with j-1, has a path to ending with j : += nums[j] * (j-i+1)
starting with i, has a path to starting with i+1 : -= currWindSum
"""


def countUnderScore(nums, target):
    i, j = 0, 0
    currWindSum = 0
    totalSum = 0
    sumEndingJ = 0
    count = 0
    while j < len(nums):
        currWindSum += nums[j]
        sumEndingJ += nums[j]*(j-i+1)
        while currWindSum > target:
            sumEndingJ -= currWindSum
            currWindSum -= nums[i]
            i += 1

        totalSum += sumEndingJ
        count += j-i+1
        j += 1
    return count, totalSum


# if duplicates exist, will it converge?
"""
using [1, 2, 3, 4]
the sorted sub-array-sum should be [1, 2, 3, 3, 4, 5, 6, 7, 9, 10]

    print(sumKSums([1, 2, 3, 4], 3))    => 3
    is it right?

I know this
    print(countUnderScore([1, 2, 3, 4], 3)) => (4,9)
    so this function will calculate 4 <=3 subarray-sums..
    however, it is still able to find out it is the 3 (target-sum) having 3<= sums..
    so it kind of right.. because if it returns 4, it should be 5<=
    if it returns 2, it should be 2<=
    so yeah.. either on the exact hit or l/r converge, it is still right

yeah, I know another form which takes the inteligence of the repeating into account
        if count >= k:
            # >: it is possible m is the one because of repeating
            # =: of course eligible
            r = m
        else:
            l = m+1
"""


def sumKSums_figure_out_bst_part(nums, k):
    l = min(nums)
    r = sys.maxsize

    while l < r:
        m = l+(r-l)//2
        count = countUnderScore(nums, m)[0]
        # if duplicates exist, will it converge?

        if count == k:
            break
        elif count < k:
            l = m+1
        else:
            r = m-1

    return l


def sumKSums(nums, k):
    l = min(nums)
    r = sys.maxsize

    while l < r:
        m = l+(r-l)//2
        count = countUnderScore(nums, m)[0]
        # if duplicates exist, will it converge?

        if count == k:
            # this is another nuance for this form
            # this will be m to do the final countAndSum
            # if converge to end, it will be l to do
            l = m
            break
        elif count < k:
            l = m+1
        else:
            # r = m-1 # this is wrong, m could be in repeating zone and still eligible
            r = m

    count, S = countUnderScore(nums, l)

    return S - (count-k)*l


class Solution:
    def rangeSum(self, nums: List[int], n: int, left: int, right: int) -> int:
        return (sumKSums(nums, right) - sumKSums(nums, left-1)) % (10**9+7)


"""
[6,6,2,6,4]
5
1
3

failed here..
I suspect the binary search..
        else:
            r = m-1
            ^ should be r = m 
                because when it is in the repeating zones.. m is still possible the end resuls
                so indeed the other form is better
                change it and this case passed.. let me see

Runtime: 520 ms, faster than 63.78% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 14.1 MB, less than 90.71% of Python3 online submissions for Range Sum of Sorted Subarray Sums.

cool.. cool

Runtime: 428 ms, faster than 81.11% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
Memory Usage: 13.9 MB, less than 97.21% of Python3 online submissions for Range Sum of Sorted Subarray Sums.
"""

if __name__ == '__main__':
    print(sumKSums_figure_out_bst_part([1, 2, 3, 4], 3))

    print(countUnderScore([1, 2, 3, 4], 1))
    print(countUnderScore([1, 2, 3, 4], 5))

    s = Solution()
    print(s.rangeSum([1, 2, 3, 4], 4, 1, 5))
    print(s.rangeSum([1, 2, 3, 4], 4, 3, 4))
    print(s.rangeSum([1, 2, 3, 4], 4, 1, 10))
