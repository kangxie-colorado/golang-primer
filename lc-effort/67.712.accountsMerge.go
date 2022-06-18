// https://leetcode.com/problems/accounts-merge/

/*
	seems like just build a reverce map
	email to idx.. then can merge on the go

*/

package main

import (
	"fmt"
	"sort"
)

func sortUniq(emails []string) []string {
	sort.Slice(emails, func(i1, i2 int) bool { return emails[i1] < emails[i2] })
	res := []string{}
	for i, e := range emails {
		if i == 0 || e != emails[i-1] {
			res = append(res, e)
		}
	}

	return res
}

func _1_naive_accountsMerge(accounts [][]string) [][]string {
	m := make(map[string]int)

	for i, acc := range accounts {

		seen := false
		for _, email := range acc[1:] {
			merges := []int{i}
			if j, found := m[email]; found {
				merges = append(merges, j)
			}
		}

		if !seen {
			for _, email := range acc[1:] {
				m[email] = i
			}
		}
	}
	exported := make([]int, 1000)
	res := [][]string{}
	for _, v := range m {
		if exported[v] == 1 {
			continue
		}
		res = append(res, append(accounts[v][:1], sortUniq(accounts[v][1:])...))
		exported[v] = 1
	}

	return res
}

/*
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))

	this reveals the difficulties behind this problem
	it is not can hashmap and merge
	there is some union find behind it..
*/

func accountsMerge(accounts [][]string) [][]string {
	m := make(map[string]int)
	unions := make([]int, len(accounts))
	for u := range unions {
		unions[u] = u
	}

	for i, acc := range accounts {

		seen := false

		for _, email := range acc[1:] {
			if j, found := m[email]; found {
				union(i, j, unions)
				seen = true
			}
		}

		if !seen {
			for _, email := range acc[1:] {
				m[email] = i
			}
		} else {
			for _, email := range acc[1:] {
				m[email] = find(i, unions)
			}
		}
	}

	unionMap := make(map[int][]int)
	for i := range unions {
		unionMap[find(i, unions)] = append(unionMap[find(i, unions)], i)
	}

	res := [][]string{}
	for _, accs := range unionMap {
		name := accounts[accs[0]][0]
		merged := []string{}
		for _, acc := range accs {
			merged = append(merged, accounts[acc][1:]...)
		}

		res = append(res, append([]string{name}, sortUniq(merged)...))
	}

	return res
}

/*
Runtime: 55 ms, faster than 82.52% of Go online submissions for Accounts Merge.
Memory Usage: 8.4 MB, less than 63.11% of Go online submissions for Accounts Merge.
*/

func testAccountsMerge() {
	fmt.Println(accountsMerge([][]string{{"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"John", "johnsmith@mail.com", "john00@mail.com"}, {"Mary", "mary@mail.com"}, {"John", "johnnybravo@mail.com"}}))
	fmt.Println(accountsMerge([][]string{{"David", "David0@m.co", "David1@m.co"}, {"David", "David3@m.co", "David4@m.co"}, {"David", "David4@m.co", "David5@m.co"}, {"David", "David2@m.co", "David3@m.co"}, {"David", "David1@m.co", "David2@m.co"}}))
}
