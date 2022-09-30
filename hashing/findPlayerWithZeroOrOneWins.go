package hashing

import "sort"

func findWinners(matches [][]int) [][]int {
	m := map[int]int{}

	for _, match := range matches {
		m[match[0]] += 0
		m[match[1]]++
	}

	ans := make([][]int, 2)
	ans[0] = make([]int, 0)
	ans[1] = make([]int, 0)

	for k, v := range m {
		if v == 0 {
			ans[0] = append(ans[0], k)
		}
		if v == 1 {
			ans[1] = append(ans[1], k)
		}
	}

	sort.Ints(ans[0])
	sort.Ints(ans[1])

	return ans
}

// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
