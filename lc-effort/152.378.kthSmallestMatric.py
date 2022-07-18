"""
https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/

not much idea..
but maybe this is a dijstra? bfs+min-heap

"""


import heapq
from logging import makeLogRecord


class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        # 1 find the smallest number, must be in 1st column
        n = len(matrix)
        minNum = 1000000000
        minX = 0
        for r in range(n):
            if matrix[r][0] < minNum:
                minNum = matrix[r][0]
                minX = r

        # 2 BFS with min-heap
        h = []
        heapq.heappush(h, (minNum, minX, 0))  # val, x, y

        while k > 0:
            node = heapq.heappop(h)
            if node[1] > 0 and node[0] <= matrix[node[1]-1][node[2]]:
                heapq.heappush(h,
                               (matrix[node[1]-1][node[2]], node[1]-1, node[2]))
            if node[1] + 1 < n and node[0] <= matrix[node[1]+1][node[2]]:
                heapq.heappush(h,
                               (matrix[node[1]+1][node[2]], node[1]+1, node[2]))
            if node[2] + 1 < n and node[0] <= matrix[node[1]][node[2]+1]:
                heapq.heappush(h,
                               (matrix[node[1]][node[2]+1], node[1], node[2]+1))

        return node[0]


# let me pause this version here and keep it
# because I quickly see what is has gone wrong
# if the matrix is like
"""
1       2       3
100     200     300
4       5       6

the view into 4 5 6 row will be blocked 
so instead of finding a minimum. just add all 3 into the heap.. thus it opens up access to all elements

"""


class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        h = []
        n = len(matrix)
        marker = -1000000001
        for r in range(n):
            heapq.heappush(h, (matrix[r][0], r, 0))  # val, x, y
            matrix[r][0] = marker

        hs = set()  # tuple(x,y)

        while k > 0:
            node = heapq.heappop(h)
            if node[1] > 0 and node[0] <= matrix[node[1]-1][node[2]]:
                heapq.heappush(h,
                               (matrix[node[1]-1][node[2]], node[1]-1, node[2]))
                matrix[node[1]-1][node[2]] = marker
            if node[1] + 1 < n and node[0] <= matrix[node[1]+1][node[2]]:
                heapq.heappush(h,
                               (matrix[node[1]+1][node[2]], node[1]+1, node[2]))
                matrix[node[1]+1][node[2]] = marker
            if node[2] + 1 < n and node[0] <= matrix[node[1]][node[2]+1]:
                heapq.heappush(h,
                               (matrix[node[1]][node[2]+1], node[1], node[2]+1))
                matrix[node[1]][node[2]+1] = marker

            k -= 1

        return node[0]


"""
Runtime: 217 ms, faster than 77.91% of Python online submissions for Kth Smallest Element in a Sorted Matrix.
Memory Usage: 17.5 MB, less than 50.59% of Python online submissions for Kth Smallest Element in a Sorted Matrix.

not bad...
good read
https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/discuss/85173/Share-my-thoughts-and-Clean-Java-Code

Build a minHeap of elements from the first row.
Do the following operations k-1 times :
Every time when you poll out the root(Top Element in Heap), you need to know the row number and column number of that element(so we can create a tuple class here), replace that root with the next element from the same column.

yes... I don't have to look to other column..
because the min is still maintained and the next column in another row must be still greater than the first column in that row..
so they cannot be the smallest

so it can truly become O(n) memory.. while my solution is actually kind of between O(n) and O(n^2)
actually not, I only push also once.. because after push it... I marked it.. so yeah... but still... extra memory

let me do that..
"""


class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """
        h = []
        n = len(matrix)

        for r in range(n):
            heapq.heappush(h, (matrix[r][0], r, 0))  # val, x, y

        while k > 0:
            node = heapq.heappop(h)
            if node[2] + 1 < n and node[0] <= matrix[node[1]][node[2]+1]:
                heapq.heappush(h,
                               (matrix[node[1]][node[2]+1], node[1], node[2]+1))

            k -= 1

        return node[0]


"""
so I read this

public class Solution {
    public int kthSmallest(int[][] matrix, int k) {
        int lo = matrix[0][0], hi = matrix[matrix.length - 1][matrix[0].length - 1] + 1;//[lo, hi)
        while(lo < hi) {
            int mid = lo + (hi - lo) / 2;
            int count = 0,  j = matrix[0].length - 1;
            for(int i = 0; i < matrix.length; i++) {
                while(j >= 0 && matrix[i][j] > mid) j--;
                count += (j + 1);
            }
            if(count < k) lo = mid + 1;
            else hi = mid;
        }
        return lo;
    }
}

it is binary search, for sure... but will it be faster?
each time.. it searches from right to left. to find the elemented that is smaller or equal than/to the mid

the minheap is grow left to right..
this shrink right to left..

sitll O(K) time each time.. 
I feel this is actually slower???

let me try

and shit, when I am implement, I begin to have question why started with the same column from last row

Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth smallest element in the matrix.

I missed this important information
each row is sorted 
each column is sorted

"""


def bSearch(nums, target):
    # look for how many numbers are smaller or equal to me
    # in other words, find the first number bigger than me
    # it can all numbers smalle than or eqaul to me, so the space includes len(nums)
    l, r = 0, len(nums)
    while l < r:
        m = l + (r-l)//2
        if nums[m] <= target:
            l = m + 1
        else:
            r = m

    return l


class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """

        n = len(matrix)
        l, r = matrix[0][0], matrix[n-1][n-1]

        """
        some thinking is due here
        why l is the return value, because m is not necessary existent in the matrix everytime
        so it can be proved by contradiction 

        what this algorithm does is to find the smallest number  that has k number <= it
        if we say l is not in the matrix and there must be a l2 in the matrix 
        that we say l2 is the smaller number which exist in the matrix and also k number <= it,
         the l should have k+1 number <= it

        which contrdicts above... so l has to be in the matrix 
        """
        while l < r:
            m = l + (r-l)//2
            i, j = 0, n-1
            count = 0
            for i in range(n):
                """
                what is smart here:
                j is the result of last row
                but my solution actually discarded that insights.. 
                hahah... I thought I was smart to use binary search.. but instead I discarded 
                previous calculations.. 

                let me redo..
                """
                j = bSearch(matrix[i], m)
                count += j

            if count < k:
                l = m+1
            else:
                r = m
        return l


"""
Runtime: 149 ms, faster than 95.78% of Python online submissions for Kth Smallest Element in a Sorted Matrix.
Memory Usage: 17.4 MB, less than 70.66% of Python online submissions for Kth Smallest Element in a Sorted Matrix.

wow....
"""


class Solution(object):
    def kthSmallest(self, matrix, k):
        """
        :type matrix: List[List[int]]
        :type k: int
        :rtype: int
        """

        n = len(matrix)
        l, r = matrix[0][0], matrix[n-1][n-1]

        while l < r:
            m = l + (r-l)//2
            i, j = 0, n-1
            count = 0
            for i in range(n):
                while j >= 0 and matrix[i][j] > m:
                    j -= 1
                count += j+1

            if count < k:
                l = m+1
            else:
                r = m
        return l


"""
Runtime: 181 ms, faster than 90.22% of Python online submissions for Kth Smallest Element in a Sorted Matrix.
Memory Usage: 17.6 MB, less than 50.59% of Python online submissions for Kth Smallest Element in a Sorted Matrix.

hahah... pratically my row-binary-search is faster?
not making

for each m, that is O(n*lgn), this is O(n)
"""

if __name__ == '__main__':
    s = Solution()
    m = [[1, 5, 9], [10, 11, 13], [12, 13, 15]]
    print(s.kthSmallest(m, 8))
    m = [[1, 5, 9], [10, 11, 13], [12, 13, 15]]
    print(s.kthSmallest(m, 7))
    m = [[1, 5, 9], [10, 11, 13], [12, 13, 15]]
    print(s.kthSmallest(m, 6))
