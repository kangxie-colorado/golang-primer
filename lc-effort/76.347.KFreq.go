// https://leetcode.com/problems/top-k-frequent-elements/
package main

import "container/heap"

/*
easily I am thinking this is a hashmap
then priority queue
*/

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	pq := make(PriorityQueueNC, 0)
	heap.Init(&pq)

	for k, v := range m {
		heap.Push(&pq, &NumAndCount{num: k, count: v, index: 0})
	}

	res := []int{}
	for k > 0 {
		item := heap.Pop(&pq)
		e, _ := item.(*NumAndCount)
		res = append(res, e.num)
		k--
	}

	return res
}

/*
Runtime: 24 ms, faster than 23.56% of Go online submissions for Top K Frequent Elements.
Memory Usage: 5.9 MB, less than 34.66% of Go online submissions for Top K Frequent Elements.

not really too bad..
let me try just using the array

nah.. I think.. 1 <= nums.length <= 105
when array is small, it is actually wasting

let me just sort and count
nah... still

rerun it
Runtime: 11 ms, faster than 91.62% of Go online submissions for Top K Frequent Elements.
Memory Usage: 5.5 MB, less than 85.65% of Go online submissions for Top K Frequent Elements.

actually this is as good as I can get

func topKFrequent(nums []int, k int) []int {

    freq := make([][]int, len(nums) + 1)
    tempMap := make(map[int]int)

    for i := 0; i < len(nums); i++ {
        if count, ok := tempMap[nums[i]]; ok {
            tempMap[nums[i]] = count + 1
        } else {
            tempMap[nums[i]] = 1
        }
    }

    for key, val := range tempMap {
        freq[val] = append(freq[val], key)
    }

    res := []int{}
    for i := len(freq) - 1; i >= 0; i-- {
        for _, num := range freq[i] {
            res = append(res, num)
            if len(res) == k {
                return res
            }
        }
    }

    return nil

}

gotta admit, this reverse mapping is smart too

cool.. it has a name
bucketSort..

because at most len+1 buckets..
*/

type NumAndCount struct {
	num   int
	count int
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueueNC []*NumAndCount

func (pq PriorityQueueNC) Len() int { return len(pq) }

func (pq PriorityQueueNC) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count > pq[j].count
}

func (pq PriorityQueueNC) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueueNC) Push(x interface{}) {
	n := len(*pq)
	item := x.(*NumAndCount)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueueNC) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueueNC) update(item *NumAndCount, num, count int) {
	item.num = num
	item.count = count
	heap.Fix(pq, item.index)
}
