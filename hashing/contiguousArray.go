package hashing

func findMaxLength(nums []int) int {
	m := map[int]int{}
	// This line is tricky but, necessary for calculating len of subarr
	m[0] = -1
	maxLen, sum := 0, 0
	for idx, val := range nums {
		sum += (2 * val) - 1
		if midx, ok := m[sum]; ok {
			maxLen = max(maxLen, idx-midx)
		} else {
			m[sum] = idx
		}
	}
	return maxLen
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// https://leetcode.com/problems/contiguous-array/
