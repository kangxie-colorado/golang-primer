# https://leetcode.com/problems/find-median-from-data-stream/
# so if you figure out the algorithm this is not super difficult
# the thing is the is ordered list of integer so the number increases
# so actually only two number are needed to know the media
# think 1 2 3 4 5 6
# I need to quickly know 3 and 4
# naturally this is heap problem
# save first half to a max heap [3 2 1]
# save 2nd half to a min heap [4 5 6]
# when two heap are same size, take average of both's top number
# when one heap is bigger, take its own top... and important is to maintain the size diff to be only 1
# so when 7 comes... it will be added to 2nd half's min heap [4 5 6 7]
# when 8 comes, the heaps are not same size, so pop 4 out of min heap and add it to max heap; so heaps becomes 
# [4 3 2 1] and [5 6 7 8].. we still got the median

# python's heapq naturally implements the min-heap
# but there is no built-in max-heap... and the standard workaround is to negate the values (since this is integer)
import heapq
class MinHeap:
    def __init__(self, l=[]):
        self.l = l
    
    def __len__(self):
        return len(self.l)

    def top(self):
        return self.l[0]

    def push(self, n):
        heapq.heappush(self.l, n)
    
    def pop(self):
        return heapq.heappop(self.l)

class MaxHeap:
    def __init__(self, l=[]):
        self.l = l
    
    def __len__(self):
        return len(self.l)
    
    def top(self):
        return -self.l[0]

    def push(self, n):
        heapq.heappush(self.l, -n)
    
    def pop(self):
        return -heapq.heappop(self.l)

class MedianFinder:
    def __init__(self):
        self.maxH = MaxHeap() # lower half
        self.minH = MinHeap() # higher half
        self.median = 0
        self.increasing = True

    def _1_addNum(self, num: int) -> None:
        # this assumed the nums are either increaasing or decreasing
        # which ain't true
        if len(self.minH) == len(self.maxH):
            self.minH.push(num)
            self.median = self.minH.top()
        else:
            mintop = self.minH.pop()
            if mintop > num and self.increasing:
                self.minH, self.maxH = self.maxH, self.minH
                self.increasing = False

            self.maxH.push(mintop)
            self.minH.push(num)
            self.median = (self.minH.top()+self.maxH.top())/2

    def addNum(self, num):
        if len(self.minH) and num > self.minH.top():
            self.minH.push(num) 
        else:
            self.maxH.push(num)

        if abs(len(self.minH) - len(self.maxH))>1:
            # need a rebanlance
            if len(self.minH) > len(self.maxH):
                self.maxH.push(self.minH.pop())    
            else:
                self.minH.push(self.maxH.pop())

    def findMedian(self) -> float:
        if len(self.minH) == len(self.maxH):
            self.median = (self.minH.top()+self.maxH.top())/2
        else:
            if len(self.maxH) > len(self.minH):
                self.median = self.maxH.top()
            else:
                self.median = self.minH.top()

        return self.median

def testHeap():
    # maxH = MaxHeap()
    maxH = MinHeap()
    maxH.push(3)
    maxH.push(1)
    maxH.push(4)
    maxH.push(2)

    print(maxH.top())
    print(maxH.pop())
    print(maxH.top())
    maxH.push(1)
    maxH.push(4)

    while len(maxH):
        print(maxH.pop())

def main():
    # case of increasing
    # median_finder = MedianFinder()
    # median_finder.addNum(1)
    # print(median_finder.findMedian())
    # median_finder.addNum(3)
    # print(median_finder.findMedian())

    # case of decreasing 
    # ["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
    # [[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]
    median_finder = MedianFinder()
    median_finder.addNum(-1)
    print(median_finder.findMedian())
    median_finder.addNum(-2)
    print(median_finder.findMedian())
    median_finder.addNum(-3)
    print(median_finder.findMedian())
    median_finder.addNum(-4)
    print(median_finder.findMedian())
    median_finder.addNum(-5)
    print(median_finder.findMedian())


if __name__ == '__main__':
    main()

"""
["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[-1],[],[-2],[],[-3],[],[-4],[],[-5],[]]

failed.. shit, you said it is increasing order??
so it is ordered but not necessary increasing and could be decreasing...

okay.. then failed here

["MedianFinder","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian","addNum","findMedian"]
[[],[6],[],[10],[],[2],[],[6],[],[5],[],[0],[],[6],[],[3],[],[1],[],[0],[],[0],[]]

okay.. fuck it is not ordered at all
I just need to figure out the side to add and be flexiable about it
"""


"""
Runtime: 1089 ms, faster than 23.46% of Python3 online submissions for Find Median from Data Stream.
Memory Usage: 35.9 MB, less than 69.24% of Python3 online submissions for Find Median from Data Stream.


"""