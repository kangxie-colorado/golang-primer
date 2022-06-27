// https://leetcode.com/problems/longest-happy-string/

package main

import (
	"container/heap"
	"fmt"
)

/*
	we can firstly looking at combining 2 smaller numebers. let it be a b
	and let a <= b

	then to accomondate b, we at least need (b-1)/2 a...
	e.g. bbabbabb b=6 a=2
	if b=2, 2 a cannot take all of them, at most take bbabbabb

	then take this string to acconmondate c
	now a or b can both be the separator ..

	but there is a catch... with a c, the extra b might just be okay..

	so maybe I should focus on two bigger numbers and accomondate the smallest?

	hmm... nah.. 1/1/7, that will leave c unused.. still have to come back
	so maybe I need to comeback...

	let me code a bit


*/

func longestDiverseStringHelper(a int, b int, c int, stra, strb, strc string) string {
	// a <= b <= c
	// combine 2 bigger first, it will be at least b length
	// actuall bbccbbccbbcc... min(4b+2, b+c)
	// so a surely can be absorbed..
	// then looking at c...

	// form -> cbcb...bc
	str := ""
	b2 := b
	for b > 0 {
		str += strc + strb
		b--
		c--
	}
	if c > 0 {
		str += strc
		c--
	}

	moreC := 0
	idx := 0
	for moreC < b2+1 && c > 0 {
		// if more c, then replace "c" with "cc" as much as it can
		str = str[:idx] + strc + strc + str[idx+1:]
		idx += 3

		c--
		moreC++
	}

	// now it is in this form : ccbccbccb..cbcbc
	// now insert a
	// can just change c into "ca" for a times
	// becaome a<=c.. so this must be ok

	// if we still have c then it must be first c full blown, ccbccb...ccbcc
	// and 1 a can at most take 2 c, no matter how you place it..
	// to the end, in the middle, the net gain is the same.. so just append to the end..
	// ccbccb...ccbccaccacc..acc
	for c > 2 && a > 0 {
		str += stra + strc + strc
		c -= 2
		a--
	}
	// c is done
	// now if we still have a left, just replace b to ba...
	prefix := ""
	idx = 0
	for a > 0 {
		prefix += stra + str[idx:idx+1]
		idx++
		a--
	}

	str = prefix + str[idx:]

	// a can inserted into
	// the end... carry two more c
	// besiedss b, then between a and b, two more c
	// between cc, then each end can become cc, two more c..
	return str
}

func _1_longestDiverseString(a int, b int, c int) string {

	stra, strb, strc := "a", "b", "c"

	if a > b {
		a, b = b, a
		stra, strb = strb, stra
	}

	if a > c {
		a, c = c, a
		stra, strc = strc, stra
	}

	if b > c {
		b, c = c, b
		strb, strc = strc, strb
	}

	// fmt.Println(a, stra, b, strb, c, strc)

	return longestDiverseStringHelper(a, b, c, stra, strb, strc)
}

/*
Success
Details
Runtime: 1 ms, faster than 36.36% of Go online submissions for Longest Happy String.
Memory Usage: 2.1 MB, less than 72.73% of Go online submissions for Longest Happy String.
Next challenges:

Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Happy String.
Memory Usage: 2.2 MB, less than 45.45% of Go online submissions for Longest Happy String.
Next challenges:

the priority queue looks so elegant. let me code it out
*/

// An Item is something we manage in a priority queue.
type abc struct {
	count int
	str   string
	index int // The index of the item in the heap.

}

// A PriorityQueue implements heap.Interface and holds Items.
type abcPQ []*abc

func (pq abcPQ) Len() int { return len(pq) }

func (pq abcPQ) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].count > pq[j].count
}

func (pq abcPQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *abcPQ) Push(x interface{}) {
	n := len(*pq)
	item := x.(*abc)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *abcPQ) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *abcPQ) update(item *abc, count int, str string) {
	item.count = count
	item.str = str
	heap.Fix(pq, item.index)
}

func longestDiverseString(a int, b int, c int) string {
	str := ""

	pq := make(abcPQ, 0)

	if a > 0 {
		heap.Push(&pq, &abc{count: a, str: "a", index: 0})
	}
	if b > 0 {
		heap.Push(&pq, &abc{count: b, str: "b", index: 1})
	}
	if c > 0 {
		heap.Push(&pq, &abc{count: c, str: "c", index: 2})
	}
	heap.Init(&pq)

	for pq.Len() > 1 {
		s := heap.Pop(&pq)
		s1 := s.(*abc)
		s = heap.Pop(&pq)
		s2 := s.(*abc)

		if s1.count >= 2 {
			str += s1.str + s1.str
			s1.count -= 2
		} else {
			str += s1.str
			s1.count -= 1
		}

		if s2.count >= 2 && s2.count >= s1.count {
			str += s2.str + s2.str
			s2.count -= 2
		} else {
			str += s2.str
			s2.count -= 1
		}

		if s1.count > 0 {
			heap.Push(&pq, s1)
		}

		if s2.count > 0 {
			heap.Push(&pq, s2)
		}
	}

	if pq.Len() == 0 {
		return str
	}

	r := pq.Pop()
	sr := r.(*abc)

	if sr.count >= 2 {
		str += sr.str + sr.str
	} else {
		str += sr.str
	}

	return str
}

/*
Success
Details
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Happy String.
Memory Usage: 2.2 MB, less than 59.09% of Go online submissions for Longest Happy String.
Next challenges:

so I think I might could pop 1 at a time..
no.. you cannot...

you might end pop the same str next time, and stalemate..
*/

func testLongestDiverseString() {
	fmt.Println(longestDiverseString(1, 8, 12))
	fmt.Println(longestDiverseString(1, 1, 7))
	fmt.Println(longestDiverseString(7, 1, 0))
}
