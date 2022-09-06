"""
https://leetcode.com/problems/top-k-frequent-elements/

I knew heap.
also bucket sort or bucket something

let me do that
"""


from typing import Counter, DefaultDict, List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        buckets = [set()]*(len(nums)+1)
        for i in range(len(buckets)):
            buckets[i] = set()
        m = DefaultDict(int)  # num:bucket

        for n in nums:
            times = m[n]+1

            m[n] = times
            buckets[times].add(n)
            if times > 1:
                buckets[times-1].remove(n)
        res = []

        for i in range(len(nums), -1, -1):
            res.extend(list(buckets[i]))
            k -= len(buckets[i])
            if k == 0:
                break

        return res


"""
while I think this is O(N)
I think there is no need to update the buckets so many time..
since I am using the memory O(N) why do I need to do that
"""


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        buckets = [set()]*(len(nums)+1)
        for i in range(len(buckets)):
            buckets[i] = set()

        C = Counter(nums)

        for num, count in C.items():
            buckets[count].add(num)

        res = []
        for i in range(len(nums), -1, -1):
            res.extend(list(buckets[i]))
            # because the answer is guranteed to be unique so no need to consider half
            # of the set can be used and left the other half
            # because that would make the answer non-uniq
            k -= len(buckets[i])
            if k == 0:
                break

        return res


"""
still not good
replace set with []
"""


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        buckets = [[]]*(len(nums)+1)
        for i in range(len(buckets)):
            buckets[i] = []

        C = Counter(nums)

        for num, count in C.items():
            buckets[count].append(num)

        res = []
        for i in range(len(nums), -1, -1):
            res.extend((buckets[i]))
            # because the answer is guranteed to be unique so no need to consider half
            # of the set can be used and left the other half
            # because that would make the answer non-uniq
            k -= len(buckets[i])
            if k == 0:
                break

        return res


"""
Runtime: 132 ms, faster than 75.18% of Python3 online submissions for Top K Frequent Elements.
Memory Usage: 19.4 MB, less than 28.59% of Python3 online submissions for Top K Frequent Elements.
"""


if __name__ == '__main__':
    s = Solution()
    print(s.topKFrequent(nums=[1, 1, 1, 2, 2, 3], k=2))
