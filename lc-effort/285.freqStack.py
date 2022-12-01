"""
class FreqStack:

    def __init__(self):
        self.maxHeap = []
        self.freqMap = defaultdict(int)

    def push(self, val: int) -> None:
        self.freqMap[val] += 1
        heapq.heappush(
            self.maxHeap, (-self.freqMap[val], -len(self.maxHeap), val))


    def pop(self) -> int:
        _,_,val = heapq.heappop(self.maxHeap)
        self.freqMap[val] -= 1
        return val

"""


from collections import defaultdict
import heapq


class FreqStack:

    def __init__(self):
        self.maxHeap = []
        self.freqMap = defaultdict(int)
        self.seq = 0

    def push(self, val: int) -> None:
        self.freqMap[val] += 1
        heapq.heappush(
            self.maxHeap, (-self.freqMap[val], -self.seq, val))
        self.seq += 1

    def pop(self) -> int:
        _, _, val = heapq.heappop(self.maxHeap)
        self.freqMap[val] -= 1
        return val


"""
Runtime: 416 ms, faster than 83.16% of Python3 online submissions for Maximum Frequency Stack.
Memory Usage: 22.7 MB, less than 63.86% of Python3 online submissions for Maximum Frequency Stack.


doesnt seem like a hard problem tbh

https://leetcode.com/problems/maximum-frequency-stack/discuss/163410/C%2B%2BJavaPython-O(1)
as always... Lee is one-up everyone...lol

his idea is to 
1. use freqMap to store the freqs and track
2. use m as the buckets: m[bucket] = [] # a stack, naturally ranked by the pos

let me also do this
"""


class FreqStack:

    def __init__(self):
        self.freqBuckets = defaultdict(list)  # frequencies buckets
        self.freqMap = defaultdict(int)
        self.maxFreq = 0

    def push(self, val: int) -> None:
        self.freqMap[val] += 1
        self.maxFreq = max(self.maxFreq, self.freqMap[val])
        self.freqBuckets[self.freqMap[val]].append(val)

    def pop(self) -> int:
        val = self.freqBuckets[self.maxFreq].pop()
        # this is tricky part
        # if an element A appears 3 time
        # freqBuckets[3] will have A for the 3rd time
        # freqBuckets[2] will have A for the 2nd time appearance
        # freqBuckets[1] will have A for the 1st time appearance
        # therefore, when maxFreq decrease from 3.. it will always be 2
        # no need to worry, there is no buckets[2]
        if not self.freqBuckets[self.maxFreq]:
            self.maxFreq -= 1
        self.freqMap[val] -= 1
        return val


""""
Runtime: 1013 ms, faster than 18.26% of Python3 online submissions for Maximum Frequency Stack.
Memory Usage: 23.2 MB, less than 9.39% of Python3 online submissions for Maximum Frequency Stack.
"""


if __name__ == "__main__":

    calls = ["FreqStack", "push", "push", "push", "push", "push", "push", "pop", "push",
             "pop", "push", "pop", "push", "pop", "push", "pop", "pop", "pop", "pop", "pop", "pop"]
    params = [[], [4], [0], [9], [3], [4], [2], [], [6],
              [], [1], [], [1], [], [4], [], [], [], [], [], []]

    fs = None
    for call, param in zip(calls, params):
        if call == "FreqStack":
            fs = FreqStack()
        elif call == "push":
            fs.push(param[0])
        else:
            print(fs.pop())
