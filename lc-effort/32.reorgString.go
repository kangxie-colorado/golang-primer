// https://leetcode.com/problems/reorganize-string/

/*
analysis
	this has something similar to do with the 22.longestHappyString.go
	just use a priority quest, pop the biggest then push it back

	it will utilize a fact, that the 2nd biggest will be at the top next time
	if it is the same pop again... ah... no all the same

	hmm.. then pop twice.. and write two..
	yeapp, that will do it

*/

package main

import "container/heap"

// An Item is something we manage in a priority queue.
type StringWithCount struct {
	count int
	str   byte
	index int // The index of the item in the heap.

}

// A PriorityQueue implements heap.Interface and holds Items.
type StringWithCountPQ []*StringWithCount

func (pq StringWithCountPQ) Len() int { return len(pq) }

func (pq StringWithCountPQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count > pq[j].count
}

func (pq StringWithCountPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *StringWithCountPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*StringWithCount)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *StringWithCountPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *StringWithCountPQ) update(item *StringWithCount, count int, str byte) {
	item.count = count
	item.str = str
	heap.Fix(pq, item.index)
}

func reorganizeString(s string) string {
	counter := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	pq := make(StringWithCountPQ, 0)
	index := 0
	for k, v := range counter {
		heap.Push(&pq, &StringWithCount{v, k, index})
		index++
	}
	heap.Init(&pq)

	str := []byte{}
	for pq.Len() > 1 {
		s := heap.Pop(&pq)
		s1 := s.(*StringWithCount)
		s = heap.Pop(&pq)
		s2 := s.(*StringWithCount)

		str = append(str, s1.str)
		s1.count -= 1
		str = append(str, s2.str)
		s2.count -= 1

		if s1.count > 0 {
			heap.Push(&pq, s1)
		}

		if s2.count > 0 {
			heap.Push(&pq, s2)
		}
	}

	if pq.Len() == 0 {
		return string(str)
	}

	r := pq.Pop()
	sr := r.(*StringWithCount)

	if sr.count >= 2 {
		return ""
	} else {
		str = append(str, sr.str)
	}

	return string(str)

}
